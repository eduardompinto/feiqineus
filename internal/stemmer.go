package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Stemmer struct {
	host string
	port string
}

func NewStemmer() *Stemmer {
	return &Stemmer{
		host: EnvOrDefault("STEMMER_HOST", "localhost"),
		port: EnvOrDefault("STEMMER_PORT", "8000"),
	}
}

func (s *Stemmer) Normalize(text string) (string, error) {
	url := fmt.Sprintf("http://%s:%s/stemmer", s.host, s.port)
	resp, err := http.Post(url, "text/plain", strings.NewReader(text))
	if err != nil {
		return "", err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Printf("Can't close stemmer response body to text %s", text)
		}
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	return string(b), nil
}
