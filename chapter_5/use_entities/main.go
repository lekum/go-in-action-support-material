package main

import (
	"fmt"
	"github.com/lekum/go-in-action-support-material/chapter_5/entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "Jannette"
	a.Email = "jannette@email.com"
	fmt.Printf("Admin: %v\n", a)
}
