package main

import (
	"fmt"
	"yydderf/cw"
)

func main() {
	cwSystem := cw.System{}
	fmt.Print(cwSystem)

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	cwSystem.CountCyberWarriors(PTTArticles)
	FBArticles := cwSystem.LoadFB("./data/fb.json")
	cwSystem.CountKeyWord(FBArticles, PTTArticles)
}
