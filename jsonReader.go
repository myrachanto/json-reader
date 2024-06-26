package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

// Struct to match your JSON structure
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// readJSONFile reads the json files
func readJSONFile(filename string, wg *sync.WaitGroup, ch chan<- Item) {
	defer wg.Done()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error reading file:", filename, err)
		return
	}

	var data Item
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Println("Error unmarshalling JSON from file:", filename, err)
		return
	}

	ch <- data
}

func main() {
	filenames := []string{"file1.json", "file2.json", "file3.json"} // List your JSON files here

	var wg sync.WaitGroup
	ch := make(chan Item, len(filenames))

	for _, filename := range filenames {
		wg.Add(1)
		go readJSONFile(filename, &wg, ch)
	}

	wg.Wait()
	close(ch)
	// fanin compile the read json files
	var allData []Item
	for data := range ch {
		allData = append(allData, data)
	}

	// Process allData as needed
	fmt.Println(allData)
}
