package internal

type ValidStruct interface {
	IsValid() bool
}

type Response struct {
	Message string `json:"message"`
}

type SuspiciousMessage struct {
	Message *string `json:"message"`
	Link    *string `json:"link"`
}

func (sm *SuspiciousMessage) IsValid() bool {
	return sm.Message != nil && sm.Link != nil
}
