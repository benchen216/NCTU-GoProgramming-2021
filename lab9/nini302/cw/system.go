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
	Ip     string
	Author string
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

func (s System) CountCyberWarriors(PTA PTTArticles, usr_num int) {
	mymap := make(map[System]int)

	for _, d := range PTA.Articles {
		mymap[System{d.Ip, d.Author}]++
	}
	my_map := make(map[string][]string)
	for i, _ := range mymap {
		if i.Author != "" {
			my_map[i.Ip] = append(my_map[i.Ip], i.Author)
		}
	}
	order := []string{}
	for i, j := range my_map {
		if len(j) > usr_num && i != "None" {
			order = append(order, i)
		}
	}
	sort.Strings(order)
	for _, i := range order {
		j := my_map[i]
		sort.Strings(j)
		fmt.Print(i, ", total: ", len(j), "\n")
		fmt.Print("[", strings.Join(j, ", "), "]")
		fmt.Println("")
	}
}

func (s System) CountKeyWord(PTA PTTArticles, FB FBArticles, key_num int) {
	for _, i := range os.Args[3:] {
		mymap := make(map[string]int)
		for _, e := range PTA.Articles {
			if strings.Contains(e.Article_title, i) {
				mymap[e.Author]++
			}
		}
		for _, e := range FB.Articles {
			if strings.Contains(e.Article_title, i) {
				mymap[e.Author]++
			}
		}
		names := []string{}
		for i, j := range mymap {
			if j > key_num {
				names = append(names, i)
			}
		}
		sort.Strings(names)
		fmt.Print(i, ",total:", len(names), "\n")
		fmt.Print("[", strings.Join(names, ","), "]")
		fmt.Println("")
	}

}
