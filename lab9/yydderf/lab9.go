package main

import (
	"fmt"
    "os"
    "strconv"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"ydderf/cw"
)

func main() {
	cwSystem := cw.System{}
    if len(os.Args[1:]) < 3 {
        fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
        os.Exit(1)
    }

	PTTArticles := cwSystem.LoadPTT("../data/ptt.json")
	FBArticles := cwSystem.LoadFB("../data/fb.json")

    num, _ := strconv.Atoi(os.Args[1])
    cnt, _ := strconv.Atoi(os.Args[2])
    keywords := os.Args[3:]

    cwSystem.CountCyberWarriors(PTTArticles, num)
    cwSystem.CountKeyWord(PTTArticles, FBArticles, cnt, keywords)
}
