package main

import (
	"fmt"
	"justin01010/cw"
)

func main() {
	cwSystem := cw.System{}
	fmt.Print(cwSystem)

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	cwSystem.CountCyberWarriors(PTTArticles)
	FBArticles := cwSystem.LoadFB("./data/fb.json")
	cwSystem.CountKeyWord(FBArticles, PTTArticles)
}
