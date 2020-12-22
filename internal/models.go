package internal

// ValidStruct if it's possible to validate the struct
type ValidStruct interface {
	IsValid() bool
}

// To be exposed on http handlers
// SimpleResponse simple response body with the Message as string.
type SimpleResponse struct {
	Message string `json:"message"`
}

// SuspiciousMessage is possible fake news.
// it contains the text (like a whats app message) and the link to follow it
type SuspiciousMessage struct {
	Text *string `json:"message"`
	Link *string `json:"link"`
}

// IsValid if Text and Link are present
func (sm *SuspiciousMessage) IsValid() bool {
	return sm.Text != nil && sm.Link != nil
}
