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

	result := cc.HasValidTag(cm)
	if(result == false) {
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

	result := cc.HasValidTag(cm)
	if(result == true) {
		t.Fail()
	}

}

func TestHasValidTag_EmptyTagIsRejected(t *testing.T) {
	cm := cc.CommitMessage {
		Type:  "",
		Body: []string{ "This is a body" },
		Footer: []string{ "This is a footer"},
		Description: "This is a description.",
	}
	result := cc.HasValidTag(cm)
	if(result == true) {
		t.Fail()
	}
}

func TestHasValidTag_NilTagRejected(t *testing.T) {
	t.Skip()
	cm := cc.CommitMessage {
		//Type:  nil,
		Body: []string{ "This is a body" },
		Footer: []string{ "This is a footer"},
		Description: "This is a description.",
	}
	result := cc.HasValidTag(cm)
	if(result == true) {
		t.Fail()
	}
}

func utilityCommitMessageEqual(cm1, cm2 cc.CommitMessage) bool {

	//Why doesn't it allow me to compare cm1/2 with nil?
	//if cm1 == nil || cm2 == nil {
	//	return false //We are claiming that nil objects are not equal.
	//}

	result := true //Assume true until proven otherwise.
	result = result && cm1.Type == cm2.Type
	result = result && cm1.Description == cm2.Description
	result = result && len(cm1.Body) == len(cm2.Body)
	result = result && len(cm1.Footer) == len(cm2.Footer)

	//We need to bail here if the lists are of different lengths, otherwise we'll run into problems during the looping.
	if result == false {
		return result
	}

	for i := 0; i < len(cm1.Body); i++ {
		result = result && cm1.Body[i] == cm2.Body[i]
	}

	for i := 0; i < len(cm1.Footer); i++ {
		result = result && cm2.Footer[i] == cm2.Footer[i]
	}

	return result
}

func TestParseCommitMessage_SimpleMessage(t *testing.T) {
	messageRaw := "fix: this is a valid message."
	expected := cc.CommitMessage{
		Type:        "fix",
		Description: "this is a valid message",
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_FullMessage(t *testing.T) {
	messageRaw :=
`fix: this is a description.

this is a body1
this is a body2

this is a footer1
this is a footer2`

	expected := cc.CommitMessage {
		Type:  "fix",
		Description: "this is a description.",
		Body: []string{ "this is a body1", "this is a body2" },
		Footer: []string{ "this is a footer1", "this is a footer2"},
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_ExtraDescriptionFails(t *testing.T) {
	messageRaw :=
		`fix: this is a description1.
this is a description2

this is a body1
this is a body2

this is a footer1
this is a footer2`
		_, err := cc.ParseCommitMessage(messageRaw)
		if(err == nil) {
			t.Fail()
		}
}

func TestParseCommitMessage_ExtraTypeIsTreatedAsDescription(t *testing.T) {
	messageRaw :=
`fix: add: this is a description.

this is a body1
this is a body2

this is a footer1
this is a footer2`

	expected := cc.CommitMessage {
		Type:  "fix",
		Description: "add: this is a description.",
		Body: []string{ "this is a body1", "this is a body2" },
		Footer: []string{ "this is a footer1", "this is a footer2"},
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_NoBodyStillAllowsFooter(t *testing.T) {
	messageRaw :=
		`fix: add: this is a description.


this is a footer1
this is a footer2`

	expected := cc.CommitMessage {
		Type:  "fix",
		Description: "add: this is a description.",
		Body: []string{},
		Footer: []string{ "this is a footer1", "this is a footer2"},
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_NoFooter(t *testing.T) {
	messageRaw :=
		`fix: add: this is a description.

this is a body1
this is a body2`

	expected := cc.CommitMessage {
		Type:  "fix",
		Description: "add: this is a description.",
		Body: []string{ "this is a body1", "this is a body2" },
		Footer: []string{},
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_BlankLinesIgnored(t *testing.T) {
	messageRaw :=
		`fix: add: this is a description.

this is a body1
this is a body2

this is a footer1
this is a footer2

`

	expected := cc.CommitMessage {
		Type:  "fix",
		Description: "add: this is a description.",
		Body: []string{ "this is a body1", "this is a body2" },
		Footer: []string{ "this is a footer1", "this is a footer2"},
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_BlankLinesIgnoredNoFooter(t *testing.T) {
	messageRaw :=
		`fix: add: this is a description.

this is a body1
this is a body2



`

	expected := cc.CommitMessage {
		Type:  "fix",
		Description: "add: this is a description.",
		Body: []string{ "this is a body1", "this is a body2" },
		Footer: []string{ },
	}
	actual, err := cc.ParseCommitMessage(messageRaw)
	if ( !utilityCommitMessageEqual(expected, actual)) || err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_EmptyType(t *testing.T) {
	messageRaw :=
		`: this is a description.

this is a body1
this is a body2

this is a footer1
this is a footer2`

	_, err := cc.ParseCommitMessage(messageRaw)
	if err != nil {
		t.Fail()
	}
}

func TestParseCommitMessage_NoType(t *testing.T) {
	messageRaw :=
		`this is a description.

this is a body1
this is a body2

this is a footer1
this is a footer2`

	_, err := cc.ParseCommitMessage(messageRaw)
	if err != nil {
		t.Fail()
	}
}