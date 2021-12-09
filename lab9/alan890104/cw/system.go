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

type System struct {
	// you can add some data type if you like
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

type pair struct {
	Ip     string
	Author string
}

func (s System) String() string {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
	PTTArticles := s.LoadPTT("./data/ptt.json")
	FBArticles := s.LoadFB("./data/fb.json")

	usr_num, _ := strconv.Atoi(os.Args[1])
	key_num, _ := strconv.Atoi(os.Args[2])

	s.CountCyberWarriors(PTTArticles, usr_num)
	s.CountKeyWord(PTTArticles, FBArticles, key_num)

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

func (System) CountCyberWarriors(ptt PTTArticles, usr_num int) {
	mymap := make(map[pair]int)
	for _, d := range ptt.Articles {
		mymap[pair{d.Ip, d.Author}]++
	}
	mymap_2 := make(map[string][]string)
	for i, _ := range mymap {
		if i.Author != "" {
			mymap_2[i.Ip] = append(mymap_2[i.Ip], i.Author)
		}
	}
	order := []string{}
	for i, j := range mymap_2 {
		if len(j) > usr_num && i != "None" {
			order = append(order, i)
		}
	}
	sort.Strings(order)
	for _, i := range order {
		j := mymap_2[i]
		sort.Strings(j)
		fmt.Print(i, ", total: ", len(j), "\n")
		fmt.Print("[", strings.Join(j, ", "), "]")
		fmt.Println("")
	}
}

func (s System) CountKeyWord(ptt PTTArticles, fb FBArticles, key_num int) {
	for _, i := range os.Args[3:] {
		mymap := make(map[string]int)
		for _, d := range ptt.Articles {
			if strings.Contains(d.Article_title, i) {
				mymap[d.Author]++
			}
		}
		for _, d := range fb.Articles {
			if strings.Contains(d.Article_title, i) {
				mymap[d.Author]++
			}
		}
		names := []string{}
		for i, j := range mymap {
			if j > key_num {
				names = append(names, i)
			}
		}
		sort.Strings(names)
		fmt.Print(i, ", total: ", len(names), "\n")
		fmt.Print("[", strings.Join(names, ", "), "]")
		fmt.Println("")
	}
}
