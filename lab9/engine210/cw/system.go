package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	ptt_articles PTTArticles
	fb_articles  FBArticles
}

func (System) String() string {
	return "There's nothing here."
}

func (s *System) LoadPTT(url string) PTTArticles {
	var articles PTTArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	s.ptt_articles = articles
	return articles
}

func (s *System) LoadFB(url string) FBArticles {
	var articles FBArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	s.fb_articles = articles
	return articles
}

func (s *System) CountCyberWarriors(ip_user_num int) {
	// fmt.Printf("balballba%s\n", s.ptt_articles.Articles[0].Author)
	// fmt.Printf("balballba%s\n", s.ptt_articles.Articles[0].Ip)

	ip_map := make(map[string]map[string]struct{})

	for _, v := range s.ptt_articles.Articles {
		_, found := ip_map[v.Ip]
		if found {
			ip_map[v.Ip][v.Author] = struct{}{}
		} else {
			ip_map[v.Ip] = make(map[string]struct{})
			ip_map[v.Ip][v.Author] = struct{}{}
		}
	}
	var keys []string
	for k, _ := range ip_map {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := ip_map[k]
		if len(v) > ip_user_num && k != "None" {
			fmt.Printf("%s, total: %d\n", k, len(v))
			var l []string
			for kk, _ := range v {
				l = append(l, kk)
			}
			sort.Strings(l)
			output := "[" + strings.Join(l, `, `) + `]`
			fmt.Println(output)
		}
	}
}

func (s *System) CountKeyWord(keyword_count int, keywords []string) {
	keyword_map := make(map[string]map[string]int) //map[keyword]map[author]count
	for _, keyword := range keywords {
		keyword_map[keyword] = make(map[string]int)
	}

	for _, v := range s.ptt_articles.Articles {
		for _, keyword := range keywords {
			if strings.Contains(v.Article_title, keyword) {
				_, found := keyword_map[keyword][v.Author]
				if found {
					keyword_map[keyword][v.Author] += 1
				} else {
					keyword_map[keyword][v.Author] = 1
				}
			}
		}
	}
	for _, v := range s.fb_articles.Articles {
		for _, keyword := range keywords {
			if strings.Contains(v.Article_title, keyword) {
				_, found := keyword_map[keyword][v.Author]
				if found {
					keyword_map[keyword][v.Author] += 1
				} else {
					keyword_map[keyword][v.Author] = 1
				}
			}
		}
	}

	for k, v := range keyword_map {
		fmt.Printf("%s, ", k)
		var l []string
		for kk, vv := range v {
			if vv > keyword_count {
				l = append(l, kk)
			}
		}
		fmt.Printf("total: %d\n", len(l))
		sort.Strings(l)
		output := "[" + strings.Join(l, `, `) + `]`
		fmt.Printf("%s\n", output)
	}
}
