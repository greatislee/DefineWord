package main

import (
	"DefineWord/AssembleRequestUrl"
	"DefineWord/Parse"
	"fmt"
	"log"
)

func main() {

	// parse Args
	cl, err := Parse.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// AssembleRequestUrl
	urlS, err := AssembleRequestUrl.AssembleGoogleRequestUrl(cl)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(urlS.Q)
}
