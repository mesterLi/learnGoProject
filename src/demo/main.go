package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func main() {
	//var f *os.File
	//if isExist := checkFileIsExist("./aaa.html"); isExist {
	//	f, _ = os.OpenFile("./aaa.html", os.O_APPEND, 0666)
	//	log.Print("打开文件")
	//} else {
	//	f, _ = os.Create("./aaa.html")
	//	log.Print("创建文件")
	//}
	resp, err := http.Get("https://www.doutula.com/photo/4486735")
	if err != nil {
		log.Fatal("err:{}", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	//log.Print(json.Marshal(resp.Body))
	if err != nil {
		log.Fatal("NewDocumentFromReader err:{}", err)
	}
	fmt.Print("doc:{}", doc.Find(".list-group-item").Nodes)
	doc.Find(".list-group-item").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, _ := s.Find("a").Attr("href")
		//title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band)
	})
}

func checkFileIsExist(filename string) bool {
	exit := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exit = false
	}
	return exit
}