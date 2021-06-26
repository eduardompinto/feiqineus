package internal

// Replier receive a SuspiciousMessage and
type Replier struct {
	verifiedMessageGetter VerifiedMessageGetter
	verifiedMessageSaver  VerifiedMessageSaver
	textNormalizer        TextNormalizer
}

// NewReplier creates a new instance of the Replier struct
func NewReplier(getter VerifiedMessageGetter, saver VerifiedMessageSaver) Replier {
	return Replier{
		verifiedMessageGetter: getter,
		verifiedMessageSaver:  saver,
		textNormalizer:        NewStemmer(),
	}
}

func (r *Replier) CheckMessage(sm SuspiciousMessage) (*VerifiedMessage, error) {
	var text = *sm.Text
	t, err := r.textNormalizer.Normalize(text)
	if err != nil {
		return nil, err
	}
	vm := r.verifiedMessageGetter.Get(t)
	if vm != nil {
		return vm, nil
	}
	vm, err = r.createVerifiedMessage(sm, t)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func (r *Replier) createVerifiedMessage(sm SuspiciousMessage, textNormalized string) (*VerifiedMessage, error) {
	vm := NewVerifiedMessage(*sm.Text, textNormalized, sm.Link)
	if err := r.verifiedMessageSaver.Save(&vm); err != nil {
		return nil, err
	}
	return &vm, nil
}
