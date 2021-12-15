package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"alex10101025/cw"
)

func main() {
	if len(os.Args) <= 3 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}
	cwSystem := cw.System{}

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	//fmt.Printf("%+v\n", PTTArticles.Articles[0])

	FBArticles := cwSystem.LoadFB("./data/fb.json")
	// fmt.Printf("%+v\n", FBArticles.Articles[0])
	IP_USER_NUM, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	KEYWORD_COUNT, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	KEYWORDS := os.Args[3:]
	//fmt.Print("aaa\n")
	fmt.Print(cwSystem.CountCyberWarriors(PTTArticles, IP_USER_NUM))
	fmt.Print(cwSystem.CountKeyWord(PTTArticles, FBArticles, KEYWORD_COUNT, KEYWORDS))
}
