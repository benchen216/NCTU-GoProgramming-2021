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

func contains(arr []string, target string) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}

func usercount(arr []string, target string) int {
	cnt := 0
	for _, a := range arr {
		if a == target {
			cnt++
		}
	}
	return cnt
}

func unique(arr []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, key := range arr {
		if key == "" {
			continue
		}
		if !keys[key] {

			list = append(list, key)
			keys[key] = true
		}
	}
	return list
}

type System struct {
	// you can add some data type if you like
	ptt              PTTArticles
	fb               FBArticles
	ip_map           map[string][]string
	ip_user_num      int
	keyword          map[int]string
	keyword_num      int
	keyword_user_map map[int][]string
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

func (s System) CountCyberWarriors() {
	//find accs by id
	theip := []string{}
	for i := 0; i < len(s.ptt.Articles); i++ {
		if !contains(s.ip_map[s.ptt.Articles[i].Ip], s.ptt.Articles[i].Author) {
			if s.ptt.Articles[i].Author == "" {
				continue
			}
			s.ip_map[s.ptt.Articles[i].Ip] = append(s.ip_map[s.ptt.Articles[i].Ip], s.ptt.Articles[i].Author)
			// theip = append(theip, s.ptt.Articles[i].Ip)
		}
	}
	delete(s.ip_map, "None")
	for keys := range s.ip_map {
		theip = append(theip, keys)
	}
	theip = unique(theip)
	sort.Strings(theip)
	theip = theip[:len(theip)-1]
	// fmt.Println(theip)
	for _, ip := range theip {
		arr := s.ip_map[ip]
		if len(arr) > s.ip_user_num {
			fmt.Printf("%s, total: %d\n", ip, len(arr))
			sort.Strings(arr)
			out := "[" + strings.Join(arr, ", ") + "]"
			fmt.Println(out)
		}
	}

}

func (s System) CountKeyWord() {
	//find fb
	s.keyword_user_map = make(map[int][]string)
	for i := 0; i < len(s.fb.Articles); i++ {
		for index, key := range s.keyword {

			if strings.Contains(s.fb.Articles[i].Article_title, string(key)) {
				s.keyword_user_map[index] = append(s.keyword_user_map[index], s.fb.Articles[i].Author)
			}
		}
	}
	//find ptt
	for i := 0; i < len(s.ptt.Articles); i++ {
		for index, key := range s.keyword {
			if strings.Contains(s.ptt.Articles[i].Article_title, string(key)) {
				s.keyword_user_map[index] = append(s.keyword_user_map[index], s.ptt.Articles[i].Author)
			}
		}
	}

	//check count

	for index := 0; index < len(s.keyword); index++ {
		key := s.keyword[index]
		names := unique(s.keyword_user_map[index])
		// fmt.Println(string(key))
		// sort.Strings(names)
		out := ""
		result_user := []string{}
		for _, name := range names {
			if usercount(s.keyword_user_map[index], name) > s.keyword_num {
				result_user = append(result_user, name)
			}
		}
		fmt.Printf("%s, total: %d\n", key, len(result_user))
		sort.Strings(result_user)
		out = "[" + strings.Join(result_user, ", ") + "]"
		fmt.Println(out)
	}
}

func (s System) String(ptt PTTArticles, fb FBArticles) {

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
	s.ptt = ptt
	s.fb = fb
	//ipusernum
	s.ip_user_num, _ = strconv.Atoi(os.Args[1])
	s.ip_map = make(map[string][]string)
	s.keyword_num, _ = strconv.Atoi(os.Args[2])
	s.keyword = make(map[int]string)
	for index, key := range os.Args[3:] {
		nameRune := []rune(key)
		s.keyword[index] = string(nameRune)
		//fmt.Println(index)
	}
	s.CountCyberWarriors()
	s.CountKeyWord()
}
