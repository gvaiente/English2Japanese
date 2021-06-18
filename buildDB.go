package main

import (
	"io/ioutil"
	"strings"
	"sync"
)

type word struct {
	key  string
	jp   string
	def  []string
	list []word
}

func BuildDB() map[string]*word {
	limit := make(chan int, 1)
	var words []word
	var wg sync.WaitGroup
	data, _ := ioutil.ReadFile("dst.txt")
	split := strings.Split(string(data), "\n")

	for i := 0; i < len(split)-2; i++ {
		wg.Add(1)
		limit <- 1
		go func() {
			defer wg.Done()
			w := new(word)
			s := strings.Split(string(split[i]), "\t")
			w.key = s[0]
			w.jp = s[1]
			defs := strings.Split(s[2], ";")
			for j := range defs {
				w.def = append(w.def, defs[j])
			}

			words = append(words, *w)
			//fmt.Println(words[len(words)-1])
			<-limit
		}()
	}
	wg.Wait()
	//fmt.Println(len(words))
	//fmt.Println(len(split))

	m := make(map[string]*word)

	for i := range words {
		if v, ok := m[words[i].key]; !ok {
			m[words[i].key] = &words[i]
		} else {
			v.list = append(v.list, words[i])
		}

	}
	return m
}
