package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var logFile string
	flag.StringVar(&logFile, "file", "logfile", "The file to parse")
	myMap := make(map[string]int)

	flag.Parse()

	openTheFile, err := os.Open(logFile)
	if err != nil {
		fmt.Printf("error opening file %s: %v", logFile, err)
		return
	}

	defer openTheFile.Close()

	file := bufio.NewScanner(openTheFile)
	file.Split(bufio.ScanWords)

	for file.Scan() {
		myMap[file.Text()] += 1
	}
	// order the map highest count first
	// print first 10
	for key, val := range myMap {
		fmt.Printf("%v: %v\n", key, val)
	}
}
