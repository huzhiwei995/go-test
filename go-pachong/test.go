package main

import (
	"strings"
	"fmt"
)

func main()  {
	urlimg := "https://anime-pictures.net/pictures/download_image/609673-1100x619-original-denki-long+hair-single-looking+at+viewer-fringe.jpg"
	path := strings.Split(urlimg, "/")
	fmt.Println(path)
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Println(name)
}