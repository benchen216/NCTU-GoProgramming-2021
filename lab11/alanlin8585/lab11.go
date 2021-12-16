package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
	"os"
)

var search_type string
var search_max int

func init() {
	// Define the other flags here
	// flag.IntVar() 
    flag.IntVar(&search_max, "max", 10, "Max Printing")
	flag.StringVar(&search_type, "w", "ptt", "Web page")
}

func main() {
    if len(os.Args) <= 1 {
        flag.PrintDefaults()
        return
    }
        
	flag.Parse()
	
	// flag.PrintDefaults()

	c := colly.NewCollector()

	/*c.OnHTML("div#topbar.bbs-content > a#logo", func(e *colly.HTMLElement) {
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
	})*/
	
	if search_type == "ptt" {
	    c.OnHTML("div[class='push']", func(e *colly.HTMLElement) {
	        if e.Index + 1 <= search_max {
	            fmt.Print(e.Index + 1, ". 名字: ", e.ChildText("span[class='f3 hl push-userid']"), ", 留言", e.ChildText("span[class='f3 push-content']"), ", 時間: ", e.ChildText("span[class='push-ipdatetime']"), "\n")
	        }
	    })
	    c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else if search_type == "ncku" {
	    c.OnHTML("div[class='content_maintext tab-content'] > div[class='content_tab tab-pane active'] > table[class='teacherTableStyle'] > tbody > tr > td[class='teacherInfo']", func(e *colly.HTMLElement) {
	        if e.Index + 1 <= search_max {
	            url := e.ChildAttr("p[class='infop'] > a", "href")
	            if len(url) == 0 {
	                url = "NULL"
                }
	            fmt.Print(e.Index + 1, ". 姓名: ", e.ChildText("span[class='content_title2']"), ", 網站: ", url, "\n")
	        }
	    })
	    c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}

	c.Wait()
}
