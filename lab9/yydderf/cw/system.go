package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "sort"
    "strings"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

type System struct {
    ip_acc map[string]map[string]bool
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

func (s *System) CountCyberWarriors(ptt PTTArticles, num int) {
    s.ip_acc = map[string]map[string]bool{}
    for _, article := range ptt.Articles {
        tmp_ip := article.Ip
        tmp_author := article.Article.Author
        if s.ip_acc[tmp_ip] == nil {
            s.ip_acc[tmp_ip] = map[string]bool{}
        }
        if _, ok := s.ip_acc[tmp_ip][tmp_author]; !ok {
            s.ip_acc[tmp_ip][tmp_author] = true
        }
    }
    delete(s.ip_acc, "None")
    ips := make([]string, 0, len(s.ip_acc))
    for ip := range s.ip_acc {
        ips = append(ips, ip)
    }
    sort.Strings(ips)
    for i := 0; i < len(s.ip_acc); i++ {
        ip := ips[i]
        authors := s.ip_acc[ips[i]]
        if len(authors) > num {
            author_list := make([]string, 0, len(authors))
            fmt.Printf("%s, total: %d\n", ip, len(authors))
            for k, _ := range authors {
                author_list = append(author_list, k)
            }
            sort.Strings(author_list)
            output := "[" + strings.Join(author_list, `, `) + "]"
            fmt.Println(output)
        }
    }
    return
}

func (s *System) CountKeyWord(ptt PTTArticles, fb FBArticles, cnt int, keywords []string){
    // db establishment
    data := map[string][]string{}
    var count int
    for _, article := range ptt.Articles {
        tmp_author := article.Article.Author
        tmp_title := article.Article.Article_title
        data[tmp_author] = append(data[tmp_author], tmp_title)
    }
    for _, article := range fb.Articles {
        tmp_author := article.Article.Author
        tmp_title := article.Article.Article_title
        data[tmp_author] = append(data[tmp_author], tmp_title)
    }
    // loop through all the keywords
    for _, keyword := range keywords {
        usr_art := []string{}
        for usr, articles := range data {
            count = 0
            for _, article := range articles {
                if strings.Contains(article, keyword) == true {
                    count += 1
                }
            }
            if count > cnt {
                usr_art = append(usr_art, usr)
            }
        }
        sort.Strings(usr_art)
        fmt.Printf("%s, total: %d\n", keyword, len(usr_art))
        output := "[" + strings.Join(usr_art, `, `) + "]"
        fmt.Println(output)
    }
    return
}
