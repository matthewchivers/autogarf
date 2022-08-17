package main

import (
	"log"
	"path/filepath"
	"time"
)

var OS string

func main() {
	config, err := readConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	directory := filepath.FromSlash(config.ClientDir)

	var clients = []*Client{}
	clients = getClientDirectories(directory, clients)

	for _, client := range clients {
		client.populateStatementList()
	}

	// Copy the files
	for _, client := range clients {
		log.Printf("%s", client.name)
		// for _, statement := range client.statements {
		// 	log.Printf("%s", statement)
		// }
		latestFileName := client.getLatestFileName()
		newFileName := client.getNewFileName()
		if client.latestStatementDate.After(time.Now()) {
			log.Printf("The most recent statement is already the most current. Not copying: %s", latestFileName)
		} else {
			log.Printf("Duplicating the most recent statement \"%s\" to have this month's date \"%s\"", latestFileName, newFileName)
			copyFile(client.getCurrentFilePath(), client.getNewFilePath())
		}
	}
}
