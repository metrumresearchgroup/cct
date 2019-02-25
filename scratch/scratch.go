package main

import (
	"encoding/json"
	"fmt"
	"github.com/metrumresearchgroup/cct/cc"
	"os"

	"github.com/spf13/cobra"
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
	// RootCmd represents the base command when called without any subcommands
	var RootCmd = &cobra.Command{
		Use:   "cct",
		Short: "conventional commits",
		Long:  fmt.Sprintf("cct cli version %s", 5),
	}

	fmt.Println(RootCmd.Use)

	RootCmd.Execute()

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
