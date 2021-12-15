package main

import (
	"fmt"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"alex10101025/cw"
)

func main() {

	cwSystem := cw.System{}

	cwSystem.Ptt = cwSystem.LoadPTT("./data/ptt.json")
	//fmt.Printf("%+v\n", PTTArticles.Articles[0])

	cwSystem.Fb = cwSystem.LoadFB("./data/fb.json")
	// fmt.Printf("%+v\n", FBArticles.Articles[0])
	fmt.Println(cwSystem)
}
