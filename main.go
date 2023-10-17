package main

import (
	"flag"
	"fmt"
	"github.com/refaldyrk/lidya/constant"
	"github.com/refaldyrk/lidya/helper"
	"os"
)

func main() {
	filename := os.Args[1]
	helpFlag := flag.Bool("h", false, "-h")
	langFlag := flag.String("lang", "", "lidya --lang=go")
	flag.Parse()

	if *helpFlag {
		fmt.Println(constant.HELP_STRING)
		return
	}

	if filename == "help" {
		fmt.Println(constant.HELP_STRING)
		return
	}

	if filename != "" {
		helper.GenerateCode(filename, *langFlag)
	} else {
		fmt.Println("Filename Missing!")
		os.Exit(1)
	}

	fmt.Println("Thank U, Using Lidya :D")
}
