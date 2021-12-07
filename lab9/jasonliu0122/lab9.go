package main

import (
//	"fmt"

	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"lab9/cw"
)

func main() {
//	fmt.Print("100")
	cwSystem := cw.System{}
//	fmt.Print("101")
	cwSystem.String()
//	fmt.Print("102")
	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
//	fmt.Printf("%+v\n", PTTArticles.Articles[0])
	cwSystem.CountCyberWarriors(PTTArticles);
	FBArticles := cwSystem.LoadFB("./data/fb.json")
//	fmt.Printf("%+v\n", FBArticles.Articles[0])
	cwSystem.CountKeyWord(PTTArticles,FBArticles);	
//	fmt.Println(cwSystem)
}
