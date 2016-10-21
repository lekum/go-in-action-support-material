package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func printCountdown(idx int, wgp *sync.WaitGroup) {
	time.Sleep(time.Duration(idx) * time.Second)
	fmt.Println("My index is", strconv.Itoa(idx))
	(*wgp).Done()
}

func main() {
	threads := 30
	var wgp sync.WaitGroup
	wgp.Add(threads)
	for i := 0; i < threads; i++ {
		go printCountdown(i, &wgp)
	}
	fmt.Println("Waiting...")
	wgp.Wait()
	fmt.Println("Done!")
}
