package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Stemmer struct {
	host string
	port string
}

func envOrDefault(env string, def string) string {
	e, ok := os.LookupEnv(env)
	if ok {
		return e
	}
	return def
}

func NewStemmer() *Stemmer {
	return &Stemmer{
		host: envOrDefault("STEMMER_HOST", "localhost"),
		port: envOrDefault("STEMMER_PORT", "8081"),
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
