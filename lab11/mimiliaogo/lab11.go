package main

import (
	"flag"
	"fmt"
	// "strings"
	"strconv"
	"github.com/gocolly/colly"
)

var max_num int;
var website string;
func init() {
	// Define the other flags here
	// flag.IntVar()
	flag.IntVar(&max_num, "max", 4, "help message for flagname")
	flag.StringVar(&website, "w", "ptt", "help message for flagname")
}

func main() {
	flag.Parse()
	// flag.PrintDefaults()

	c := colly.NewCollector()

	// c.OnHTML("div#topbar.bbs-content > a#logo", func(e *colly.HTMLElement) {
	// 	fmt.Println("demo1: " + e.Text)
	// })

	// c.OnHTML("div[class='bbs-content'] > a[id='logo']", func(e *colly.HTMLElement) {
	// 	fmt.Println("demo2: " + e.Text)
	// })

	// c.OnHTML(".f2", func(e *colly.HTMLElement) {
	// 	fmt.Println("demo3: " + e.Text)
	// })

	// c.OnHTML("#logo", func(e *colly.HTMLElement) {
	// 	fmt.Println("demo4: " + e.Attr("href"))
	// })

	if (website == "ptt") { 
		c.OnHTML("#main-content", func(e *colly.HTMLElement) {
			e.ForEach("div.push", func(index int, e *colly.HTMLElement) {
				// fmt.Println("demo mimi: " + e.Text)
				if (index < max_num) {
					name := e.ChildText(".push-userid")
					comment := e.ChildText(".push-content")
					time := e.ChildText(".push-ipdatetime")
					// strings.Replace(comment, ": ", "")
					// fmt.Println(strconv.Itoa(index) + " 名字: " + name + ", 留言") 
					fmt.Printf("%s. 名字: %s, 留言%s, 時間: %s\n", strconv.Itoa(index+1), name, comment, time)
				}
			})
		})
	} else if (website == "ncku") {
		c.OnHTML("#tab1 > table > tbody  ", func(e *colly.HTMLElement) {
			e.ForEach("td.teacherInfo", func(index int, e *colly.HTMLElement) {
				// fmt.Println(e.ChildAttr("a", "text"))
				if (index < max_num) {
					prof_name := e.ChildText("span.content_title2")
					link := e.ChildAttr("p > a", "href")
					if (link == "") {
						link = "NULL"
					}
					// fmt.Println(link, prof_name)
					fmt.Printf("%s. 姓名: %s, 網站: %s\n", strconv.Itoa(index+1), prof_name, link)
				}
			})
		})
		// 1. 姓名: 張燕光, 網站: http://cial.csie.ncku.edu.tw/
	}
	if (website == "ptt") {
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else if (website == "ncku") {
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}
	c.Wait()
}
