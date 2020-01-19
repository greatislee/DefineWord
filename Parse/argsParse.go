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
	commandL := &CommandLine{
		Word: os.Args[1],
	}
	return commandL, nil
}
