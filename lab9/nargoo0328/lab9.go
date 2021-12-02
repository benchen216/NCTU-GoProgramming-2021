package main

import (
	"fmt"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"nargoo0328/cw"
)

func main() {
	cwSystem := cw.System{}
	fmt.Print(cwSystem)
	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	cwSystem.CountCyberWarriors(PTTArticles)
	FBArticles := cwSystem.LoadFB("./data/fb.json")
	cwSystem.CountKeyWord(PTTArticles,FBArticles)
}
