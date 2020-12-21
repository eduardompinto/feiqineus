package internal

// ValidStruct if it's possible to validate the struct
type ValidStruct interface {
	IsValid() bool
}

// Response simple response body with the Message field.
// To be exposed on http handlers
type Response struct {
	Message string `json:"message"`
}

// SuspiciousMessage is a request body
type SuspiciousMessage struct {
	Message *string `json:"message"`
	Link    *string `json:"link"`
}

// IsValid if Message and Link are present
func (sm *SuspiciousMessage) IsValid() bool {
	return sm.Message != nil && sm.Link != nil
}
