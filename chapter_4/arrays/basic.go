package main

import "fmt"

func printSlice(sl []int) {
	for i, v := range sl {
		fmt.Println(i, "->", v)
	}
}

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, "->", v)
	}
}

func main() {
	a := []int{1, 2, 5, 2}
	printSlice(a)
	b := [...]int{10, 20, 30, 40, 50}
	printArray(b)
}
