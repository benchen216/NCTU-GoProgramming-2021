package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

var (
	max int
	webpage string
)

func init() {
	// Define the other flags here
	flag.IntVar(&max, "max", 10, "Max Printing")
	flag.StringVar(&webpage, "w", "ptt", "Web page")
}

func check_flag() bool {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return false
	}
	
	return true
}

func handle_ppt(c *colly.Collector) {
	c.OnHTML("div#main-content", func(e *colly.HTMLElement) {
		e.ForEach("div.push", func(i int, e *colly.HTMLElement) {
			if i >= max {
				return
			}
			
			fmt.Printf("%v. 名字: %s, ", i + 1, e.ChildText(".push-userid"))
			fmt.Printf("留言%s, ", e.ChildText(".push-content"))
			fmt.Printf("時間: %s\n", e.ChildText(".push-ipdatetime"))
		})
	})
}

func handle_ncku(c *colly.Collector) {
	c.OnHTML("div#tab1", func(e *colly.HTMLElement) {
		e.ForEach("td.teacherInfo", func(i int, e *colly.HTMLElement) {
			if i >= max {
				return
			}
				
			fmt.Printf("%v. 姓名: %s, ", i + 1, e.ChildText("span.content_title2 > a"))
			
			link := e.ChildAttr("p.infop > a", "href")
			if link == "" {
				link = "NULL"
			}
			fmt.Printf("網站: %s\n", link)
		})
	})
}

func main() {
	if !check_flag() {
		return
	}
	
	c := colly.NewCollector()

	if webpage == "ncku" {
		handle_ncku(c)
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	} else {
		handle_ppt(c)		
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	}

	c.Wait()
}
