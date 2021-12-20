package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var sint int
var w string

func init() {
	// Define the other flags here

	flag.IntVar(&sint, "max", 0, "Max Printing (default 10)")
	flag.StringVar(&w, "w", "", "Web page (default \"ptt\")")

}

func main() {

	flag.Parse()
	if sint == 0 && len(w) == 0 {
		flag.PrintDefaults()
		return
	} else if sint == 0 {
		sint = 10
	} else if w == "" {
		w = "ptt"
	}

	c := colly.NewCollector()

	if w == "ncku" {
		var name, addr []string
		c.OnHTML("div[id=tab1]", func(el *colly.HTMLElement) {
			el.ForEach(".teacherInfo", func(_ int, e *colly.HTMLElement) {
				name = append(name, e.ChildText(".content_title2"))
				addr = append(addr, e.ChildAttr(".infop a[href]", "href"))

			})
		})
		// c.OnHTML(".infop a[href]", func(e *colly.HTMLElement) {
		// 	addr = append(addr, e.Attr("href"))
		// 	// fmt.Println(e.Attr("href"))
		// })

		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
		for i := 1; i <= sint; i++ {
			if i > len(name) {
				return
			}
			sa := strconv.Itoa(i)
			if len(addr[i-1]) == 0 {
				addr[i-1] = "NULL"
			}
			addr[i-1] = strings.Replace(addr[i-1], " ", "", -1)

			fmt.Printf("%s. 姓名: %s, 網站: %s\n", sa, name[i-1], addr[i-1])
		}
	} else {
		var name, cmt, time []string
		// var name, cmt []string

		c.OnHTML(".push-userid", func(e *colly.HTMLElement) {
			name = append(name, e.Text)
			// fmt.Println(e.Text)
		})
		c.OnHTML(".push-content", func(e *colly.HTMLElement) {
			cmt = append(cmt, e.Text)
		})
		c.OnHTML(".push-ipdatetime", func(e *colly.HTMLElement) {
			time = append(time, e.Text)
		})

		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
		// fmt.Println(len(name))
		for i := 1; i <= sint; i++ {
			sa := strconv.Itoa(i)
			name[i-1] = strings.Replace(name[i-1], " ", "", -1)
			fmt.Printf("%s. 名字: %s, 留言%s, 時間:%s", sa, name[i-1], cmt[i-1], time[i-1])
		}
	}
	//
	c.Wait()
}
