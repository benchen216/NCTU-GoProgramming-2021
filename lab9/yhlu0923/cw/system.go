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

// TODO
func (System) CountCyberWarriors(PTT_Articles PTTArticles, FB_Articles FBArticles, IP_USER_NUM int) {
	/* Save the data in my map */
	ip_times := make(map[string][]string)
	for i := range PTT_Articles.Articles {
		ip := PTT_Articles.Articles[i].Ip
		if PTT_Articles.Articles[i].Article.Author == "" {
			continue
		}
		tmp_slice := []string{PTT_Articles.Articles[i].Article.Author}
		value, isExist := ip_times[ip]
		if !isExist {
			ip_times[ip] = tmp_slice
		} else {
			flag := false
			for j := range value {
				if value[j] == tmp_slice[0] {
					flag = true
					break
				}
			}
			if !flag {
				ip_times[ip] = append(value, tmp_slice[0])
			}
		}
	}

	/* Sort by Keys */
	keys := make([]string, 0, len(ip_times))
	for k := range ip_times {
		if len(ip_times[k]) > IP_USER_NUM && k != "None" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	// fmt.Printf("%+v\n", keys)

	/* Print Format */
	for _, k := range keys {

		if len(ip_times[k]) > IP_USER_NUM {
			fmt.Printf("%s, total: %d\n", k, len(ip_times[k]))
			sort.Strings(ip_times[k])
			// fmt.Println(k, ip_times[k])
			fmt.Printf("[")
			for i, name := range ip_times[k] {
				if i == 0 {
					fmt.Printf("%s", name)
				} else {
					fmt.Printf(", %s", name)
				}
			}
			fmt.Printf("]\n")
		}
	}
}

func (System) CountKeyWord(PTT_Articles PTTArticles, FB_Articles FBArticles, KEYWORD_COUNT int, KEYWORDS []string) {

	answer := make(map[string]map[string]int)
	// var answer []map[string]int
	for _, keyword := range KEYWORDS {
		/* Save the data in my map */
		author_times := make(map[string]int)
		for i := range PTT_Articles.Articles {
			title := PTT_Articles.Articles[i].Article.Article_title
			author := PTT_Articles.Articles[i].Article.Author

			if strings.Contains(title, keyword) {
				value, isExist := author_times[author]
				if !isExist {
					author_times[author] = 1
				} else {
					author_times[author] = value + 1
				}
			}
		}

		for i := range FB_Articles.Articles {
			title := FB_Articles.Articles[i].Article.Article_title
			author := FB_Articles.Articles[i].Article.Author

			if strings.Contains(title, keyword) {
				value, isExist := author_times[author]
				if !isExist {
					author_times[author] = 1
				} else {
					author_times[author] = value + 1
				}
			}
		}

		answer[keyword] = author_times
	}

	for _, keyword := range KEYWORDS {
		/* Sort by Keys */
		keys := make([]string, 0, len(answer[keyword]))
		for k := range answer[keyword] {
			if answer[keyword][k] > KEYWORD_COUNT {
				keys = append(keys, k)
			}
		}
		sort.Strings(keys)
		// fmt.Printf("%+v\n", keys)

		fmt.Printf("%s, total: %d\n", keyword, len(keys))

		/* Print Format */
		fmt.Printf("[")
		for i, name := range keys {
			if i == 0 {
				fmt.Printf("%s", name)
			} else {
				fmt.Printf(", %s", name)
			}
		}
		fmt.Printf("]\n")
	}
}
