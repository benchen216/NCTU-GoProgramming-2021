package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"errors"
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
	IP_count int
	Word_count int
	Keywords []string
}

func (s *System) Init_value() bool {
	var err error
	defer func() {
		checkErr(err)
	}()
	
	if len(os.Args) < 4 {
		err = errors.New("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		return false
	}

	if s.IP_count, err = strconv.Atoi(os.Args[1]); err != nil {
		return false
	}
	
	if s.Word_count, err = strconv.Atoi(os.Args[2]); err != nil {
		return false
	}
	
	s.Keywords = os.Args[3:]
	
	return true
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

func (s System) CountCyberWarriors(data PTTArticles) {
	user := make(map[string] map[string] struct{})
	
	// read data
	for _, x := range(data.Articles) {
		if x.Ip == "None" || x.Article.Author == "" {
			continue
		}
		
		if val, ok := user[x.Ip]; ok {
			if _, ok := val[x.Article.Author]; !ok {
				val[x.Article.Author] = struct{}{}
			}
		} else {
			user[x.Ip] = make(map[string] struct{})
			user[x.Ip][x.Article.Author] = struct{}{}
		}
	}
	
	// filter data
	type pack struct {
		ip string
		name []string
	}
	
	result := []pack{}
	
	for i, x := range(user) {
		if len(x) > s.IP_count {
			names := make([]string, 0, len(x))
			for y, _ := range(x) {
				names = append(names, y)
			}
			
			result = append(result, pack{i, names})
		}
	}
	
	// sort data
	sort.Slice(result, func (l, r int) bool {
		return result[l].ip < result[r].ip
	})
	
	for _, x := range(result) {
		sort.Slice(x.name, func(l, r int) bool {
			return x.name[l] < x.name[r]
		})
	}
	
	// output result
	for _, x := range(result) {
		fmt.Printf("%v, total: %d\n[%v]\n", x.ip, len(x.name), strings.Join(x.name, ", "))
	}
}

func (s System) CountKeyWord(data_p PTTArticles, data_f FBArticles) {
	result := make([]map[string] int, len(s.Keywords))
	for i := 0; i < len(s.Keywords); i++ {
		result[i] = make(map[string] int)
	}
	
	// read data
	for _, x := range(data_p.Articles) {
		for i, key := range(s.Keywords) {
			if strings.Contains(x.Article.Article_title, key) {
				if x.Article.Author != "" {
					result[i][x.Article.Author]++					
				}
			}
		} 
	}
	
	for _, x := range(data_f.Articles) {
		for i, key := range(s.Keywords) {
			if strings.Contains(x.Article.Article_title, key) {
				if x.Article.Author != "" {
					result[i][x.Article.Author]++					
				}
			}
		}
	}
	
	// filter & output
	for i, x := range(result) {
		var user []string
		
		for y, n := range(x) {
			if n > s.Word_count {
				user = append(user, y)
			}
		}
		
		sort.Slice(user, func (l, r int) bool {
			return user[l] < user[r]
		})
		
		fmt.Printf("%v, total: %d\n[%v]\n", s.Keywords[i], len(user), strings.Join(user, ", "))
	}
}
