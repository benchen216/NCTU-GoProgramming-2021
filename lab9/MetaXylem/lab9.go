package main

import (
	"fmt"
	//"os"
	"asset/cw"
)

func main() {
	cwSystem := cw.System{}
	fmt.Print(cwSystem)


	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	cwSystem.CountCyberWarriors(PTTArticles)
	//fmt.Printf("%+v\n", PTTArticles.Articles[0])

	FBArticles := cwSystem.LoadFB("./data/fb.json")
	cwSystem.CountKeyWord(FBArticles, PTTArticles)
	//fmt.Printf("%+v\n", FBArticles.Articles[0])
}