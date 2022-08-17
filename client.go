package main

type Client struct {
	name      string
	directory string
}

func newClient(name, directory string) *Client {
	c := Client{name: name}
	c.directory = directory
	return &c
}
