package main

import (
	"awesomeProject/pkg/crawler"
	"awesomeProject/pkg/crawler/spider"
	"awesomeProject/pkg/index"
	"fmt"
	"sort"
)

func main() {
	s := spider.New()
	var srch, leftP, rightP int
	var d string
	var g1 []crawler.Document
	var g2 []crawler.Document
	g1, _ = s.Scan("https://go.dev", 2)
	g2, _ = s.Scan("https://golang.org", 2)
	for i := range g2 {
		g1 = append(g1, g2[i])
	}

	//сортировка полученных значений по id
	sort.Slice(g1, func(i, j int) bool {
		return g1[i].ID < g1[j].ID
	})

	var ind map[string]index.Elem
	ind = make(map[string]index.Elem)
	ind = index.Indx(g1)

	fmt.Println("Enter word: ")
	_, err := fmt.Scan(&d)
	if err != nil {
		return
	}
	//Нахождение слова в описаниях ссылок с помошью бинарного поиска
	for i := range ind[d].Num {
		srch = ind[d].Num[i]
		leftP = 0
		rightP = len(g1) - 1
		for leftP <= rightP {
			var mid = int((leftP + rightP) / 2)
			var midVal = g1[mid].ID

			if midVal == srch {
				fmt.Println(g1[mid].URL)
				fmt.Println(g1[mid].Title)
				fmt.Println("--------")
				break
			} else if midVal < srch {
				leftP = mid + 1
			} else {
				rightP = mid - 1
			}
		}

	}
}
