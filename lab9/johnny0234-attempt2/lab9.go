package main

import (
	"fmt"

	// [TODO] set your module name, go mod init my_module->M
	"johnny0234/cw"

	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}

	cwSystem := cw.System{}

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	cwSystem.CountCyberWarriors(PTTArticles)
	//fmt.Printf("%+v\n", PTTArticles.Articles[0])

	FBArticles := cwSystem.LoadFB("./data/fb.json")
	cwSystem.CountKeyWord(PTTArticles, FBArticles)
	//fmt.Printf("%+v\n", FBArticles.Articles[0])

	//fmt.Println(cwSystem)
}
