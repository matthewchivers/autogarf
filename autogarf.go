package main

import (
	"log"
	"path/filepath"
	"time"
)

var OS string

func main() {
	// Fetch configuration
	config, err := readConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Prep the working directory based on the configuration
	directory := filepath.FromSlash(config.ClientDir)

	// Get list of client directories
	var clients = []*Client{}
	clients = getClientDirectories(directory, clients)

	// Get list of statements for each client
	for _, client := range clients {
		client.populateStatementList()
	}

	// Loop through client folders - copy incumbent file to a new file with this month's date
	for _, client := range clients {
		log.Printf("%s", client.name)
		incumbentFileName := client.getIncumbentFileName()
		newFileName := client.getNewFileName()
		// Only create a new file if there isn't one already present for this month
		if client.incumbentStatementDate.After(time.Now()) {
			log.Printf("The most recent statement (\"%s\") already has this month's date. No action being taken.", incumbentFileName)
		} else {
			log.Printf("Duplicating the most recent statement \"%s\" to have this month's date \"%s\"", incumbentFileName, newFileName)
			copyFile(client.getIncumbentFilePath(), client.getNewFilePath())
		}
	}
}
