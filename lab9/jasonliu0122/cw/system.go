package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
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

type cyberwarrior struct{
	bool_map map[string] bool;
	id []string
}

func (System) CountCyberWarriors(pttArticles PTTArticles) {

	ip_user_num, _ := strconv.Atoi(os.Args[1])
	times := make(map[string]cyberwarrior)
	var list_ip []string

	for _,article := range pttArticles.Articles {
		if  article.Article.Author=="" || article.Ip == "None" {
			continue
		}
		cw_value,tf := times[article.Ip]
		if tf {
			_,tf2 := cw_value.bool_map[article.Article.Author]
			if !tf2 {
				cw_value.bool_map[article.Article.Author] = true
				cw_value.id = append(cw_value.id,article.Article.Author)
				times[article.Ip] = cw_value
			}
		}else{ 
			list_ip = append(list_ip,article.Ip)
			var tmp  cyberwarrior
			tmp.bool_map = make(map[string]bool)
			tmp.id = append(tmp.id,article.Article.Author)
			tmp.bool_map[article.Article.Author] = true
			times[article.Ip] = tmp
		}
	}

	sort.Strings(list_ip)

	for _,ip := range list_ip {
		if len(times[ip].id)>ip_user_num {
			sort.Strings(times[ip].id)
			fmt.Printf("%s, total: %d\n",ip,len(times[ip].id))
			for i,name := range times[ip].id {
				if i == 0{
					fmt.Printf("[%s",name)	
				}else {
					fmt.Printf(", %s",name)
				}
			}
			fmt.Printf("]\n")
		}
	}


}

func (System) CountKeyWord(pttArticles PTTArticles,fbArticles FBArticles){

	word_num, _ := strconv.Atoi(os.Args[2])

	for _,word := range os.Args[3:] {	
		times := make(map[string]int)
		list_name := []string{}
		for _,article := range pttArticles.Articles {
			if strings.Contains(article.Article.Article_title,word){
				times[article.Article.Author]++
			}
		}
		for _,article := range fbArticles.Articles {
			if strings.Contains(article.Article.Article_title,word){
				times[article.Article.Author]++
			}
		}
		for name,num := range times {
			if num>word_num{
				list_name = append(list_name,name)
			}
		}
		fmt.Printf("%s, total: %d\n",word,len(list_name))
		sort.Strings(list_name)
		fmt.Printf("[")
		for i,name := range list_name {
			if i == 0 {
				fmt.Printf("%s",name)	
			}else {
				fmt.Printf(", %s",name)
			}
		}
		fmt.Printf("]\n")
	}

}

