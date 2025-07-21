package main

import "errors"

func getHeadChar(line string) (rune, error) {
	if line == "" {
		return 0, errors.New("Unexpected end of input")
	}
	return []rune(line)[0], nil
}
