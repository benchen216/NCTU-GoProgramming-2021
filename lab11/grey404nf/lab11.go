package main

import (
	"flag"
	"fmt"
	"strings"
	"os"

	"github.com/gocolly/colly"
)

func init() {
	// Define the other flags here
	// flag.IntVar()
}

func main() {
	cnt1:=0
	cnt2:=0
	cnt3:=0
	max:=flag.Int("max", 10, "Max Printing")
	w:=flag.String("w", "ptt", "Web page")
	
	flag.Parse()
	if flag.NFlag()==0{
		flag.PrintDefaults()
		os.Exit(0)
	}
	
	c := colly.NewCollector()
	
	var s [1000][3]string
	
	if *w=="ptt" {
		c.OnHTML("span[class='f3 hl push-userid']", func(e *colly.HTMLElement) {
			s[cnt1][0]=e.Text
			cnt1++
			
		})
		
		c.OnHTML("span[class='f3 push-content']", func(e *colly.HTMLElement) {
			s[cnt2][1]=e.Text
			cnt2++
		})
		
		c.OnHTML("span[class='push-ipdatetime']", func(e *colly.HTMLElement) {
			s[cnt3][2]=e.Text
			cnt3++
		})
		
		c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
	}else if *w=="ncku" {
		c.OnHTML("div#tab1 > table > tbody > tr > td > span > a", func(e *colly.HTMLElement) {
			s[cnt1][0]=e.Text
			cnt1++
		})
		
		c.OnHTML("div#tab1 > table > tbody > tr > td > p > a", func(e *colly.HTMLElement) {
			s[cnt2][1]=e.Attr("href")
			cnt2++
		})
		
		c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
	}
	
	c.Wait()
	
	if *max>cnt1 {
		*max=cnt1
	}
	
	if *w=="ptt" {
		for i:=0;i<*max;i++ {
			s[i][0]=strings.ReplaceAll(s[i][0], " ", "")
			//s[i][1]=strings.ReplaceAll(s[i][1], " ", "")
			s[i][2]=strings.ReplaceAll(s[i][2], " ", "")
			fmt.Printf("%d. 名字: %s, 留言%s, 時間: %s", i+1, s[i][0], s[i][1], s[i][2])
		}
	} else if *w=="ncku" {
		for i:=0;i<*max;i++ {
			s[i][0]=strings.ReplaceAll(s[i][0], " ", "")
			s[i][1]=strings.ReplaceAll(s[i][1], " ", "")
			if s[i][1]=="" {
				s[i][1]="NULL"
			}
			fmt.Printf("%d. 姓名: %s, 網站: %s\n", i+1, s[i][0], s[i][1])
		}
	}
}
