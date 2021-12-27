package cw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

type System struct {
	// you can add some data type if you like
	ptt_articles PTTArticles
	fb_articles  FBArticles
}

func (System) String() string {
	return "There's nothing here."
}

func (s *System) LoadPTT(url string) PTTArticles {
	var articles PTTArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	s.ptt_articles = articles
	return articles
}

func (s *System) LoadFB(url string) FBArticles {
	var articles FBArticles
	jsonBlob, _ := ioutil.ReadFile(url)
	checkErr(json.Unmarshal(jsonBlob, &articles))
	s.fb_articles = articles
	return articles
}

func (s *System) CountCyberWarriors() {
	fmt.Printf("%s\n", s.ptt_articles.Articles[0].Author)
	// for i, a := range articles {
	// 	fmt.Printf("%+v\n", a)
	// }
}

/* [TODO]
func (System) CountKeyWord(){

}
*/
