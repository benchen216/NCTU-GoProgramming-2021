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

type Users struct {
	n     int
	users []string
}

func (System) CountCyberWarriors(ptt PTTArticles) {
	n, _ := strconv.Atoi(os.Args[1])
	ipM := make(map[string]Users)
	var ips []string

	for i := 0; i < len(ptt.Articles); i++ {
		if ptt.Articles[i].Ip != "None" && ptt.Articles[i].Author != "" {
			tmp := ipM[ptt.Articles[i].Ip]

			flag := 0
			for j := 0; j < tmp.n; j++ {
				if ptt.Articles[i].Author == tmp.users[j] {
					flag = 1
					break
				}
			}

			if flag == 0 {
				tmp.n++
				tmp.users = append(tmp.users, ptt.Articles[i].Author)
				ipM[ptt.Articles[i].Ip] = tmp

				if tmp.n == n+1 {
					ips = append(ips, ptt.Articles[i].Ip)
				}
			}
		}
	}

	sort.Strings(ips)

	for i := 0; i < len(ips); i++ {
		tmp := ipM[ips[i]]
		fmt.Printf("%s, total: %d\n", ips[i], tmp.n)
		sort.Strings(tmp.users)
		fmt.Printf("[")
		for j := 0; j < len(tmp.users); j++ {
			if j != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", tmp.users[j])
		}
		fmt.Printf("]\n")
	}
}

func (System) CountKeyWord(ptt PTTArticles, fb FBArticles) {
	n, _ := strconv.Atoi(os.Args[2])
	keyword := os.Args[3:]
	Authors := make(map[string]map[string]int)
	var authors [][]string
	for i := 0; i < len(keyword); i++ {
		authors = append(authors, make([]string, 0))
	}

	for i := 0; i < len(ptt.Articles); i++ {
		for j := 0; j < len(keyword); j++ {
			if strings.Contains(ptt.Articles[i].Article_title, keyword[j]) {
				var tmp []int

				for k := 0; k < len(keyword); k++ {
					tmp = append(tmp, Authors[ptt.Articles[i].Author][keyword[k]])
				}
				tmp[j]++
				Authors[ptt.Articles[i].Author] = make(map[string]int)
				for k := 0; k < len(keyword); k++ {
					Authors[ptt.Articles[i].Author][keyword[k]] = tmp[k]
				}

				if tmp[j] == n+1 {
					authors[j] = append(authors[j], ptt.Articles[i].Author)
				}
			}
		}
	}

	for i := 0; i < len(fb.Articles); i++ {
		for j := 0; j < len(keyword); j++ {
			if strings.Contains(fb.Articles[i].Article_title, keyword[j]) {
				var tmp []int

				for k := 0; k < len(keyword); k++ {
					tmp = append(tmp, Authors[fb.Articles[i].Author][keyword[k]])
				}
				tmp[j]++
				Authors[fb.Articles[i].Author] = make(map[string]int)
				for k := 0; k < len(keyword); k++ {
					Authors[fb.Articles[i].Author][keyword[k]] = tmp[k]
				}

				if tmp[j] == n+1 {
					authors[j] = append(authors[j], fb.Articles[i].Author)
				}
			}
		}
	}

	for i := 0; i < len(keyword); i++ {
		fmt.Printf("%s, total: %d\n", keyword[i], len(authors[i]))
		sort.Strings(authors[i])
		fmt.Printf("[")
		for j := 0; j < len(authors[i]); j++ {
			if j != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", authors[i][j])
		}
		fmt.Printf("]\n")
	}
}
