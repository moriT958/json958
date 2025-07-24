package main

import (
	"encoding/json"
	"errors"
)

func getHeadChar(line string) (rune, error) {
	if line == "" {
		return 0, errors.New("Unexpected end of input")
	}
	return []rune(line)[0], nil
}

// map data into v
func parse(data []byte, v any) error {
	return nil
}

func main() {
	var NilData any
	json.Unmarshal([]byte(`{"name": "example"}`), NilData)
}
