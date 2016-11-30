package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type Duration int64

func main() {
	var bill user
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(bill)
	fmt.Println(lisa)
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)
}
