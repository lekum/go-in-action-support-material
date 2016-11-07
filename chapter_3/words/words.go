package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

func main() {
	filename := os.Args[1]
	log.SetOutput(os.Stdout)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text.\n", count)
}
