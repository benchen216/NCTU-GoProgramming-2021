package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
    "sort"
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
	id []string
}

func (System) CountCyberWarriors(ptt PTTArticles) {
	counts, _ := strconv.Atoi(os.Args[1])
	ip := make(map[string]data)
	ans := []string{}
	for _, post := range ptt.Articles {
		if post.Ip == "None" || post.Author == "" {
			continue
		}
		usr := ip[post.Ip]
		chk := true
		for _, id := range usr.id {
			if id == post.Author {
				chk = false
				break
			}
		}
		if chk {
			usr.id = append(usr.id, post.Author)
			ip[post.Ip] = usr
			if len(ip[post.Ip].id) == counts+1 {
				ans = append(ans, post.Ip)
			}
		}
	}
	sort.Strings(ans)
	for _, val := range ans {
		fmt.Printf("%s, total: %d\n", val, len(ip[val].id))
		sort.Strings(ip[val].id)
		fmt.Printf("[%s", ip[val].id[0])
		if len(ip[val].id) > 1 {
			for _, i := range ip[val].id[1:] {
				fmt.Printf(", %s", i)
			}
		}
		fmt.Printf("]\n")
	}
}

func (System) CountKeyWord(fb FBArticles, ptt PTTArticles) {
	counts, _ := strconv.Atoi(os.Args[2])
	keyword := os.Args[3:]
	datas := make(map[string]map[string]int)
	ans := make(map[string]data)
	for _, k := range keyword {
		datas[k] = make(map[string]int)
	}
	for _, post := range ptt.Articles {
		for _, k := range keyword {
			if strings.Contains(post.Article_title, k) {
				datas[k][post.Author] += 1
				if datas[k][post.Author] == counts+1 {
					tmp := ans[k]
					tmp.id = append(tmp.id, post.Author)
					ans[k] = tmp
				}
			}
		}
	}
	for _, post := range fb.Articles {
		for _, k := range keyword {
			if strings.Contains(post.Article_title, k) {
				datas[k][post.Author] += 1
				if datas[k][post.Author] == counts+1 {
					tmp := ans[k]
					tmp.id = append(tmp.id, post.Author)
					ans[k] = tmp
				}
			}
		}
	}
	for _, k := range keyword {
		fmt.Printf("%s, total: %d\n", k, len(ans[k].id))
		sort.Strings(ans[k].id)
		fmt.Printf("[")
		if len(ans[k].id) > 0 {
			fmt.Printf("%s", ans[k].id[0])
		}
		if len(ans[k].id) > 1 {
			for _, i := range ans[k].id[1:] {
				fmt.Printf(", %s", i)
			}
		}
		fmt.Printf("]\n")
	}
}
