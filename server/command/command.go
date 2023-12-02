package command

import (
	"fmt"
	"strings"
)

func Do(command string, msg string) (string, error) {
	args := strings.Split(msg, " ")
	switch command {
	case "GET":
		return get(args[1])
	case "SET":
		return set(args[1], args[2])
	case "DEL":
		return del(args[1])
	}
	return "Invalid", nil
}

func get(key string) (string, error) {
	fmt.Printf("%s will be get", key)
	return "", nil
}

func set(key string, value string) (string, error) {
	fmt.Printf("%s and %s will be set", key, value)
	return "", nil
}

func del(key string) (string, error) {
	fmt.Printf("%s will be del", key)
	return "", nil
}
