package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"time"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"io"
	"bytes"
)
var CountNum int
var logger *log.Logger
var logFile *os.File
var imgPath string
var l int
func init() {
	filePath := "./test.txt"
	l = 0
	logFile,err := os.OpenFile(filePath,os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create log err : %v\n", err)
	}
	logger = log.New(logFile,"[INFO]",log.Ldate)
	logger.SetFlags(log.Lshortfile)
	imgPath = "./image/"
}
func main() {
	i := 0
	CountNum = 0
	for {
		err := work("https://anime-pictures.net/pictures/view_posts/"+strconv.Itoa(i)+"?aspect=16%3A9&order_by=date&ldate=0&lang=en")
		if err != nil {
			fmt.Println("连接错误：",err)
			break
		}
		if CountNum > 1 {
			break
		} else {
			fmt.Println("int : ",CountNum)
		}
		i++
		time.Sleep(2*time.Second)
	}
	defer logFile.Close()
}
func work(url string) (err error) {
	Info("url是：",url)
	c := colly.NewCollector(colly.MaxDepth(2))
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	c.OnHTML("*", func(e *colly.HTMLElement) {
		e.DOM.Find(".img_block_big").Each(func(i int, cs *goquery.Selection) {
			a,_ := cs.Find("a").Eq(0).Attr("href")
			c.Visit(e.Request.AbsoluteURL(a))
		})
	})
	c.OnHTML(".download_icon", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = "https://anime-pictures.net"+link
		err := getImg(link)
		if err != nil {
			Info("error")
		}
		l++
	})
	c.OnResponse(func(r *colly.Response) {
		CountNum++
	})
	c.OnError(func(r *colly.Response, err error) {
		Info("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.Visit(url)
	return nil
}
func getImg(urlimg string) error {
	name := "./image/"+"H_"+strconv.Itoa(l)+".jpg"
	out, err := os.Create(name)
	if err != nil{
		Info("创建文件失败",err)
		return err
	}
	defer out.Close()
	resp, err := http.Get(urlimg)
	if err != nil{
		Info("请求图片url失败",err)
		return err
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		Info("读取文件失败",err)
		return err
	}
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		Info("下载文件失败",err)
	}
	return err
}
func Info(v ...interface{}) {
	logger.SetPrefix("[Info]")
	logger.Println(v)
}