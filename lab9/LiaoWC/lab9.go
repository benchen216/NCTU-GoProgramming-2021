package main

import (
	"kkk/cw"

	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	/* Check argc */
	if len(os.Args) <= 3 {
		fmt.Print("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>\n")
		os.Exit(1)
	}

	/* Fetch argv */
	IPUserNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	KeywordCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	Keywords := os.Args[3:]

	/* Init cwSystem */
	cwSystem := cw.System{}

	/* Load json */
	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	FBArticles := cwSystem.LoadFB("./data/fb.json")

	/* IPUserNum */
	fmt.Print(cwSystem.CountCyberWarriors(IPUserNum,PTTArticles))

	/* KeywordCount */
	fmt.Print(cwSystem.CountKeyWord(KeywordCount, Keywords,PTTArticles,FBArticles))
}
