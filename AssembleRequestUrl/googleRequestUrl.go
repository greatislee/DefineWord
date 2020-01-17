package AssembleRequestUrl

import (
	"DefineWord/Parse"
	"errors"
)

type UrlStruct struct {
	client string
	sl     string
	tl     string
	hl     string
	dt     [8]string
	souce  string
	ssel   string
	tsel   string
	kc     string
	tk     float32
	Q      string
}

func AssembleGoogleRequestUrl(line *Parse.CommandLine) (*UrlStruct, error) {
	if line == nil {
		return nil, errors.New("invalid commandline")
	}

	t := &UrlStruct{
		Q: line.Word,
	}
	return t, nil
}
