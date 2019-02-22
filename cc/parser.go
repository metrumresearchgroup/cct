package cc

import (
	"errors"
	"fmt"
	"strings"
)

// NewCommitMessage creates a new commit message from a string
func NewCommitMessage(s string) (CommitMessage, error) {
	var cm CommitMessage
	sa := strings.Split(s, "\n")
	if len(sa) > 1 {
		// TODO: fix
		// this is wrong/incomplete but was used for demonstration
		cm.Body = sa[1:]
		cm.Footer = []string{sa[len(sa)-1]}
	}
	// get the first line
	head := strings.SplitN(sa[0], ":", 2)
	if len(head) != 2 {
		cm.Type = "MISSING"
		cm.Description = sa[0]
		return cm, errors.New(fmt.Sprintf("invalid header %s", sa[0]))
	}
	cm.Type = head[0]
	cm.Description = head[1]
	// parse out the two elements of the first line

	// do stuff with rest for body and footer
	return cm, nil
}
