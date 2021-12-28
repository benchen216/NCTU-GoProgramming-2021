package main

import (
	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"fmt"
	"lab9/cw"
	"os"
	"strconv"
)

func main() {
	cwSystem := cw.System{}

	cwSystem.LoadPTT("../data/ptt.json")
	// PTTArticles := cwSystem.LoadPTT("../data/ptt.json")
	// fmt.Printf("%+v\n", PTTArticles.Articles[0])

	cwSystem.LoadFB("../data/fb.json")
	// FBArticles := cwSystem.LoadFB("../data/fb.json")
	// fmt.Printf("%+v\n", FBArticles.Articles[0])
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}

	ip_user_num, _ := strconv.Atoi(os.Args[1])
	cwSystem.CountCyberWarriors(ip_user_num)

	keywords := os.Args[3:]
	// st := []string{"蔡英文", "韓國瑜", "宋楚瑜"}
	keyword_count, _ := strconv.Atoi(os.Args[2])
	cwSystem.CountKeyWord(keyword_count, keywords)

	// fmt.Println(cwSystem)
}
