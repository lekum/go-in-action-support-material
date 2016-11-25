package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}
	slice2 := [][]int{{10}, {100, 200}}
	fmt.Println(slice2)
	slice2[0] = append(slice2[0], 20)
	fmt.Println(slice2)
}
