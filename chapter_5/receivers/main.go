package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s<%s>\n",
		u.name,
		u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{name: "Bill", email: "bill@gmail.com"}
	sendNotification(u)

}
