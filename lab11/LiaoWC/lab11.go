package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

var MaxNum int
var Web string

func init() {
	// Define the other flags here
	// flag.IntVar()
	flag.IntVar(&MaxNum, "max", 10, "Max Printing")
	flag.StringVar(&Web, "w", "ptt", "Web page")
}

func main() {
	// Init
	flag.Parse()
	c := colly.NewCollector()

	// Check argv
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	// Define crawler
	if Web == "ptt" {
		c.OnHTML("#main-container", func(e *colly.HTMLElement) {
			e.ForEach(".push", func(idx int, e *colly.HTMLElement) {
				if idx < MaxNum {
					userid := e.ChildText(".push-userid")
					content := e.ChildText(".push-content")
					ipdatetime := e.ChildText(".push-ipdatetime")
					fmt.Printf("%s. 名字: %s, 留言%s, 時間: %s\n", strconv.Itoa(idx+1), userid, content, ipdatetime)
				}
			})
		})
	} else if Web == "ncku" {
		c.OnHTML("#tab1 tbody", func(e *colly.HTMLElement) {
			e.ForEach(".teacherInfo", func(idx int, e *colly.HTMLElement) {
				if idx < MaxNum {
					name := e.ChildText(".content_title2")
					href := e.ChildAttr(".infop a", "href")
					if href == "" {
						href = "NULL"
					}
					fmt.Printf("%s. 姓名: %s, 網站: %s\n", strconv.Itoa(idx+1), name, href)
				}
			})
		})
	}

	// Visit
	if Web == "ptt" {
		err := c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
		if err != nil {
			log.Fatalln("c.Visit failed")
		}
	} else {
		err := c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
		if err != nil {
			log.Fatalln("c.Visit failed")
		}
	}

	c.Wait()
}
