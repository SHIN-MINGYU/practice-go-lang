// will use go query

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)


var baseURL string = "https://search.naver.com/search.naver?display=15&f=&filetype=0&page=3&query=%ED%81%AC%EB%A1%A4%EB%A7%81&research_url=&sm=tab_pge"

func main(){
	pages := getPages()
	titles := []string{}
	c := make(chan []string)
	for i:=0; i<pages; i++ {
		go getTitles(i, c)
	}

	for i:=0; i<pages; i++ {
		titles = append(titles, <-c...)
	}

	writeToFile(titles)
}


func getTitles(page int, mainC chan<- []string) {
	titles := []string{}
	c := make(chan string)
	pageUrl := baseURL + "&start=" + strconv.Itoa(15*(page-1) + 1)

	fmt.Println("Checking : ", pageUrl)
	if(page != 0) {
		pageUrl += "&where=web"
	}
	res, err := http.Get(pageUrl)

	checkErr(err)
	checkCode(res.StatusCode)

	defer res.Body.Close()

	doc, err2 := goquery.NewDocumentFromReader(res.Body)

	checkErr(err2)

	result := doc.Find(".total_tit>a")

	result.Each(func(i int, s *goquery.Selection) {
		go getText(s,c)
	})

	for i:=0;i<result.Length();i++{
		title := <-c
		titles = append(titles, title)
	}

	mainC <- titles
}

func getText(titleTag *goquery.Selection, c chan<- string){
	c <- titleTag.Text()
}

func writeToFile(titles []string){	
	file, err := os.Create("title.csv")
	checkErr(err)

	w :=csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"title"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _,title := range titles{
		titleSlice := []string{title}
		wErr2 := w.Write(titleSlice)
		checkErr(wErr2)
	}

}



func getPages() int{
	pages := 0
	req, err := http.NewRequest("GET",baseURL,nil)
	checkErr(err)
	req.Header.Add("User-Agent", "Mozilla/5.0")

	client := &http.Client{}
	res, err2 := client.Do(req)

	checkErr(err2)
	checkCode(res.StatusCode)

	defer res.Body.Close()// I/O byte so we should close when function is close
	// prevent memory leak

	doc, err3 := goquery.NewDocumentFromReader(res.Body)
	checkErr(err3)

	doc.Find(".sc_page_inner").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	
	return pages
}

func checkErr(err error){
	if err != nil{
		log.Fatalln(err)
	}
}

func checkCode(code int){
	if(code != 200){
		log.Fatal("Request failed with Status :", code )
	}

}