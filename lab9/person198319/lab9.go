package main

import (
	// [TODO] set your module name, go mod init my_module->MOUDULE_NAME
	"fmt"
	"lab9/cw"
	"os"
	"strconv"
)

func main() {
	cwSystem := cw.System{}

	cwSystem.LoadPTT("data/ptt.json")

	cwSystem.LoadFB("data/fb.json")
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run lab9.go <IP_USER_NUM> <KEYWORD_COUNT> <KEYWORD...>")
		os.Exit(1)
	}

	user_ip_num, _ := strconv.Atoi(os.Args[1])
	cwSystem.CountCyberWarriors(user_ip_num)

	keywords := os.Args[3:]
	keyword_count, _ := strconv.Atoi(os.Args[2])
	cwSystem.CountKeyWord(keyword_count, keywords)
}
