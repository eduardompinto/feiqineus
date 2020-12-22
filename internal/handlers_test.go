package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSuspiciousReceiver_ServeHTTP(t *testing.T) {
	receiver := SuspiciousReceiver{}
	rb := strings.NewReader(`{"message": "example", "link": "http://google.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "localhost:8080/suspicious", rb)
	wr := httptest.NewRecorder()
	receiver.ServeHTTP(wr, req)
	var r SimpleResponse
	_ = json.Unmarshal(wr.Body.Bytes(), &r)
	expected := "Text received"
	if r.Message != expected {
		t.Errorf("Expected: [%s], Received: [%s]", expected, r.Message)
	}
}

func TestSuspiciousReceiver_ServeHTTP_Invalid_Body(t *testing.T) {
	receiver := SuspiciousReceiver{}
	rb := strings.NewReader(`{"msg": "example", "link": "http://google.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "localhost:8080/suspicious", rb)
	wr := httptest.NewRecorder()
	receiver.ServeHTTP(wr, req)
	var r SimpleResponse
	_ = json.Unmarshal(wr.Body.Bytes(), &r)
	expected := "Invalid request body"
	if r.Message != expected {
		t.Errorf("Expected: [%s], Received: [%s]", expected, r.Message)
	}
}
