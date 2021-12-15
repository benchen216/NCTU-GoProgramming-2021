package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (System) String() string {
	return "There's nothing here."
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

func (System) CountCyberWarriors(ptt PTTArticles, IP_USER_NUM int) string {
	ip_user := make(map[string]map[string]bool) //ip name
	for _, article := range ptt.Articles {
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
		if len(ip_user[ip]) > IP_USER_NUM {
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

func (System) CountKeyWord(ptt PTTArticles, fb FBArticles,
	KEYWORD_COUNT int, KEYWORDS []string) string {
	ret := ""
	for _, keyword := range KEYWORDS {
		user_keyword := make(map[string]int)
		for _, article := range ptt.Articles {
			if cnt := strings.Count(article.Article_title, keyword); cnt > 0 {
				if _, is_exist := user_keyword[article.Author]; is_exist == true {
					user_keyword[article.Author] += 1
				} else {
					user_keyword[article.Author] = 1
				}
			}
		}
		for _, article := range fb.Articles {
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
			if cnt > KEYWORD_COUNT {
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
