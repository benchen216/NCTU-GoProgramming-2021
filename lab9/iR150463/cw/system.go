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
	PTTCyberWarriors []PTTAccount
}

func (System) String() string {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
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

// help CountCyberWarriors
func (s System) AddAccount(Article PTTArticle) System {
	newIp := Article.Ip
	newAuthor := Article.Author

	for ip := range s.PTTCyberWarriors {
		if s.PTTCyberWarriors[ip].Account.Ip == newIp {
			foundAccount := false

			for ac := range s.PTTCyberWarriors[ip].Accounts {
				if s.PTTCyberWarriors[ip].Accounts[ac] == newAuthor {
					foundAccount = true
				}
			}

			if !foundAccount {
				s.PTTCyberWarriors[ip].Accounts = append(s.PTTCyberWarriors[ip].Accounts, newAuthor)
			}

			return s
		}
	}

	s.PTTCyberWarriors = append(s.PTTCyberWarriors, PTTAccount{Account{newIp}, []string{newAuthor}})
	return s
}

func (s System) CountCyberWarriors(Articles PTTArticles) {
	IP_USER_NUM, _ := strconv.Atoi(os.Args[1])

	for i := range Articles.Articles {
		if Articles.Articles[i].Ip == "None" || Articles.Articles[i].Article.Author == "" {
			continue
		}

		s = s.AddAccount(Articles.Articles[i])
	}

	sort.Slice(s.PTTCyberWarriors, func(i, j int) bool {
		return s.PTTCyberWarriors[i].Ip < s.PTTCyberWarriors[j].Ip
	})

	for i := range s.PTTCyberWarriors {
		accountsLen := len(s.PTTCyberWarriors[i].Accounts)

		if accountsLen > IP_USER_NUM {
			sort.Strings(s.PTTCyberWarriors[i].Accounts)

			fmt.Printf("%+v, total: %+v\n", s.PTTCyberWarriors[i].Ip, accountsLen)
			fmt.Print("[", strings.Join(s.PTTCyberWarriors[i].Accounts, ", "), "]\n")
		}
	}
}

func (System) CountKeyWord(FB FBArticles, PTT PTTArticles) {
	KEYWORD_COUNT, _ := strconv.Atoi(os.Args[2])

	for _, keyword := range os.Args[3:] {

		count := make(map[string]int)

		for _, article := range PTT.Articles {
			if strings.Contains(article.Article.Article_title, keyword) {
				count[article.Article.Author] += 1
			}
		}

		for _, article := range FB.Articles {
			if strings.Contains(article.Article.Article_title, keyword) {
				count[article.Article.Author] += 1
			}
		}

		authors := []string{}

		for name, count := range count {
			if count > KEYWORD_COUNT {
				authors = append(authors, name)
			}
		}

		fmt.Printf("%s, total: %d\n", keyword, len(authors))
		sort.Strings(authors)
		fmt.Printf("[")
		for i, name := range authors {
			if i == 0 {
				fmt.Printf("%s", name)
			} else {
				fmt.Printf(", %s", name)
			}
		}
		fmt.Printf("]\n")
	}
}
