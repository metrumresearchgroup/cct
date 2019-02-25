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

func TestHasValidTag_InvalidTagIsRejected(t *testing.T) {
	cm := cc.CommitMessage {
		Type:  "badtag",
		Body: []string{ "This is a body" },
		Footer: []string{ "This is a footer"},
		Description: "This is a description.",
	}

	_, err := cc.HasValidTag(cm)
	if(err == nil) {
		t.Fail()
	}

}

func TestValidTag_EmptyTagIsRejected(t *testing.T) {
	cm := cc.CommitMessage {
		Type:  "",
		Body: []string{ "This is a body" },
		Footer: []string{ "This is a footer"},
		Description: "This is a description.",
	}
	_, err := cc.HasValidTag(cm)
	if(err == nil) {
		t.Fail()
	}
}

func TestValidTag_NilTagRejected(t *testing.T) {
	cm := cc.CommitMessage {
		Type:  nil,
		Body: []string{ "This is a body" },
		Footer: []string{ "This is a footer"},
		Description: "This is a description.",
	}
	_, err := cc.HasValidTag(cm)
	if(err == nil) {
		t.Fail()
	}
}
