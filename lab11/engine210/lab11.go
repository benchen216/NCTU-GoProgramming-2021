package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func init() {
	// Define the other flags here
	// flag.IntVar()
}

func main() {
	// var w = flag.String("w", "https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html", "Website to crawl")
	var w = flag.String("w", "ptt", "Web page")
	var max = flag.Int("max", 10, "Max Printing")
	flag.Parse()

	argc := len(os.Args)
	if argc <= 1 {
		flag.PrintDefaults()
		return
	}

	// fmt.Println(*w)
	// fmt.Println(*max)

	c := colly.NewCollector()

	/*
		c.OnHTML("div#topbar.bbs-content > a#logo", func(e *colly.HTMLElement) {
			fmt.Println("demo1: " + e.Text)
		})

		c.OnHTML("div[class='bbs-content'] > a[id='logo']", func(e *colly.HTMLElement) {
			fmt.Println("demo2: " + e.Text)
		})

		c.OnHTML(".f2", func(e *colly.HTMLElement) {
			fmt.Println("demo3: " + e.Text)
		})

		c.OnHTML("#logo", func(e *colly.HTMLElement) {
			fmt.Println("demo4: " + e.Attr("href"))
		})
	*/

	if strings.Contains(*w, "ptt") {
		c.OnHTML("#main-content", func(e *colly.HTMLElement) {
			i := 1
			e.ForEach(".push", func(_ int, e *colly.HTMLElement) {
				if i > *max {
					return
				}
				fmt.Printf("%d. ", i)
				fmt.Printf("名字: %s, ", e.ChildText(".push-userid"))
				fmt.Printf("留言%s, ", e.ChildText(".push-content"))
				fmt.Printf("時間: %s\n", e.ChildText(".push-ipdatetime"))
				i += 1
			})
		})
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else if strings.Contains(*w, "ncku") {
		c.OnHTML("#tab1 > .teacherTableStyle > tbody", func(e *colly.HTMLElement) {
			i := 1
			e.ForEach(".teacherInfo", func(_ int, e *colly.HTMLElement) {
				if i > *max {
					return
				}
				fmt.Printf("%d. ", i)
				fmt.Printf("姓名: %s, ", e.ChildText("span > a"))
				if e.ChildAttr("p > a", "href") == "" {
					fmt.Println("網站: NULL")
				} else {
					fmt.Printf("網站: %s\n", e.ChildAttr("p > a", "href"))
				}
				i += 1
			})
		})
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}

	c.Wait()
}
