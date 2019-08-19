package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(CreateCaptcha())
}
func CreateCaptcha() string {
	return fmt.Sprintf("%06v",
		rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}