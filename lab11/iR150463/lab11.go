package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

var maxInfo int
var web string

func init() {
	// Define the other flags here
	flag.IntVar(&maxInfo, "max", 10, "Max Printing")
	flag.StringVar(&web, "w", "ptt", "Web page")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	if maxInfo > 45 {
		maxInfo = 45
	}

	c := colly.NewCollector()

	if web == "ptt" {
		var infoCount = 1
		c.OnHTML(".push", func(e *colly.HTMLElement) {
			if infoCount <= maxInfo {
				fmt.Print(strconv.Itoa(infoCount) + ". ")
				fmt.Print("名字: " + e.ChildText(".f3.hl.push-userid") + ", ")
				fmt.Print("留言" + e.ChildText(".f3.push-content") + ", ")
				fmt.Print("時間: " + e.ChildText(".push-ipdatetime"))
				fmt.Println()
				infoCount++
			} else {
				return
			}
		})

		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	} else {
		var infoCount = 1
		c.OnHTML(".teacherInfo", func(e *colly.HTMLElement) {
			if infoCount <= maxInfo {
				if len(e.ChildAttrs("a", "href")) > 1 && e.ChildAttrs("a", "href")[1] != "" {
					fmt.Print(strconv.Itoa(infoCount) + ". ")
					fmt.Print("姓名: " + e.ChildText(".content_title2") + ", ")
					fmt.Println("網站: " + e.ChildAttrs("a", "href")[1])
				} else {
					fmt.Print(strconv.Itoa(infoCount) + ". ")
					fmt.Print("姓名: " + e.ChildText(".content_title2") + ", ")
					fmt.Println("網站: " + "NULL")
				}
				infoCount++
			} else {
				return
			}
		})

		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}

	c.Wait()
}
