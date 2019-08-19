package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
)

func main() {
	fmt.Println(os.Args[0])
	r,err := http.Get(os.Args[1])
	fmt.Println(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
}