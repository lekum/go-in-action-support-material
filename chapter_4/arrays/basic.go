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
	var c [5]int
	c = b
	c[0] = 1

	fmt.Println("b after c changes:")
	printArray(b)

	fmt.Println("c after change:")
	printArray(c)

	ar := [...]*string{new(string), new(string), new(string)}
	*ar[0] = "Red"
	*ar[1] = "Blue"
	*ar[2] = "Green"
	ab := ar
	*ar[2] = "Yellow"
	for _, x := range ar {
		fmt.Println(*x)
	}
	for _, x := range ab {
		fmt.Println(*x)
	}
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	for _, j := range array {
		fmt.Println(j)
	}
	array = [4][2]int{0: {0: 11}, 2: {0: 1, 1: 2}}
	for _, j := range array {
		fmt.Println(j)
	}
}
