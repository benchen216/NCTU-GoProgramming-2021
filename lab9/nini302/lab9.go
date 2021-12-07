package main

import (
	"fmt"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"nini/cw"
)

func main() {
	cwSystem := cw.System{}

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	fmt.Printf("%+v\n", PTTArticles.Articles[0])

	FBArticles := cwSystem.LoadFB("./data/fb.json")
	fmt.Printf("%+v\n", FBArticles.Articles[0])

	fmt.Println(cwSystem)
}
