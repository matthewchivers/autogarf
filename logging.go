package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func setLoggingOutput() {
	directory := filepath.FromSlash("logs")
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	logDate := time.Now().Format("2006-01-02 150405")
	logFileName := fmt.Sprintf("%s.log", logDate)
	logName := filepath.Join("logs", logFileName)
	f, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

// Delete log files older than 6 months
func rotateLogFiles() {
	log.Printf("Rotating log files")
	directory := filepath.FromSlash("logs")
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("error reading directory: %v", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		log.Printf("Found log file: %s", file.Name())
		oneYearAgo := time.Now().Add(-365 * 24 * time.Hour)
		fileDateString := strings.Split(file.Name(), " ")[0]
		fileDate, err := time.Parse("2006-01-02", fileDateString)
		if err != nil {
			log.Fatalf("error parsing date: %v", err)
		}
		if oneYearAgo.After(fileDate) {
			log.Printf("Deleting log file: %s", file.Name())
			err := os.Remove(filepath.Join(directory, file.Name()))
			if err != nil {
				log.Fatalf("error deleting file: %v", err)
			}
		}
		if err != nil {
			log.Fatalf("error removing file: %v", err)
		}
	}
	log.Printf("Finished rotating log files")
}
