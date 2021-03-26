package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	// "golang.org/x/sync/errgroup"
)

type News struct {
	// URL   string `json:"url"`
	Title string `json:"title"`
}

type NewsList struct {
	TotalNews int    `json:"total_ebooks"`
	List      []News `json:"ebooks"`
}

func New_NewsList() *NewsList {
	return &NewsList{}
}
func main() {
	BaseURL := "https://www.thesaigontimes.vn/121624/Cuoc-cach-mang-dau-khi-da-phien.html"
	response, err := http.Get(BaseURL)
	checkerr(err)
	// fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Text())
	// Find the review items
	doc.Find("div.desktop").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find(".Title").Text()
		author := s.Find(".ReferenceSourceTG").Text()
		date := s.Find(".Date").Text()
		fmt.Println(title, author, date)
		strA := title + ", " + author + ", " + date
		dataBytes := []byte(strA)
		ioutil.WriteFile("example.txt", dataBytes, 0)
	})
	fmt.Println("DONE")
	response.Body.Close()
}
func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func (ebooks *Ebooks) GetTotalPages(url string) error {

// }
// func (newsList *NewsList) getAllNews(currentUrl string) error {

// }
// func (newsList *NewsList) getNewsByURL(url string) error {
// 	doc, err := goquery.NewDocument(url)
// 	if err != nil {
// 		return err
// 	}
// 	doc.Find("")
// 	return nil
// }
