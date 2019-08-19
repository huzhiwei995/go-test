package main

import (
	"strings"
	"os"
	"fmt"
)

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
