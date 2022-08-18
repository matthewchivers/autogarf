package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

var OS string

func main() {
	setLoggingOutput()
	rotateLogFiles()

	config := getConfig()

	directory := filepath.FromSlash(config.ClientDir)

	var clients = []*Client{}
	clients = getClientDirectories(directory, clients)
	printClientDirectories(clients)

	processStatements(clients)
	copyStatements(clients)

	log.Printf("Finished")
}

func getConfig() *conf {
	log.Printf("Getting config")
	config, err := readConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished getting config")
	return config
}

func processStatements(clients []*Client) {
	log.Printf("Processing statements:")
	for _, client := range clients {
		client.populateStatementList()
	}
	log.Printf("Finished statement processing")
}

// Loop through client folders - copy incumbent file to a new file with this month's date
func copyStatements(clients []*Client) {
	log.Printf("Copying statements:")
	for _, client := range clients {
		log.Printf("Actioning: %s", client.name)
		incumbentFileName := client.getIncumbentFileName()
		newFileName := client.getNewFileName()
		// Only create a new file if there isn't one already present for this month
		if correctDoubleSpaces(incumbentFileName) == correctDoubleSpaces(newFileName) {
			log.Printf("The most recent statement (\"%s\") already has the new date. No action being taken.", incumbentFileName)
		} else {
			log.Printf("Duplicating the incumbent statement \"%s\" and renaming with this month's date \"%s\"", incumbentFileName, newFileName)
			copyFile(client.getIncumbentFilePath(), client.getNewFilePath())
		}
	}
	log.Printf("Finished copying statements")
}

// Print all client names from the list supplied
func printClientDirectories(clients []*Client) {
	log.Printf("Printing client directories:")
	log.Printf("Found %d clients", len(clients))
	var sb strings.Builder
	for i, client := range clients {
		sb.WriteString(fmt.Sprintf("\"%s\"", client.name))
		if i < len(clients)-1 {
			sb.WriteString(", ")
		}
	}
	log.Printf("Client directories: %s", sb.String())
	log.Printf("Finished printing client directories")
}
