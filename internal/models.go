package internal

import (
	"fmt"
	"time"
)

// ValidStruct if it's possible to validate the struct
type ValidStruct interface {
	IsValid() bool
}

// SuspiciousMessage is possible fake news.
// it contains the text (like a whats app message) and the link to follow it
type SuspiciousMessage struct {
	Text *string `json:"message"`
	Link *string `json:"link"`
}

// IsValid if Text and Link are present
func (sm *SuspiciousMessage) IsValid() bool {
	return sm.Text != nil
}

type VerifiedMessage struct {
	ID             int32
	Checked        bool
	Explanation    string
	FirstAppear    time.Time
	IsFake         *bool
	Link           string
	Text           string
	TextNormalized string
}

func (vm *VerifiedMessage) IsValid() bool {
	return true
}

func (vm *VerifiedMessage) Verdict() string {
	var verdict string
	if !vm.Checked {
		verdict = vm.getUncheckedVerdictMessage()
	} else if vm.IsFake == nil {
		verdict = vm.getInconclusiveVerdictMessage()
	} else if *vm.IsFake {
		verdict = vm.getFakeVerdictMessage()
	} else {
		verdict = vm.getGenuineVerdictMessage()
	}
	return verdict
}

func (vm *VerifiedMessage) getInconclusiveVerdictMessage() string {
	return fmt.Sprintf("Não consegui chegar a uma conclusão sobre isso. Veja o que eu sei sobre isso: %s", vm.Explanation)
}

func (vm *VerifiedMessage) getUncheckedVerdictMessage() string {
	return "Ainda nao verifiquei essa mensagem"
}

func (vm *VerifiedMessage) getGenuineVerdictMessage() string {
	return fmt.Sprintf(
		`A informação parece ser verdadeira, mas ainda assim leia o que sei sobre isso: %s`,
		vm.Explanation,
	)
}

func (vm *VerifiedMessage) getFakeVerdictMessage() string {
	return fmt.Sprintf(
		`Essa mensagem é falsa! Veja o que eu sei sobre isso: %s`,
		vm.Explanation,
	)
}

func NewVerifiedMessage(text string, textNormalized string, link *string) VerifiedMessage {
	var l = ""
	if link != nil {
		l = *link
	}
	return VerifiedMessage{
		Checked:        false,
		Explanation:    "",
		FirstAppear:    time.Now(),
		IsFake:         nil,
		Link:           l,
		Text:           text,
		TextNormalized: textNormalized,
	}
}

type VerifiedMessageGetter interface {
	Get(message string) *VerifiedMessage
}

type VerifiedMessageSaver interface {
	Save(vm *VerifiedMessage) error
}

type TextNormalizer interface {
	Normalize(text string) (string, error)
}
