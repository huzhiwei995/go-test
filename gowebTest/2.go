package main

import "fmt"

type Master struct {
	name string
	age int
}
func main() {
	var bils Master
	bils.name = "mark"
	fmt.Println(bils)
}