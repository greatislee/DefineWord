package httpClient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func GetResponse(str *string) (*string, error) {
	if str == nil {
		return nil, errors.New("GetResponse: Invalid Parameter")
	}

	timeout := 3 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(*str)
	if err != nil {
		return nil, errors.New("GetResponse: " + "[Get] " + err.Error())
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("GetResponse: " + "ReadAll " + *str + " " + err.Error())
	}

	resultStr := string(result)
	return &resultStr, nil
}
