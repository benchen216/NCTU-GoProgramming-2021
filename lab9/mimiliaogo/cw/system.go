package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"sort"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

type System struct {
	// you can add some data type if you like
}

func (System) String() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
	// return "There's nothing here."
}

func (System) LoadPTT(url string) PTTArticles {
	var articles PTTArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	return articles
}

func (System) LoadFB(url string) FBArticles {
	var articles FBArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	return articles
}


func (System) CountCyberWarriors(ptt_articles PTTArticles) {
	cyber_num, _ := strconv.Atoi(os.Args[1])
	var data = map[string]map[string]bool{}
	for _, v := range ptt_articles.Articles {
		if v.Author != "" {
			if _, ok := data[v.Ip]; ok { 
				data[v.Ip][v.Author] = true
			} else {
				data[v.Ip] = make(map[string]bool)
				data[v.Ip][v.Author] = true
			}
		}
	}
	// sort ip
	ips := []string{}
	for ip, _ := range data {
		ips = append(ips, ip)
	}
	sort.Strings(ips)
	for _, ip := range ips {
		author_map := data[ip]
		if (len(author_map) > cyber_num) {
			if ip != "None" {
				fmt.Printf("%s, total: %d\n", ip, len(author_map))
				authors := []string{}
				for author, _ := range(author_map) {
					authors = append(authors, author)
				}
				sort.Strings(authors)
				fmt.Println("[" + strings.Join(authors, `, `) + "]")
			}
		}
	}
	// for ip, author_map := range data {
	// 	if (len(author_map) > cyber_num) {
	// 		if ip != "None" {
	// 			fmt.Printf("%q, total: %d\n", ip, len(author_map))
	// 			authors := []string{}
	// 			for author, _ := range(author_map) {
	// 				authors = append(authors, author)
	// 			}
	// 			sort.Strings(authors)
	// 			fmt.Println(authors)
	// 		}
	// 	}
	// }

}

func (System) CountKeyWord(ptt_articles PTTArticles, fb_articles FBArticles){
	keyword_num, _ := strconv.Atoi(os.Args[2])
	keywords := os.Args[3:]
	// same author talk the keyword > keyword_num

	var data = map[string]map[string]int{}
	// [蔡英文]{[author1: 3][author2: 12]....}
	for _, ptt := range ptt_articles.Articles {
		for _, keyword := range keywords {
			if strings.Contains(ptt.Article_title, keyword) {
				if _, exist := data[keyword]; exist {
					data[keyword][ptt.Author] = data[keyword][ptt.Author] + 1
				} else {
					data[keyword] = make(map[string]int)
					data[keyword][ptt.Author] = 1
				}
			}
		}
	}
	for _, fb := range fb_articles.Articles {
		for _, keyword := range keywords {
			if strings.Contains(fb.Article_title, keyword) {
				if _, exist := data[keyword]; exist {
					data[keyword][fb.Author] = data[keyword][fb.Author] + 1
				} else {
					data[keyword] = make(map[string]int)
					data[keyword][fb.Author] = 1
				}
			}
		}
	}

	// // sort keyword
	// keyword_sort := []string{}
	// for key, _ := range data {
	// 	keyword_sort = append(keyword_sort, key)
	// }
	// sort.Strings(keyword_sort)

	for _, keyword := range keywords {
		author_map := data[keyword]
		authors := []string{}
		for author, times := range author_map {
			if times > keyword_num {
				authors = append(authors, author)
			}
		}
		fmt.Printf("%s, total: %d\n", keyword, len(authors))
		sort.Strings(authors)
		fmt.Println("[" + strings.Join(authors, `, `) + "]")
	}
}

