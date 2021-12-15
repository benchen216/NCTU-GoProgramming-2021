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
	IP_USER_NUM   int
	KEYWORD_COUNT int
	KEYWORDS      []string
	Ptt           PTTArticles
	Fb            FBArticles
}

func (cw System) String() string {
	if len(os.Args) <= 3 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
	cw.IP_USER_NUM, _ = strconv.Atoi(os.Args[1])
	cw.KEYWORD_COUNT, _ = strconv.Atoi(os.Args[2])
	cw.KEYWORDS = os.Args[3:]
	a := cw.CountCyberWarriors()
	b := cw.CountKeyWord()
	return a + b
	//return "There's nothing here."
}

func (cw System) LoadPTT(url string) PTTArticles {
	var articles PTTArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	return articles
}

func (cw System) LoadFB(url string) FBArticles {
	var articles FBArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	return articles
}

func (cw System) CountCyberWarriors() string {
	ip_user := make(map[string]map[string]bool) //ip name
	for _, article := range cw.Ptt.Articles {
		if article.Ip == "None" {
			continue
		}
		if _, is_exist := ip_user[article.Ip]; is_exist == false {
			ip_user[article.Ip] = make(map[string]bool)
		}
		ip_user[article.Ip][article.Author] = true
	}
	var sorted_ip []string
	for ip, _ := range ip_user {
		sorted_ip = append(sorted_ip, ip)
	}
	sort.Strings(sorted_ip)
	ret := ""
	for _, ip := range sorted_ip {
		if len(ip_user[ip]) > cw.IP_USER_NUM {
			ret += ip + ", total: " + strconv.Itoa(len(ip_user[ip])) + "\n["
			var authors []string
			for author, _ := range ip_user[ip] {
				authors = append(authors, author)
			}
			sort.Strings(authors)
			for i, author := range authors {
				ret += author
				if i != len(authors)-1 {
					ret += ", "
				} else {
					ret += "]\n"
				}
			}
		}
	}
	return ret
}

func (cw System) CountKeyWord() string {
	ret := ""
	for _, keyword := range cw.KEYWORDS {
		user_keyword := make(map[string]int)
		for _, article := range cw.Ptt.Articles {
			if cnt := strings.Count(article.Article_title, keyword); cnt > 0 {
				if _, is_exist := user_keyword[article.Author]; is_exist == true {
					user_keyword[article.Author] += 1
				} else {
					user_keyword[article.Author] = 1
				}
			}
		}
		for _, article := range cw.Fb.Articles {
			if cnt := strings.Count(article.Article_title, keyword); cnt > 0 {
				if _, is_exist := user_keyword[article.Author]; is_exist == true {
					user_keyword[article.Author] += 1
				} else {
					user_keyword[article.Author] = 1
				}
			}
		}
		var sorted_users []string
		for user, cnt := range user_keyword {
			if cnt > cw.KEYWORD_COUNT {
				sorted_users = append(sorted_users, user)
			}
		}
		sort.Strings(sorted_users)
		ret += keyword + ", total: " + strconv.Itoa(len(sorted_users)) + "\n["
		for i, user := range sorted_users {
			ret += user
			if i != len(sorted_users)-1 {
				ret += ", "
			}
		}
		ret += "]\n"
	}

	return ret
}
