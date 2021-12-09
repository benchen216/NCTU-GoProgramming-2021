package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

type System struct {
	// you can add some data type if you like
	//<IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>
	IP_USER_NUM   int
	KEYWORD_COUNT int
	keywords      []string
}

func (System) String() string {
	return "There's nothing here."
}

func New() System {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}

	sys := System{}
	sys.IP_USER_NUM, _ = strconv.Atoi(os.Args[1])
	sys.KEYWORD_COUNT, _ = strconv.Atoi(os.Args[2])
	for _, keyword := range os.Args[3:] {
		sys.keywords = append(sys.keywords, keyword)
	}

	return sys
}

func (System) Check() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
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

func (sys System) CountCyberWarriors(pptArticles PTTArticles) {
	count := make(map[string][]string)
	for _, article := range pptArticles.Articles {
		if article.Ip == "None" || article.Author == "" {
			continue
		}

		var hasId bool = false
		for _, id := range count[article.Ip] {
			if id == article.Author {
				hasId = true
				break
			}
		}

		if !hasId {
			ids := count[article.Ip]
			ids = append(ids, article.Author)
			count[article.Ip] = ids
		}
	}

	var ips []string
	for ip, ids := range count {
		if len(ids) > sys.IP_USER_NUM {
			ips = append(ips, ip)
		}
	}

	sort.Strings(ips)

	for _, ip := range ips {
		sort.Strings(count[ip])
		fmt.Printf("%s, total: %d\n", ip, len(count[ip]))
		for i, id := range count[ip] {
			if i == 0 {
				fmt.Printf("[%s", id)
			} else {
				fmt.Printf(", %s", id)
			}
		}
		fmt.Printf("]\n")
	}
}

func (sys System) CountKeyWord(pttArticles PTTArticles, fbArticles FBArticles) {
	for _, keyword := range sys.keywords {
		count := make(map[string]int)

		for _, article := range pttArticles.Articles {
			if article.Author == "" {
				continue
			}

			if strings.Contains(article.Article_title, keyword) {
				count[article.Author]++
			}
		}

		for _, article := range fbArticles.Articles {
			if article.Author == "" {
				continue
			}

			if strings.Contains(article.Article_title, keyword) {
				count[article.Author]++
			}
		}

		var ids []string
		for id, i := range count {
			if i > sys.KEYWORD_COUNT {
				ids = append(ids, id)
			}
		}

		sort.Strings(ids)

		fmt.Printf("%s, total: %d\n", keyword, len(ids))
		fmt.Printf("[")
		for i, id := range ids {
			if i == 0 {
				fmt.Printf("%s", id)
			} else {
				fmt.Printf(", %s", id)
			}
		}
		fmt.Printf("]\n")
	}
}
