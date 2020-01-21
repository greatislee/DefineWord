package Parse

import (
	"errors"
	"os"
)

type CommandLine struct {
	Word string
}

func ParseArgs() (*CommandLine, error) {
	if len(os.Args) <= 1 {
		return nil, errors.New("Please Enter: dw --help")
	}
	if os.Args[1] == "--help" {
		return nil, errors.New("Usage: dw word")
	}

	var strTemp string
	for i := range os.Args {
		if i == 0 {
			continue
		}
		strTemp += os.Args[i] + " "
	}
	commandL := &CommandLine{
		Word: strTemp,
	}
	return commandL, nil
}
