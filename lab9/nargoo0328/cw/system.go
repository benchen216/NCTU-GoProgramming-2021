package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"os"
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

/* [TODO]
func (System) CountCyberWarriors() {

}

func (System) CountKeyWord(){

}
*/
type CyberWarrior struct{
	check_map map[string]bool
	ids []string
}
func (System) CountCyberWarriors(articles PTTArticles){
	threshold, _ := strconv.Atoi(os.Args[1])
	count := make(map[string]CyberWarrior)
	ip_list := []string{}
	for _,article := range articles.Articles{
		if article.Ip=="None"{
			continue
		}
		cw_value,ok := count[article.Ip]
		if ok{
			_,ok2 := cw_value.check_map[article.Article.Author]
			if !ok2{
				cw_value.check_map[article.Article.Author] = true
				cw_value.ids = append(cw_value.ids,article.Article.Author)
				count[article.Ip] = cw_value
			}
		}else{
			ip_list = append(ip_list,article.Ip)
			temp := CyberWarrior{check_map: make(map[string]bool) ,ids: []string{article.Article.Author}}
			temp.check_map[article.Article.Author] = true
			count[article.Ip] = temp
		}
	}
	sort.Strings(ip_list)
	for _,ip := range ip_list{
		if len(count[ip].ids)>threshold{
			sort.Strings(count[ip].ids)
			fmt.Printf("%s, total: %d\n",ip,len(count[ip].ids))
			fmt.Printf("[")
			for i,name := range count[ip].ids{
				if i==0{
					fmt.Printf("%s",name)	
				}else{
					fmt.Printf(", %s",name)
				}
			}
			fmt.Printf("]\n")
		}
	}
}
func (System) CountKeyWord(PTT PTTArticles, FB FBArticles){
	threshold, _ := strconv.Atoi(os.Args[2])
	for _,keyword := range os.Args[3:]{	
		count := make(map[string]int)
		name_list := []string{}
		for _,article := range PTT.Articles{
			if strings.Contains(article.Article.Article_title,keyword){
				count[article.Article.Author]+=1
			}
		}
		for _,article := range FB.Articles{
			if strings.Contains(article.Article.Article_title,keyword){
				count[article.Article.Author]+=1
			}
		}
		for name,number := range count{
			if number>threshold{
				name_list = append(name_list,name)
			}
		}
		fmt.Printf("%s, total: %d\n",keyword,len(name_list))
		sort.Strings(name_list)
		fmt.Printf("[")
		for i,name := range name_list{
			if i==0{
				fmt.Printf("%s",name)	
			}else{
				fmt.Printf(", %s",name)
			}
		}
		fmt.Printf("]\n")
	}
}