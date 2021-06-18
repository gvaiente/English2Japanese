package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const workerCount int = 2994

func main() {
	var url string
	var wg sync.WaitGroup
	sem := make(chan int, workerCount)
	defer close(sem)
	c := &http.Client{
		Timeout: 100 * time.Second,
	}
	for i := 1; i < len(os.Args); i++ {
		wg.Add(1)
		sem <- 1
		go func(i int) {
			defer wg.Done()
			if key := os.Args[i]; key != "" {
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
							fmt.Printf("%s\t%s\t%s\n", key, jword, definition)
						}
					})
					_ = subSection
				})
				_ = section
			}
			<-sem
		}(i)
		wg.Wait()

	}

}
