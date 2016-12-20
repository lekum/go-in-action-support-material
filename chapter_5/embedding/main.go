package main

import "fmt"

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user  // Embedded type
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "John Smith",
			email: "john@email.com",
		},
		level: "super",
	}
	sendNotification(&ad)
	ad.user.notify()
	ad.notify()
}