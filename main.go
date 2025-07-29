package main

import (
	"encoding/json"
)

func main() {
	var NilData any
	json.Unmarshal([]byte(`{"name": "example"}`), NilData)
}
