package cc

// CommitMessage stores the commit message
type CommitMessage struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Body        []string `json:"body"`
	Footer      []string `json:"footer"`
}
