package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

var (
	max int
	web string
)

func init() {
	flag.IntVar(&max, "max", 10, "Max Printing")
	flag.StringVar(&web, "w", "ptt", "Web page")
}

func ppt(c *colly.Collector) {
	c.OnHTML("div#main-content", func(e *colly.HTMLElement) {
		e.ForEach("div.push", func(i int, e *colly.HTMLElement) {
			if i >= max {
				return
			}
			fmt.Printf("%d. 名字: %s,", i+1, e.ChildText(".push-userid"))
			fmt.Printf(" 留言%s,", e.ChildText(".push-content"))
			fmt.Printf(" 時間: %s\n", e.ChildText(".push-ipdatetime"))
		})
	})

	c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
}

func ncku(c *colly.Collector) {
	c.OnHTML("div#tab1", func(e *colly.HTMLElement) {
		e.ForEach("td.teacherInfo", func(i int, e *colly.HTMLElement) {
			if i >= max {
				return
			}
			fmt.Printf("%d. 姓名: %s,", i+1, e.ChildText("span.content_title2 > a"))

			link := e.ChildAttr("p.infop > a", "href")
			if link == "" {
				fmt.Printf(" 網站: %s\n", "NULL")
				return
			}
			fmt.Printf(" 網站: %s\n", link)
		})
	})

	c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	c := colly.NewCollector()
	if web == "ptt" {
		ppt(c)
	} else if web == "ncku" {
		ncku(c)
	}

	c.Wait()
}
