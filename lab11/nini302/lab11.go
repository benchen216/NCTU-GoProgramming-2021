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
	flag.IntVar(&max, "max", 10, "Max Printing") // Define the other flags here
	flag.StringVar(&web, "w", "ptt", "Web page") // flag.IntVar()
}
func ppt(c *colly.Collector) {
	c.OnHTML("div#main-content", func(e *colly.HTMLElement) {
		e.ForEach("div.push", func(i int, e *colly.HTMLElement) {
			if i >= max {
				return
			}
			fmt.Printf("%d. 名字: %s, ", i+1, e.ChildText(".push-userid"))
			fmt.Printf("留言%s ", e.ChildText(".push-content"))
			fmt.Printf("時間: %s\n", e.ChildText(".push-ipdatetime"))
		})
	})
}
func ncku(c *colly.Collector) {
	c.OnHTML("div#tab1", func(h *colly.HTMLElement) {
		h.ForEach("td.teacherInfo", func(i int, h *colly.HTMLElement) {
			if i >= max {
				return
			}
			fmt.Printf("%d. 姓名: %s, ", i+1, h.ChildText("span.content_title2 > a"))

			link := h.ChildAttr("p.infop > a", "herf")
			if link == "" {
				link = "NULL"
			}
			fmt.Printf("網站: %s\n", link)
		})
	})
}
func main() {
	flag.Parse()
	// flag.PrintDefaults()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}
	c := colly.NewCollector()

	if web == "ncku" {
		ncku(c)
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	} else {
		ppt(c)
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	}
	c.Wait()
}
