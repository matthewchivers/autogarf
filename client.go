package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Client struct {
	name                   string
	directory              string
	statements             []string
	incumbentStatementName string
	incumbentStatementDate time.Time
}

// Creates a new client
func newClient(name, directory string) *Client {
	c := Client{name: name}
	c.directory = directory
	return &c
}

// Returns a list of client directories
func getClientDirectories(directory string, clients []*Client) []*Client {
	client_dir, err := os.Open(directory)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer client_dir.Close()

	directories, err := client_dir.Readdirnames(0)

	for _, dir_name := range directories {
		client_path := filepath.Join(directory, dir_name)
		clients = append(clients, newClient(dir_name, client_path))
	}
	return clients
}

// Populates the statement list
func (c *Client) populateStatementList() ([]string, error) {
	client_dir, err := os.Open(c.directory)
	if err != nil {
		return nil, err
	}
	defer client_dir.Close()

	files, err := client_dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	for _, file_name := range files {
		if strings.HasPrefix(file_name, "Statement") && strings.HasSuffix(file_name, ".docx") {
			file_name = strings.Replace(file_name, "  ", " ", -1) // Correct double-spaces in file name
			c.statements = append(c.statements, file_name)
		}
	}
	return c.statements, nil
}

// Returns the incumbent statement file name
func (c *Client) getIncumbentFileName() string {
	if c.incumbentStatementName == "" {
		incumbentStatmentIndex := 0
		incumbentStatementDate := time.Now().AddDate(-100, 0, 0) // 100 years ago
		for i, statementName := range c.statements {
			extractedDate, err := extractDate(statementName)
			if err != nil {
				log.Fatal(err)
			}
			statementDate := convertStringToDate(extractedDate)
			if statementDate.After(incumbentStatementDate) {
				incumbentStatementDate = statementDate
				incumbentStatmentIndex = i
			}
		}
		c.incumbentStatementName = c.statements[incumbentStatmentIndex]
		c.incumbentStatementDate = incumbentStatementDate
	}
	return c.incumbentStatementName
}

// Extracts the date from the statement file name
func extractDate(statementName string) (string, error) {
	post := strings.Split(statementName, " - ")[1]
	date := strings.Split(post, ".")[0]
	components := strings.Split(date, " ")
	if len(components) != 3 {
		return "", errors.New(fmt.Sprintf("invalid date format in file name %s", statementName))
	}
	return date, nil
}

// Converts a string to a date
func convertStringToDate(date string) time.Time {
	t, err := time.Parse("02 Jan 06", date)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

// Creates new file name for the new statement
// New file name will always be for the next month.
// e.g. If generating any time in March, the new file name will be for April.
func (c *Client) getNewFileName() string {
	hyphenSplit := strings.Split(c.incumbentStatementName, " - ")
	filePrefix := hyphenSplit[0]
	fileSuffix := strings.Split(hyphenSplit[1], ".")[1]

	firstOfCurrentMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
	lastOfNextMonth := firstOfCurrentMonth.AddDate(0, 2, -1)
	return fmt.Sprintf("%s - %s.%s", filePrefix, lastOfNextMonth.Format("02 Jan 06"), fileSuffix)
}

// Returns the entire path to the new statement file
func (c *Client) getNewFilePath() string {
	return filepath.Join(c.directory, c.getNewFileName())
}

// Returns the entire path to the incumbent statement file
func (c *Client) getIncumbentFilePath() string {
	return filepath.Join(c.directory, c.incumbentStatementName)
}
