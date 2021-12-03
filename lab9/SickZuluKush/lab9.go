package main

import (
	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"lab9/cw"
	"os"
)

func main() {
	cwSystem := cw.System{}
	if !(&cwSystem).Init_value() {
		os.Exit(1)
	}

	PTTArticles := cw.System{}.LoadPTT("./data/ptt.json")
	FBArticles := cw.System{}.LoadFB("./data/fb.json")
	
	cwSystem.CountCyberWarriors(PTTArticles)
	cwSystem.CountKeyWord(PTTArticles, FBArticles)
}
