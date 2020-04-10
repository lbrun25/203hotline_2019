package hotline

import (
	"fmt"
	"os"
	"regexp"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBeInteger = "Argument(s) must be positive integers.\n"
)

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if (arg == "-h") {
			return true
		}
	}
	return false
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > 2 {
		fmt.Println(tooManyArgs)
		return false
	}
	for _, arg := range argsWithoutProg {
		var re = regexp.MustCompile("[0-9]")

		match := re.FindAllString(arg, -1)
		if len(arg) != len(match) {
			fmt.Println(mustBeInteger)
			return false
		}
	}
	return true
}