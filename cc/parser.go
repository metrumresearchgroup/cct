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
	var body []string
	var footer []string

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
		if len(strings.Split(description, "\n")) > 1 {
			return cm, errors.New("description must be a single-line")
		}
	}

	if len(sections) > 1 {
		body = strings.Split(sections[1], "\n")
		if body[0] == "" {
			footer = body
			body = []string{}
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
	cm.Body = pruneWhiteSpace(body)
	cm.Footer = pruneWhiteSpace(footer)

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

func pruneWhiteSpace(stringList []string) ([]string) {
	var returnList []string = []string{}
	for _, element := range(stringList) {
		if strings.TrimSpace(element) != "" {
			returnList = append(returnList, strings.TrimSpace(element))
		}
	}
	if returnList == nil {
		returnList = []string { }
	}
	return returnList
}