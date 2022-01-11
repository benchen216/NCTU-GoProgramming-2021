package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

//flags
var (
	max int
	web string
)

func init() {
	flag.IntVar(&max, "max", 10, "Max Printing")
	flag.StringVar(&web, "w", "ptt", "Web page")
}

func ptt(c *colly.Collector) {
	c.OnHTML("div[id='main-content']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='push']", func(i int, e *colly.HTMLElement) {
			if i < max {
				str := strconv.Itoa(i+1) + ". 名字: " + e.ChildText(".push-userid") +
					", 留言" + e.ChildText(".push-content") +
					", 時間: " + e.ChildText(".push-ipdatetime") + "\n"
				strings.ReplaceAll(str, "(MISSING)", "")
				fmt.Printf(str)
			}
		})
	})

}

func ncku(c *colly.Collector) {
	c.OnHTML("div[class='content_maintext tab-content'] > div[id='tab1'] > table[class='teacherTableStyle']", func(e *colly.HTMLElement) {
		e.ForEach("td[class='teacherInfo']", func(i int, e *colly.HTMLElement) {
			if i < max {
				if len(e.ChildAttr("p[class='infop'] > a", "href")) == 0 {
					fmt.Printf(strconv.Itoa(i+1) + ". 姓名: " + e.ChildText("span[class='content_title2'] > a") + ", 網站: NULL\n")
				} else {
					fmt.Printf(strconv.Itoa(i+1) + ". 姓名: " + e.ChildText("span[class='content_title2'] > a") + ", 網站: " + e.ChildAttr("p[class='infop'] > a", "href") + "\n")
				}
			}

		})
	})

}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	c := colly.NewCollector()

	if web == "ptt" {
		ptt(c)
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else {
		ncku(c)
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}

	c.Wait()
}
