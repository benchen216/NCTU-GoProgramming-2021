package main

import (
	"fmt"
	"os"
	"strconv"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"my_module/cw"
)

func main() {
	cwSystem := cw.System{}
	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	// fmt.Printf("%+v\n", PTTArticles.Articles[0])
	FBArticles := cwSystem.LoadFB("./data/fb.json")
	// fmt.Printf("%+v\n", FBArticles.Articles[0])
	// fmt.Println(cwSystem)

	/* Error handling */
	argsWithoutProg := os.Args[1:]
	// fmt.Println(argsWithoutProg)
	if len(argsWithoutProg) < 3 {
		fmt.Printf("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>\n")
		os.Exit(1)
	}
	// Get params
	IP_USER_NUM, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	KEYWORD_COUNT, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	KEYWORDS := argsWithoutProg[2:]

	cwSystem.CountCyberWarriors(PTTArticles, FBArticles, IP_USER_NUM)
	cwSystem.CountKeyWord(PTTArticles, FBArticles, KEYWORD_COUNT, KEYWORDS)
}

// go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>
