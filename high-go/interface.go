package main

import "fmt"

type Color interface {
	getColor() string
	setColor(string)
}

type Car struct {
	color string
}
func (c Car) getColor() string {
	return c.color
}
func (c Car) setColor(s string) {
	c.color = s
}

func main() {
	car := Car{"white"}
	col := Color(car)

	car = col.(Car)         // L(1)
	car.setColor("yellow")
	fmt.Println(col)        // L(2)
	fmt.Println(car)
	car.color = "black"
	fmt.Println(col)        // L(3)
	fmt.Println(car)
}