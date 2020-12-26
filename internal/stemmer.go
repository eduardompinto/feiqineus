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

func getStemmerHost() string {
	h, ok := os.LookupEnv("STEMMER_HOST")
	if ok {
		return h
	}
	return "localhost"
}

func getStemmerPort() string {
	p, ok := os.LookupEnv("STEMMER_PORT")
	if ok {
		return p
	}
	return "8081"
}

func NewStemmer() *Stemmer {
	return &Stemmer{
		host: getStemmerHost(),
		port: getStemmerPort(),
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
			log.Printf("Can't close stemmer response body to text %s\n", text)
		}
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	return string(b), nil
}
