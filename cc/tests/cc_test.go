package tests

import (
	"github.com/metrumresearchgroup/cct/cc"
	"testing"
)

func TestHasValidTag_ValidTagPasses(t *testing.T) {
	//testMessage := "fix: this is a valid message."
	cm := cc.CommitMessage {
		Type:  "fix",
		Body: []string{ "This is a body" },
		Footer: []string{ "This is a footer"},
		Description: "This is a description.",
	}
	_, err := cc.HasValidTag(cm)
	if(err != nil) {
		t.Fail()
	}
}