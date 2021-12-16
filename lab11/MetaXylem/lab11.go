package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

var max_size int
var website string

func init() {
	flag.IntVar(&max_size, "max", 10, "Max Printing")
	flag.StringVar(&website, "w", "ptt", "Web page")
}

func main() {
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		return
	}
	flag.Parse()
	c := colly.NewCollector()

	if website == "ptt" {
		c.OnHTML(".push", func(e *colly.HTMLElement) {
			if e.Index >= max_size {
				return
			}
			fmt.Print(e.Index+1, ". 名字: ", e.ChildText(".f3.hl.push-userid"), ", 留言", e.ChildText(".f3.push-content"), ", 時間: ", e.ChildText(".push-ipdatetime"), "\n")
		})
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else if website == "ncku" {
		c.OnHTML(".teacherInfo", func(e *colly.HTMLElement) {
			if e.Index >= max_size {
				return
			}
			href := e.ChildAttrs("a", "href")[1]
			if href == "" {
				href = "NULL"
			}
			fmt.Print(e.Index+1, ". 姓名: ", e.ChildText(".content_title2"), ", 網站: ", href, "\n")
		})
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}
	c.Wait()
}
