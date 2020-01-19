package main

import (
	"DefineWord/AssembleRequestUrl"
	"DefineWord/Parse"
	"DefineWord/httpClient"
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
	url, err := AssembleRequestUrl.AssembleGoogleRequestUrl(cl)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//httpClient
	response, err := httpClient.GetResponse(url)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(*response)
}
