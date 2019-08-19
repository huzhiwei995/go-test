package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	name string
	email string
}
func (a *admin) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	user := user{"lisa","aksjhda@email.com"}
	sendNotification(&user)

	admin := admin{"nanxy","sasasas@email.com"}
	sendNotification(&admin)

}
func sendNotification(n notifier) {
	n.notify()
}