package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Authenticate(t *testing.T) {

	postBody := map[string]interface{}{
		"email": "a@a.com",
		"password": "verysecret",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))

	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("expected http response OK, but got %d", rr.Code)
	}
}