package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	var logFile string
	flag.StringVar(&logFile, "file", "logfile", "The file to parse")
	myMap := make(map[string]int)
	sorted := make([]string, 0, len(myMap))

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

	// for key, val := range myMap {
	// 	fmt.Printf("%v: %v\n", key, val)
	// }

	for v := range myMap {
		sorted = append(sorted, v)
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		return myMap[sorted[i]] > myMap[sorted[j]]
	})

	// fmt.Println("\n\n")

	count := 0
	for _, k := range sorted {
		if count == 10 {
			break
		}
		fmt.Printf("%s:%d\n", k, myMap[k])
		count++
	}
}
