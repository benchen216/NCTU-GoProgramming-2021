package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
}

func (System) String() string {
    cwSystem := System{}

    if len(os.Args) < 4 {
        fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
        os.Exit(1)
    }

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
 	FBArticles := cwSystem.LoadFB("./data/fb.json")
	
    cnt, err := strconv.Atoi(os.Args[1])
	if err == nil {
    	cwSystem.CountCyberWarriors(PTTArticles, cnt)
    }
    
    cnt2, err := strconv.Atoi(os.Args[2])
	if err == nil {
    	cwSystem.CountKeyWord(PTTArticles, FBArticles, cnt2)
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

type pair struct {
    Ip string
    Author string
}

func (System) CountCyberWarriors(articles PTTArticles, times int) {
    mymap := make(map[pair]int)
    for _, d := range articles.Articles {
        mymap[pair{d.Ip, d.Author}]++
    }
    mymap2 := make(map[string][]string)
    for i, _ := range mymap {
        if i.Author != "" {
            mymap2[i.Ip] = append(mymap2[i.Ip], i.Author)
        }
    }
    order := []string{};
    for i, j := range mymap2 {
        if len(j) > times && i != "None" {
            order = append(order, i)
        }
    }
    sort.Strings(order)
    for _, i := range order {
        j := mymap2[i]
        sort.Strings(j)
        fmt.Print(i, ", total: ", len(j), "\n");
        fmt.Print("[", strings.Join(j, ", "), "]")
        fmt.Println("");
    }
}

func (System) CountKeyWord(pttarticles PTTArticles, fbarticles FBArticles, times int) {
    for _, i := range os.Args[3:] {
        mymap := make(map[string]int)
        for _, d := range pttarticles.Articles {
            if  strings.Contains(d.Article_title, i) {
                mymap[d.Author]++
            }
        }
        for _, d := range fbarticles.Articles {
            if  strings.Contains(d.Article_title, i) {
                mymap[d.Author]++
            }
        }
        names := []string{}
        for i, j := range mymap {
            if j > times {
                names = append(names, i)
            }
        }
        sort.Strings(names)
        fmt.Print(i, ", total: ", len(names), "\n")
        fmt.Print("[", strings.Join(names, ", "), "]")
        fmt.Println("");
    }
}
