package main

import (
	"io"
	"log"
	"os"
)

// CopyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	log.Printf("Copying \"%s\" to \"%s\"", src, dst)

	// Read all of the bytes in the source file
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Create the destination file
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the bytes from the source to the destination
	_, err = io.Copy(out, in)
	cerr := out.Close()

	// If there was an error writing to the destination, return an error
	if err != nil {
		return err
	}
	// If there was an error closing the destination, return an error
	return cerr
}
