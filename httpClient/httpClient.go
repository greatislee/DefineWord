package httpClient

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetResponse(str *string) (*string, error) {
	if str == nil {
		return nil, errors.New("GetResponse: Invalid Parameter")
	}

	var errText strings.Builder

	timeout := 3 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	var Body io.Reader
	request, err := http.NewRequest("GET", *str, Body)
	if err != nil {
		errText.WriteString("GetResponse: " + "Get " + *str + " " + err.Error())
		return nil, errors.New(errText.String())
	}
	res, err := client.Do(request)
	if err != nil {
		errText.WriteString("GetResponse: " + "Do " + *str + " " + err.Error())
		return nil, errors.New(errText.String())
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errText.WriteString("GetResponse: " + "ReadAll " + *str + " " + err.Error())
		return nil, errors.New(errText.String())
	}

	resultStr := string(result)
	return &resultStr, nil
}
