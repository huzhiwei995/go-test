package main

import (
	"fake-useragent"
	browser "github.com/EDDYCJY/fake-useragent"
	"time"
	"log"
)

func main() {
	//定制请求头
	/*random := browser.Random()
	log.Printf("Random: %s",random)

	chrome := browser.Chrome()
	log.Printf("chrome : %s",chrome)*/
	client := browser.Client{
		MaxPage:3,
		Delay:200 * time.Millisecond,
		Timeout:10*time.Second,
	}
	cache := browser.Cache{
		UpdateFile: true,
	}
	b := browser.NewBrowser(client,cache)
	random := b.Random()
	log.Printf("random: %s",random)
}
