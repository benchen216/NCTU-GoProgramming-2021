package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

var max_output int
var which_article string

func init() {
	// Define the other flags here
	flag.IntVar(&max_output, "max", 10, "Max Printing")
	flag.StringVar(&which_article, "w", "ptt", "Web page")
}

func main() {
	flag.Parse()
	if flag.NFlag()==0{
		flag.PrintDefaults()
	}else{
		c := colly.NewCollector()
		index :=1
		if which_article=="ptt"{
			c.OnHTML(".push", func(e *colly.HTMLElement) {
				if index>max_output{
					return
				}
				id := e.ChildText(".push-userid")
				content := e.ChildText(".push-content")
				time := e.ChildText(".push-ipdatetime")
				fmt.Printf("%d. 名字: %s, 留言%s, 時間: %s\n",index,id,content,time)
				index++
			})
			c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
		}else{
			c.OnHTML("div", func(e *colly.HTMLElement) {
				if e.Attr("id")=="tab1"{
					e.ForEach("td.teacherInfo", func(_ int,f *colly.HTMLElement) {
						if index>max_output{
							return
						}
						name := f.ChildText(".content_title2")
						temp := f.ChildAttrs("a","href")
						url := "NULL"
						if len(temp)>=2 && temp[1]!=""{
							url = temp[1]
						}
						fmt.Printf("%d. 姓名: %s, 網站: %s\n",index,name,url)
						index++
					})
				}
			})
			c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
		}
		c.Wait()
	}
}
