package main

import (
	"NCTU-GoProgramming-2021/lab9/andytsai2000/cw"
)

func main() {
	cwSystem := cw.New()

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	cwSystem.CountCyberWarriors(PTTArticles)

	FBArticles := cwSystem.LoadFB("./data/fb.json")
	cwSystem.CountKeyWord(PTTArticles, FBArticles)
}
