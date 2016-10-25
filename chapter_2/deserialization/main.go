package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func main() {
	dataFile := "data.json"
	file, err := os.Open(dataFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		log.Fatalln(err)
	}
	for n, feed := range feeds {
		fmt.Println(strconv.Itoa(n))
		fmt.Println(feed.Name)
		fmt.Println(feed.URI)
		fmt.Println(feed.Type)
		fmt.Println()
	}
}
