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

type data struct {
	ids []string
}

func (System) CountCyberWarriors(ptt PTTArticles) {
	lim, _ := strconv.Atoi(os.Args[1])
	ip_list := make(map[string]data)
	ans := []string{}
	for _, post := range ptt.Articles {
		if post.Ip == "None" {
			continue
		}
		usr := ip_list[post.Ip]
		chk := true
		for _, id := range usr.ids {
			if id == post.Author {
				chk = false
				break
			}
		}
		if chk {
			usr.ids = append(usr.ids, post.Author)
			ip_list[post.Ip] = usr
			if len(ip_list[post.Ip].ids) == lim+1 {
				ans = append(ans, post.Ip)
			}
		}
	}
	sort.Strings(ans)
	for _, val := range ans {
		fmt.Printf("%s, total: %d\n", val, len(ip_list[val].ids))
		sort.Strings(ip_list[val].ids)
		fmt.Printf("[%s", ip_list[val].ids[0])
		if len(ip_list[val].ids) > 1 {
			for _, id := range ip_list[val].ids[1:] {
				fmt.Printf(", %s", id)
			}
		}
		fmt.Printf("]\n")
	}
}

func (System) CountKeyWord(fb FBArticles, ptt PTTArticles) {
	lim, _ := strconv.Atoi(os.Args[2])
	keywords := os.Args[3:]
	statics := make(map[string]map[string]int)
	ans := make(map[string]data)
	for _, k := range keywords {
		statics[k] = make(map[string]int)
	}
	for _, post := range ptt.Articles {
		for _, keyw := range keywords {
			if strings.Contains(post.Article_title, keyw) {
				statics[keyw][post.Author] += 1
				if statics[keyw][post.Author] == lim+1 {
					tmp := ans[keyw]
					tmp.ids = append(tmp.ids, post.Author)
					ans[keyw] = tmp
				}
			}
		}
	}
	for _, post := range fb.Articles {
		for _, keyw := range keywords {
			if strings.Contains(post.Article_title, keyw) {
				statics[keyw][post.Author] += 1
				if statics[keyw][post.Author] == lim+1 {
					tmp := ans[keyw]
					tmp.ids = append(tmp.ids, post.Author)
					ans[keyw] = tmp
				}
			}
		}
	}
	for _, keyw := range keywords {
		fmt.Printf("%s, total: %d\n", keyw, len(ans[keyw].ids))
		sort.Strings(ans[keyw].ids)
		fmt.Printf("[")
		if len(ans[keyw].ids) > 0 {
			fmt.Printf("%s", ans[keyw].ids[0])
		}
		if len(ans[keyw].ids) > 1 {
			for _, id := range ans[keyw].ids[1:] {
				fmt.Printf(", %s", id)
			}
		}
		fmt.Printf("]\n")
	}
}
