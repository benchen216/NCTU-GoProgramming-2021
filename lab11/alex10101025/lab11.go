package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

// flag.Usage = usage
// func usage() {
//     fmt.Fprintf(os.Stderr, "Usage of\n")
//     fmt.Fprintf(os.Stderr, " -max int\n")
//     flag.PrintDefaults()
// }
var MaxPrint int
var WebPage string

func init() {
	// Define the other flags here
	// flag.IntVar()
	flag.IntVar(&MaxPrint, "max", 10, "Max Printing")
	flag.StringVar(&WebPage, "w", "ptt", "Web page")
}

func main() {
	flag.Parse()
	if flag.NFlag() < 2 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	c := colly.NewCollector()

	// c.OnHTML("div#topbar.bbs-content > a#logo", func(e *colly.HTMLElement) {
	// 	fmt.Println("demo1: " + e.Text)
	// })

	// c.OnHTML("div[class='bbs-content'] > a[id='logo']", func(e *colly.HTMLElement) {
	// 	fmt.Println("demo2: " + e.Text)
	// })
	if WebPage == "ptt" {
		c.OnHTML(".push", func(e *colly.HTMLElement) {
			if e.Index >= MaxPrint {
				return
			}
			idx := strconv.Itoa(e.Index + 1)
			name := e.ChildText(".f3.hl.push-userid")
			msg := e.ChildText(".f3.push-content")
			time := e.ChildText(".push-ipdatetime")
			fmt.Println(idx + ". 名字: " + name + ", 留言: " + msg + ", 時間: " + time)
		})
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else if WebPage == "ncku" {
		c.OnHTML(".teacherInfo", func(e *colly.HTMLElement) {
			if e.Index >= MaxPrint {
				return
			}
			idx := strconv.Itoa(e.Index + 1)
			name := e.ChildText(".content_title2")
			hrefs := e.ChildAttrs("a", "href")
			web := ""
			if len(hrefs) > 1 {
				web = hrefs[1]
				if web == "" {
					web = "NULL"
				}
			}
			fmt.Println(idx + ". 名字: " + name + ", 網站: " + web)
		})
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}
	c.Wait()
}
