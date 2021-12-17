package main

import (
	"flag"
	"fmt"
    "os"

	"github.com/gocolly/colly"
)

var webStr string
var maxNum int

func init() {
	// Define the other flags here
    flag.StringVar(&webStr, "w", "ptt", "Web page")
    flag.IntVar(&maxNum, "max", 10, "Max Printing")
}

func PTT(c *colly.Collector)([]string, []string, []string) {
    users := make([]string, 0, 0)
    comms := make([]string, 0, 0)
    times := make([]string, 0, 0)
    c.OnHTML("#main-content", func(e *colly.HTMLElement) {
        e.ForEach(".push", func(index int, ps *colly.HTMLElement) {
            users = append(users, ps.ChildText(".hl.push-userid"))
            comms = append(comms, ps.ChildText(".push-content"))
            times = append(times, ps.ChildText(".push-ipdatetime"))
        })
    })

    c.Visit("https://www.ptt.cc/bbs/Stock/M.1610102078.A.16E.html")
    c.Wait()
    return users, comms, times
}

func NCKU(c *colly.Collector)([]string, []string) {
    names := make([]string, 0, 0)
    sites := make([]string, 0, 0)
    c.OnHTML("#tab1 > table > tbody > tr > td.teacherInfo", func(e *colly.HTMLElement) {
        names = append(names, e.ChildText("span"))
        tmp := e.ChildAttr("p.infop > a", "href")
        if tmp == "" {
            sites = append(sites, "NULL")
        } else {
            sites = append(sites, tmp)
        }
    })

    c.Visit("https://www.csie.ncku.edu.tw/ncku_csie/depmember/teacher")
    c.Wait()
    return names, sites
}

func main() {
    if len(os.Args) == 1 {
        flag.PrintDefaults()
        return
    }
	flag.Parse()

    c := colly.NewCollector()
    if webStr == "ptt" {
        var users []string
        var comms []string
        var times []string
        users, comms, times = PTT(c)

        if len(users) < maxNum {
            maxNum = len(users)
        }
        for i := 0; i < maxNum; i++ {
            fmt.Printf("%d. 名字: %s, 留言%s, 時間: %s\n", i+1, users[i], comms[i], times[i])
        }
    } else if webStr == "ncku" {
        var names []string
        var sites []string
        names, sites = NCKU(c)
        if len(names) < maxNum {
            maxNum = len(names)
        }
        for i := 0; i < maxNum; i++ {
            fmt.Printf("%d. 姓名: %s, 網站: %s\n", i+1, names[i], sites[i])
        }
    }

}
