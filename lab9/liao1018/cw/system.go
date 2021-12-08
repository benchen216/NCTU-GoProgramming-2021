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
}

func (s System) String() string {
	if len(os.Args) < 4 {
        fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
        os.Exit(1)
    }
	PTTArticles := s.LoadPTT("./data/ptt.json")
	FBArticles := s.LoadFB("./data/fb.json")

	userNumber, _ := strconv.Atoi(os.Args[1])
	keyNumber, _ := strconv.Atoi(os.Args[2])

	s.CountCyberWarriors(PTTArticles, userNumber)
	s.CountKeyWord(PTTArticles, FBArticles, keyNumber)

	return ""
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

type user struct {
	Ip     string
	Author string
}

func (System) CountCyberWarriors(ptt PTTArticles, userNumber int) {
	countUser := make(map[user]int)
	for _, d := range ptt.Articles {
		countUser[user{d.Ip, d.Author}]++
	}
	AuthorsSameIp := make(map[string][]string)
	for i, _ := range countUser {
		if i.Author != "" {
			AuthorsSameIp[i.Ip] = append(AuthorsSameIp[i.Ip], i.Author)
		}
	}
	IPs := []string{}
	for i, j := range AuthorsSameIp {
		if len(j) > userNumber && i != "None" {
			IPs = append(IPs, i)
		}
	}
	sort.Strings(IPs)
	for _, ip := range IPs {
		authors := AuthorsSameIp[ip]
		sort.Strings(authors)
		fmt.Print(ip, ", total: ", len(authors), "\n")
		fmt.Print("[", strings.Join(authors, ", "), "]")
		fmt.Println("")
	}
}

func (s System) CountKeyWord(ptt PTTArticles, fb FBArticles, keyNumber int) {
	for _, key := range os.Args[3:] {
		count := make(map[string]int)
		for _, d := range ptt.Articles {
			if strings.Contains(d.Article_title, key) {
				count[d.Author]++
			}
		}
		for _, d := range fb.Articles {
			if strings.Contains(d.Article_title, key) {
				count[d.Author]++
			}
		}
		authors := []string{}
		for key, j := range count {
			if j > keyNumber {
				authors = append(authors, key)
			}
		}
		sort.Strings(authors)
		fmt.Print(key, ", total: ", len(authors), "\n")
		fmt.Print("[", strings.Join(authors, ", "), "]")
		fmt.Println("")
	}
}

