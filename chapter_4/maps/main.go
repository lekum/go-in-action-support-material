package main

import "fmt"

func main() {
	colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	colors["Blue"] = "#ef0114"
	col := "Red"
	if value, exists := colors[col]; exists {
		fmt.Println("colors[", col, "] - >", value)
	} else {
		fmt.Println(col, "not found")
	}
	for key, value := range colors {
		fmt.Println(key, "->", value)
	}
	colors["AliceBlue"] = "#f0f8ff"
	delete(colors, "Blue")
	fmt.Println(colors)
}
