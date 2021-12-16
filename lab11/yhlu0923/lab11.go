package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var maxFlag int
var wFlag string

func init() {
	// Define the other flags here
	// flag.IntVar()
	flag.IntVar(&maxFlag, "max", 10, "Max Printing")
	flag.StringVar(&wFlag, "w", "ptt", "Web page")
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	c := colly.NewCollector()

	var names []string
	var contents []string
	var dates []string

	if wFlag == "ptt" {

		c.OnHTML("div.push > span.f3.hl.push-userid", func(e *colly.HTMLElement) {
			// fmt.Println("名字" + e.Text)
			names = append(names, strings.Trim(e.Text, " "))
		})

		c.OnHTML("div.push > span.f3.push-content", func(e *colly.HTMLElement) {
			// fmt.Println("留言" + e.Text)
			contents = append(contents, e.Text)
		})

		c.OnHTML("div.push > span.push-ipdatetime", func(e *colly.HTMLElement) {
			// fmt.Println("時間" + e.Text)
			dates = append(dates, e.Text)
		})

		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else {

		// c.OnHTML("td.teacherInfo > span.content_title2 ", func(e *colly.HTMLElement) {
		// prePath := "div#tab1.content_tab.tab-pane.active > table.teacherTableStyle > tbody > tr >"
		// c.OnHTML(prePath+"td.teacherInfo > span.content_title2 > a", func(e *colly.HTMLElement) {
		count := 0

		c.OnHTML("div#tab1.content_tab.tab-pane.active > table.teacherTableStyle > tbody > tr > td.teacherInfo", func(e *colly.HTMLElement) {
			data_name := e.ChildText("span.content_title2 > a")
			// fmt.Println(data_name)
			data_web := e.ChildAttr(`p.infop > a`, "href")
			// fmt.Println(data_web)
			if data_web == "" {
				data_web = "NULL"
			}

			count += 1
			s := fmt.Sprintf("%d. 姓名: %s, 網站: %s\n", count, data_name, data_web)
			contents = append(contents, s)
		})

		// c.OnHTML(prePath+"td.teacherInfo > p.infop > a", func(e *colly.HTMLElement) {
		// 	// fmt.Println("網站: " + e.Attr("href"))
		// 	website := strings.Trim(e.Attr("href"), " ")
		// 	if website == "" {
		// 		contents = append(contents, "NULL")
		// 	} else {
		// 		contents = append(contents, website)
		// 	}
		// })
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}
	c.Wait()

	if wFlag == "ptt" {

		for i := 0; i < maxFlag; i++ {
			fmt.Printf("%d. 名字: %s, 留言%s, 時間:%s", i+1, names[i], contents[i], dates[i])
		}
	} else {

		maxFlag = min(maxFlag, len(contents))
		for i := 0; i < maxFlag; i++ {
			// fmt.Printf("%d. 姓名: %s, 網站: %s\n", i+1, names[i], contents[i])
			fmt.Printf("%s", contents[i])
		}

		// fmt.Printf("Teacher num: %d, Website num: %d\n", len(names), len(contents))
	}
}
