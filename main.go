package main

import (
	"fmt"

	parser "github.com/Sotaneum/go-args-parser"
)

func main() {
	defaultOptions := map[string]string{
		"url": "",
		"remoteName": "origin",
		"targetFolder": "",
        "accessToken": "",
        "intervalMinute": "10",
		"gitUserName":"",
		"gitUserEmail":"",
	}
	fmt.Println("initialize")
	options := parser.ArgsJoinEnv(defaultOptions)
	New(options)
}