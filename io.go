package main

import (
	"bufio"
	"log"
	"os"
)

func readFile() []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f := bufio.NewScanner(file)

	var lines []string

	// Loop over the file line by line, storing it in the lines slice
	for f.Scan() {
		lines = append(lines, f.Text())
	}

	return lines
}
