package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

func init() {
	// Define the other flags here
	// flag.IntVar()
}

func main() {
	flag.Parse()
	// flag.PrintDefaults()

	c := colly.NewCollector()

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

	c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	// c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	c.Wait()
}
