package AssembleRequestUrl

import (
	"DefineWord/Parse"
	"errors"
	"net/url"
	"unicode"
)

type UrlStruct struct {
	client string
	dt     string
	ie     string
	oe     string
	sl     string //origin language
	tl     string //target language
	q      string
}

// if han  return true
func checkHan(checkStr string) bool {
	if len(checkStr) == 0 {
		return false
	}
	for _, v := range checkStr {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func AssembleGoogleRequestUrl(line *Parse.CommandLine) (*string, error) {
	if line == nil {
		return nil, errors.New("invalid commandline")
	}

	t := UrlStruct{
		client: "gtx",
		dt:     "t",
		ie:     "UTF-8",
		oe:     "UTF-8",
		sl:     "",
		tl:     "",
		q:      "",
	}

	if checkHan(line.Word) {
		t.sl = "zh-CN"
		t.tl = "en"
	} else {
		t.sl = "en"
		t.tl = "zh-CN"
	}
	t.q = url.QueryEscape(line.Word)
	urlStr := "http://translate.google.cn/translate_a/single?" + "client=" + t.client + "&" + "dt=" + t.dt + "&" + "ie=" + t.ie + "&" + "oe=" + t.oe + "&" + "sl=" + t.sl + "&" + "tl=" + t.tl + "&" + "q=" + t.q
	return &urlStr, nil
}
