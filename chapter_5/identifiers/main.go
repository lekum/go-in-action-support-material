package main

import (
	"fmt"
	"github.com/lekum/go-in-action-support-material/chapter_5/counters"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)
}
