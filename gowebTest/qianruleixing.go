package main

import "fmt"

type user struct {
	name string
	email string
}

type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("%s ====== %s",u.name,u.email)
}
func (a *admin) notify()  {
	fmt.Printf("%s +++++++++ %s",a.user.name,a.email)
}

type notifier interface {
	notify()
}

func main() {
	ad := admin{
		user:user{
			name:"lisi",
			email:"lisi@gmail.com",
		},
		level:"papas",
	}
	sendnotify(&ad)
	ad.user.notify()
}
func sendnotify(n notifier){
	n.notify()
}