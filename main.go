package main

import (
	"fmt"
	"HFish/utils/setting"
	"os"
)

func main() {
	//setting.Run()
	args := os.Args
	if args == nil || len(args) < 2 {
		setting.Help()
	} else {
		if args[1] == "help" || args[1] == "--help" {
			setting.Help()
		} else if args[1] == "init" || args[1] == "--init" {
			setting.Init()
		} else if args[1] == "version" || args[1] == "--version" {
			fmt.Println("v0.6.3")
		} else if args[1] == "run" || args[1] == "--run" {
			setting.Run()
		} else {
			setting.Help()
		}
	}
}
