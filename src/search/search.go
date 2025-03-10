package search

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Status     int64    `json:"status"`
	Link       string   `json:"link"`
	Videos     []string `json:"videos"`
	ImageLinks []string `json:"images"`
	Links      []string `json:"links"`
	TextSize   float64  `json:"size"`
	Time       string   `json:"time"`
	Title      string   `json:"title"`
}

func NewResult() *SearchResult {
	result := SearchResult{}
	result.ImageLinks = make([]string, 0)
	result.Links = make([]string, 0)
	result.Videos = make([]string, 0)

	return &result
}

func ParsePage(link string, channel chan SearchResult) {
	result := NewResult()

	resp, err := http.Get(link)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	result.Status = int64(resp.StatusCode)
	result.Link = link

	if resp.StatusCode != 404 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		ht_text := ""
		// Находим элементы с классом "greeting" и получаем их текст
		doc.Find(".tl_article_content").Each(func(i int, s *goquery.Selection) {
			ht_text = s.Contents().Text()

			s.Find("a").Each(func(_ int, a *goquery.Selection) {
				href, exists := a.Attr("href")
				if exists {
					result.Links = append(result.Links, href)
				}
			})
		})

		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			src, exists := s.Attr("src")

			if exists {
				result.ImageLinks = append(result.ImageLinks, src)
			}
		})

		doc.Find("video").Each(func(i int, s *goquery.Selection) {
			src, exists := s.Attr("src")

			if exists {
				result.Videos = append(result.Videos, src)
			}
		})

		doc.Find("time").Each(func(i int, s *goquery.Selection) {
			result.Time = s.Text()
		})

		doc.Find("h1").Each(func(i int, s *goquery.Selection) {
			result.Title = s.Text()
		})

		//	fmt.Println(ht_text)

		//data, err := io.ReadAll(resp.Body)

		//	if err != nil {
		//		fmt.Println(err)
		//	}

		result.TextSize = float64(len(ht_text)) / 1024.0
	}

	channel <- *result
}

func GenerateLink(text string, i int, j int, k int) string {
	link := "https://telegra.ph/"
	formatted_day := fmt.Sprintf("%02d", j)
	formatted_month := fmt.Sprintf("%02d", i)
	formatted_part := fmt.Sprintf("%02d", k)

	date_text := ""

	if k != 0 {
		date_text = text + "-" + formatted_month + "-" + formatted_day + "-" + formatted_part
		link += date_text
	} else {
		date_text = text + "-" + formatted_month + "-" + formatted_day
		link += date_text
	}

	return link
}

func RunSearch(query string, channel chan SearchResult) {
	link_name := Translit(query)

	for i := 1; i <= 12; i++ {
		for j := 1; j <= 31; j++ {
			for k := 0; k <= 10; k++ {
				go ParsePage(GenerateLink(link_name, i, j, k), channel)
				time.Sleep(1 * time.Millisecond)
			}
		}
	}
}
