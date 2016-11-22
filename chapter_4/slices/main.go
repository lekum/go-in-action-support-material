package main

import "fmt"

func main() {
	slice1 := make([]string, 5)    // length and capacity of 5 elements
	slice2 := make([]string, 3, 5) // length of 3 and capacity of 5 elements
	fmt.Println(slice1, slice2)
	empty_slice := make([]int, 0)
	var nil_slice []int
	fmt.Println(empty_slice, nil_slice)
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)
	slice[1] = 25
	fmt.Println(slice)
	fmt.Println(slice[1:3])
	slice3 := slice[2:4]
	slice3[0] = -3
	fmt.Println(slice)
	slice3 = append(slice3, 80)
	fmt.Println(slice)
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice4 := source[2:3:4]
	fmt.Println(slice4)
}
