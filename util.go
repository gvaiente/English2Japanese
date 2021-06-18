package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func learnWord(key string) {
	fmt.Printf("Looking up \"%s\"", key)
	var url string
	file, _ := os.OpenFile("dst.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	c := &http.Client{
		Timeout: 100 * time.Second,
	}
	url = fmt.Sprintf("https://tangorin.com/words?search=%s", key)
	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err2 := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err2)
	}
	section := doc.Find(".wordDefinition").Each(func(i int, s *goquery.Selection) {
		subSection := s.Find(".w-jp").Each(func(i int, ss *goquery.Selection) {
			z, exist := ss.Find("sup").Attr("title")
			if exist != true {
				z = ""
			}
			if strings.Compare(z, "very common") == 0 || strings.Compare(z, "common") == 0 {
				jword := strings.ToLower(ss.Find(".roma").First().Text())
				//kanji := strings.ReplaceAll(ss.Find("a").Text(), "Inflection", "")
				definition := s.Find(".w-def").Find("span").First().Text()
				out := fmt.Sprintf("%s\t%s\t%s\n", key, jword, definition)
				if _, err := file.WriteString(out); err != nil {
					log.Fatal(err)
				}
			}
		})
		_ = subSection
	})
	_ = section
	fmt.Printf("\"%s\" Added to the definitions database\n", key)
}
