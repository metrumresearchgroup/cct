package cc

import (
	"errors"
	"fmt"
	"strings"
)


//Todo: Set valid tags in init code.
//List of approved tags
var ValidTags []string = []string{
	"wip",
	"add",
	"doc",
	"fix",
}

// NewCommitMessage creates a new commit message from a string
func NewCommitMessage(s string) (CommitMessage, error) {
	message, err := ParseCommitMessage(s)
	if err != nil {
		if !HasValidTag(message) {
			err = errors.New(fmt.Sprintf("unrecognized commit type: %s", message.Type))
		}
	}
	return message, err
}

func ParseCommitMessage(s string) (CommitMessage, error) {
	var cm CommitMessage
	var header string
	var commitType string
	var description string
	var body []string = []string{}
	var footer []string = []string{}

	sections := strings.Split(s, "\n\n")

	if len(sections) > 0 {
		header = sections[0]
	} else {
		return cm, errors.New("no commit message supplied")
	}

	splitHeader := strings.SplitN(header, ":", 2)
	if len(splitHeader) != 2 {
		return cm, errors.New("header must be of the form <type>: <description>")
	} else {
		commitType = splitHeader[0]
		description = splitHeader[1]
	}

	if len(sections) > 1 {
		body = strings.Split(sections[1], "\n")
		for index, element := range(body) {
			body[index] = strings.TrimSpace(element)
		}
	}

	if len(sections) > 2 {
		footer = strings.Split(sections[2], "\n")
		for index, element := range(footer) {
			footer[index] = strings.TrimSpace(element)
		}
	}

	cm.Type = strings.TrimSpace(commitType)
	cm.Description = strings.TrimSpace(description)
	cm.Body = body
	cm.Footer = footer

	return cm, nil
}

//Function to check that a tag is in the list of approved tags.
func HasValidTag(cm CommitMessage) bool {
	commitType := cm.Type
	found := false
	for _, item := range ValidTags {
		if commitType == item {
			found = true
			break
		}
	}
	return found
}