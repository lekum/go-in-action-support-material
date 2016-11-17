package main

import "fmt"

func main() {
	slice1 := make([]string, 5)    // length and capacity of 5 elements
	slice2 := make([]string, 3, 5) // length of 3 and capacity of 5 elements
	fmt.Println(slice1, slice2)
	empty_slice := make([]int, 0)
	var nil_slice []int
	fmt.Println(empty_slice, nil_slice)
}
