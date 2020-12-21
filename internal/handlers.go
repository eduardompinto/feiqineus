package internal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const ContentType = "Content-Type"
const ApplicationJson = "application/json"
const ApplicationText = "application/text"

type SuspiciousReceiver struct{}

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
	wr.Header().Add(ContentType, ApplicationJson)

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
	wr.Header().Add(ContentType, ApplicationText)
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
