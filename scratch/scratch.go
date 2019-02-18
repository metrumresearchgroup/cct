package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("can test stuff here")
}

// PrettyPrint allows pretty printing of datastructures to the consol
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
