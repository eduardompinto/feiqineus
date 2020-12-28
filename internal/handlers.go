package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"
const applicationText = "application/text"

// SuspiciousMessageHandler handler to receive suspicious news
type SuspiciousMessageHandler struct {
	replier *Replier
}

// SimpleResponse simple response body with the Message as string.
type SimpleResponse struct {
	Message string `json:"message"`
}

// NewSuspiciousReceiver constructor to keep db unexported
func NewSuspiciousReceiver(replier *Replier) *SuspiciousMessageHandler {
	return &SuspiciousMessageHandler{replier: replier}
}

// ServeHTTP handle POST requests, hashes the content and store it on db
func (h *SuspiciousMessageHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		h.handlePost(wr, req)
	case http.MethodGet:
		log.Println("Do nothing")
	default:
		return
	}
}

func (h *SuspiciousMessageHandler) handlePost(wr http.ResponseWriter, req *http.Request) {
	defer closeReq(req)
	wr.Header().Add(contentType, applicationJSON)

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeSimpleResponse(wr, "can't read message body", 500)
		return
	}

	var msg SuspiciousMessage
	if err = json.Unmarshal(b, &msg); err != nil {
		writeSimpleResponse(wr, "Can't parse the request body", 500)
		return
	}
	if !msg.IsValid() {
		writeSimpleResponse(wr, "Invalid request body", 400)
		return
	}
	vm, err := h.replier.CheckMessage(msg)
	if err != nil {
		log.Println(err)
		writeSimpleResponse(wr, fmt.Sprintf("Can't check message, %v", err), 500)
	} else {
		writeSimpleResponse(wr, vm.Verdict(), 200)
	}

}

func writeSimpleResponse(wr http.ResponseWriter, msg string, status int) {
	wr.WriteHeader(status)
	r, err := json.Marshal(SimpleResponse{Message: msg})
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
