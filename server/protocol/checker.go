package protocol

import (
	"strings"
)

func Check(msg string) (bool, string) {
	args := strings.Split(msg, " ")
	switch args[0] {
	case "GET":
		return getCheck(args), "GET"
	case "SET":
		return setCheck(args), "SET"
	case "DEL":
		return delCheck(args), "DEL"
	default:
		return false, "UNKNOWN"
	}

}

func getCheck(args []string) bool {
	if len(args) != 3 {
		return false
	}
	return true
}
func setCheck(args []string) bool {
	if len(args) != 4 {
		return false
	}
	return true
}
func delCheck(args []string) bool {
	if len(args) != 3 {
		return false
	}
	return true
}
