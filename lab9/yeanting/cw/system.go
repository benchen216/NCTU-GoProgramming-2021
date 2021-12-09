package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sort"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

type System struct {
	// you can add some data type if you like
	IP_count map[string][]string
	Keyw_count map[string][]string
	PTT PTTArticles
	FB FBArticles
	IP_USER_NUM int
	KEYWORD_COUNT int
	KEYWORD []string
}

func (sys System) String() string {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
	sys.IP_USER_NUM, _ = strconv.Atoi(os.Args[1])
	sys.KEYWORD_COUNT, _ = strconv.Atoi(os.Args[2])
	for idx, arg := range os.Args {
		if idx > 2 {
			sys.KEYWORD = append(sys.KEYWORD, arg)
		}
	}
	rv1 := sys.CountCyberWarriors()
	rv2 := sys.CountKeyWord()
	var rv string
	for _, ele := range rv1 {
		rv += ele
	}
	for _, ele := range rv2 {
		rv += ele
	}
	rv = strings.TrimSuffix(rv, "\n")
	return rv
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

func (sys System) CountCyberWarriors() []string {
	for _, ele := range sys.PTT.Articles {
		exist := 0
		for _, val := range sys.IP_count[ele.Ip] {
			if val == ele.Article.Author {
				exist = 1
				break
			}
		}
		if exist == 0 {
			sys.IP_count[ele.Ip] = append(sys.IP_count[ele.Ip], ele.Article.Author)
		}
	}
	var rv []string  // return value
	for key, val := range sys.IP_count {
		if key == "None"{
			continue
		}
		if len(val) > sys.IP_USER_NUM {
			sort.Strings(val)
			cnt := strconv.Itoa(len(val))
			rv = append(rv, 
				key + ", total: " + cnt + "\n" + 
				"[" + strings.Join(val, ", ") + "]" + "\n")
		}
	}
	return rv
}

func (sys System) CountKeyWord() []string {
	var rv []string  // return value
	for _, ele := range sys.KEYWORD {
		a_cnt := make(map[string]int)  // author say keyword count
		for _, val := range sys.PTT.Articles {
			if strings.Contains(val.Article_title, ele) {
				a_cnt[val.Author] += 1
			}
		}
		for _, val := range sys.FB.Articles {
			if strings.Contains(val.Article_title, ele) {
				a_cnt[val.Author] += 1
			}
		}
		for key, val := range a_cnt {
			if val > sys.KEYWORD_COUNT {
				sys.Keyw_count[ele] = append(sys.Keyw_count[ele], key)
			}
		}
		sort.Strings(sys.Keyw_count[ele])
		cnt := strconv.Itoa(len(sys.Keyw_count[ele]))
		rv = append(rv, 
			ele + ", total: " + cnt + "\n" + 
			"[" + strings.Join(sys.Keyw_count[ele], ", ") + "]" + "\n")
	}
	return rv
}
