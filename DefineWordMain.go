package main

import (
	"DefineWord/AssembleRequestUrl"
	"DefineWord/Parse"
	"fmt"
	"log"
	"os"
)

func main() {

	// parse Args
	if os.Args[1] == "--help" {
		fmt.Println("Usage: dw word")
		return
	}
	cl, err := Parse.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// AssembleRequestUrl
	_, err = AssembleRequestUrl.AssembleGoogleRequestUrl(cl)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
