package main

import (
	"log"
	"os"
	"path/filepath"
)

var OS string

func main() {
	config, err := readConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	directory := filepath.FromSlash(config.ClientDir)

	var clients = []*Client{}

	clients = getClients(directory, clients)

	for _, client := range clients {
		log.Printf("Client -\n\tName: %s\n\tDirectory: %s", client.name, client.directory)
	}

}

func getClients(directory string, clients []*Client) []*Client {
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
