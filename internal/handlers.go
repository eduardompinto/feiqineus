package internal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"
const applicationText = "application/text"

// SuspiciousReceiver handler to receive suspicious news
type SuspiciousReceiver struct {
	db Database
}

// NewSuspiciousReceiver constructor to keep db unexported
func NewSuspiciousReceiver(db Database) *SuspiciousReceiver {
	return &SuspiciousReceiver{db: db}
}

// ServeHTTP handle POST requests, hashes the content and store it on db
func (h *SuspiciousReceiver) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		h.handlePost(wr, req)
	case http.MethodGet:
		log.Println("Do nothing")
	default:
		return
	}
}

func (h *SuspiciousReceiver) handlePost(wr http.ResponseWriter, req *http.Request) {
	defer closeReq(req)
	wr.Header().Add(contentType, applicationJSON)

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeSimpleResponse(wr, "can't read message body", 500)
		return
	}

	var msg SuspiciousMessage
	ok := parse(wr, b, &msg)
	if ok {
		writeSimpleResponse(wr, "Message received", 200)
	}
}

func parse(wr http.ResponseWriter, b []byte, v ValidStruct) bool {
	err := json.Unmarshal(b, v)
	if err != nil {
		writeSimpleResponse(wr, "Can't parse the request body", 500)
		return false
	}
	if !v.IsValid() {
		writeSimpleResponse(wr, "Invalid request body", 400)
		return false
	}
	return true
}

func writeSimpleResponse(wr http.ResponseWriter, msg string, status int) {
	wr.WriteHeader(status)
	r, err := json.Marshal(Response{Message: msg})
	if err == nil {
		_, err = wr.Write(r)
		if err != nil {
			log.Println("can't write response body")
		}
		return
	}
	log.Println("can't marshal json response")
	wr.Header().Add(contentType, applicationText)
	_, err = io.WriteString(wr, msg)
	if err != nil {
		log.Println("can't write response body")
	}
}

func closeReq(req *http.Request) {
	err := req.Body.Close()
	if err != nil {
		log.Println("error during close on request body")
	}
}
