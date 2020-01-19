package main

import (
	"DefineWord/AssembleRequestUrl"
	"DefineWord/Parse"
	"DefineWord/httpClient"
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
