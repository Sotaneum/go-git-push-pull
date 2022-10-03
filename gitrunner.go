package main

import (
	"fmt"
	"strconv"
	"time"

	gitfiles "github.com/Sotaneum/go-git-files"
)

func update(db *gitfiles.GitFiles, intervalMinute int64) {
	target := int(intervalMinute * 60)
	term := 1
	fmt.Println(target)
	for  {
		if term >= target {
			db.Push()
			fmt.Println("push")
			term = 0
		}
		term = term + 1
		time.Sleep(time.Second * 1)
	}
}

func New(options map[string]string){
	user := &gitfiles.User{
		Name:  options["gitUserName"],
		Email: options["gitUserEmail"],
	}
	fmt.Println(user)
	core := gitfiles.GitFiles{}
	core.Init(*user, options["targetFolder"], options["url"], options["accessToken"], options["remoteName"])
	minute, err := strconv.ParseInt(options["intervalMinute"], 10, 10)
	if err!=nil{
		panic("minute error")
	}
	update(&core, minute)
}