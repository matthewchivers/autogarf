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
	name                        string
	directory                   string
	statements                  []Statement
	incumbentStatementNameIndex *int
}

type Statement struct {
	fileName string
	date     time.Time
}

// Creates a new client
func newClient(name, directory string) *Client {
	c := Client{name: name}
	c.directory = directory
	return &c
}

// Add statement to client
func (c *Client) newStatement(statement Statement) {
	c.statements = append(c.statements, statement)
	log.Printf("Found statement: \"%s\" with date \"%s\" - added to list for client \"%s\"", statement.fileName, statement.date.Format("2006-01-02"), c.name)
}

// Returns a list of client directories
func getClientDirectories(directory string, clients []*Client) []*Client {
	log.Printf("Getting client directories:")
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
	log.Printf("Finished getting client directories:")
	return clients
}

// Populates the statement list
func (c *Client) populateStatementList() ([]Statement, error) {
	client_dir, err := os.Open(c.directory)
	if err != nil {
		return nil, err
	}
	defer client_dir.Close()

	files, err := client_dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	for _, raw_file_name := range files {
		if strings.HasPrefix(raw_file_name, "Statement") && strings.HasSuffix(raw_file_name, ".docx") {
			file_name := correctDoubleSpaces(raw_file_name)
			fileDate, err := getFileDate(file_name)
			if err != nil {
				return nil, err
			}
			c.newStatement(Statement{raw_file_name, fileDate})
		}
	}
	return c.statements, nil
}

// Returns the incumbent statement file name
func (c *Client) getIncumbentFileName() string {
	if c.incumbentStatementNameIndex == nil {
		c.incumbentStatementNameIndex = new(int)
		*c.incumbentStatementNameIndex = 0
		for i, statement := range c.statements {
			if statement.date.After(c.statements[*c.incumbentStatementNameIndex].date) {
				*c.incumbentStatementNameIndex = i
			}
		}
	}
	return c.statements[*c.incumbentStatementNameIndex].fileName
}

// Extracts the date from the statement file name
func getFileDate(statementName string) (time.Time, error) {
	dateString, err := extractDateString(statementName)
	if err != nil {
		return time.Now(), err
	}
	return convertStringToDate(dateString), nil
}

// Extracts the date string from the statement file name
func extractDateString(statementName string) (string, error) {
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
	incumbentFileName := c.getIncumbentFileName()
	formattedIncumbentFileName := correctDoubleSpaces(incumbentFileName)
	hyphenSplit := strings.Split(formattedIncumbentFileName, " - ")
	filePrefix := hyphenSplit[0]
	fileSuffix := strings.Split(hyphenSplit[1], ".")[1]

	firstOfCurrentMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
	lastOfNextMonth := firstOfCurrentMonth.AddDate(0, 1, -1)
	return fmt.Sprintf("%s - %s.%s", filePrefix, lastOfNextMonth.Format("02 Jan 06"), fileSuffix)
}

// Returns the entire path to the new statement file
func (c *Client) getNewFilePath() string {
	return filepath.Join(c.directory, c.getNewFileName())
}

// Returns the entire path to the incumbent statement file
func (c *Client) getIncumbentFilePath() string {
	return filepath.Join(c.directory, c.getIncumbentFileName())
}

// Removes double spaces from a string
func correctDoubleSpaces(fileName string) string {
	return strings.Join(strings.Fields(fileName), " ")
}
