package internal

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Replier: receive a SuspiciousMessage and
type Replier struct {
	db Database
}

// NewReplier creates a new instance of the Replier struct
func NewReplier(db Database) Replier {
	return Replier{db: db}
}

// HashMessage take the portuguese stemme and hashes it
func (r *Replier) HashMessage(sm SuspiciousMessage) (string, error) {
	st, err := getStemmedText(*sm.Text)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write([]byte(*st))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), nil
}

func getStemmedText(text string) (*string, error) {
	resp, err := http.Post("http://stemmer:8081/stemmer", "text/plain", strings.NewReader(text))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	b, _ := ioutil.ReadAll(resp.Body)
	s := string(b)
	return &s, nil
}
