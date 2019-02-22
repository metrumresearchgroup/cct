package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/metrumresearchgroup/cct/cc"
)

func main() {
	commitMsg := `feat: implemented cool stuff

some body text

some more body text

References #12
Closes #13`
	commitMsg = `bad message
no good
`
	msg, err := cc.NewCommitMessage(commitMsg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	PrettyPrint(msg)
}

// PrettyPrint allows pretty printing of datastructures to the consol
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
