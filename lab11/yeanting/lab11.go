package main

import (
	"flag"
	"fmt"
	"strings"
	"github.com/gocolly/colly"
)

var w string
var max int
func init() {
	// Define the other flags here
	flag.StringVar(&w, "w", "ptt", "Web page")
	flag.IntVar(&max, "max", 10, "Max Printing")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 || flag.NArg()>0 {
		flag.PrintDefaults()
	}

	c := colly.NewCollector()

	if w == "ptt" {
		var userid []string
		var content []string
		var time []string
		c.OnHTML("div#main-container > div#main-content.bbs-screen.bbs-content > div.push > span.f3.hl.push-userid", func(e *colly.HTMLElement) {
			name := strings.Replace(e.Text, " ", "", -1)
			userid = append(userid, name)
		})
		c.OnHTML("div#main-container > div#main-content.bbs-screen.bbs-content > div.push > span.f3.push-content", func(e *colly.HTMLElement) {
			content = append(content, e.Text)
		})
		c.OnHTML("div#main-container > div#main-content.bbs-screen.bbs-content > div.push > span.push-ipdatetime", func(e *colly.HTMLElement) {
			time = append(time, e.Text)
		})

		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
		c.Wait()

		for i:=0; i<max; i=i+1 {
			fmt.Printf("%d. 名字: %s, 留言%s, 時間:%s\n", (i+1), userid[i], content[i], time[i])
		}

	} else {
		var t_name []string
		var t_web []string

		c.OnHTML("span[class='content_title2'] > a", func(e *colly.HTMLElement) {
			t_name = append(t_name, e.Text)
		})
	
		c.OnHTML(".infop > a", func(e *colly.HTMLElement) {
			tmp := e.Attr("href")
			if tmp[0] == ' ' {
				if strings.Replace(tmp, " ", "", -1) == "" {
					t_web = append(t_web, " NULL")
				} else {
					t_web = append(t_web, tmp)
				}
			}
		})
		
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
		c.Wait()

		for i:=0; i<max; i=i+1 {
			fmt.Printf("%d. 姓名: %s, 網站:%s\n", (i+1), t_name[i], t_web[i])
		}
	}
	

	
	
	
}
