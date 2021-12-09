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

func contains(strArr []string, str string) bool {
	for _, s := range strArr {
		if s == str {
			return true
		}
	}
	return false
}

func (System) CountCyberWarriors(ipUserNum int, pttArticles PTTArticles) string {
	ipAuthors := map[string][]string{}
	/* Classify all articles by IP */
	for _, article := range pttArticles.Articles {
		// Skip those whose ip is ""
		if len(article.Ip) == 0 || article.Ip == "None" || len(article.Author) == 0 {
			continue
		}

		// If key doesn't exist, create one
		if _, exist := ipAuthors[article.Ip]; exist == false {
			ipAuthors[article.Ip] = []string{}
		}

		/* If an IP already contains a username, the username won't be added again */
		if !contains(ipAuthors[article.Ip], article.Author) {
			ipAuthors[article.Ip] = append(ipAuthors[article.Ip], article.Author)
		}
	}

	/* Filter IP by number of users and create output string */
	var ips []string
	for ip, _ := range ipAuthors {
		ips = append(ips, ip)
	}
	sort.Strings(ips)
	result := ""
	for _, ip := range ips {
		authors := ipAuthors[ip]
		if len(authors) > ipUserNum {
			result += ip + ", total: " + strconv.Itoa(len(authors)) + "\n"
			result += "["
			sort.Strings(authors)
			for i, author := range authors {
				result += author
				if i != len(authors)-1 {
					result += ", "
				}
			}
			result += "]\n"
		}
	}
	return result
}

func (System) CountKeyWord(keywordCount int, keywords []string, pttArticles PTTArticles, fbArticles FBArticles) string {
	result := ""
	for _, keyword := range keywords {
		/* Combine PTT & FB articles */
		type Article struct {
			Author       string
			ArticleTitle string
		}
		var articles []Article
		for _, article := range pttArticles.Articles {
			articles = append(articles, Article{Author: article.Author, ArticleTitle: article.Article_title})
		}
		for _, article := range fbArticles.Articles {
			articles = append(articles, Article{Author: article.Author, ArticleTitle: article.Article_title})
		}

		/* */
		userCount := map[string]int{}
		for _, article := range articles {
			cnt := strings.Count(article.ArticleTitle, keyword)
			if cnt > 0 {
				/* If user does not exist, create one */
				// TODO: here, I only count the number of articles contain the keyword. If an article contains the same keyword which the number of it is more than one, I may get an error.
				if _, exist := userCount[article.Author]; exist == false {
					userCount[article.Author] = 1
				} else {
					userCount[article.Author] += 1
				}
			}
		}

		/*  */
		var users []string
		for user, count := range userCount {
			if count <= keywordCount {
				continue
			}
			users = append(users, user)
		}
		sort.Strings(users)
		result += keyword + ", total: " + strconv.Itoa(len(users)) + "\n["
		for i, user := range users {
			if i > 0 {
				result += ", "
			}
			result += user
		}
		result += "]\n"
	}
	return result
}
