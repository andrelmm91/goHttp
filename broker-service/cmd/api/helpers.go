package main
package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func (App *config) readJSON(w http.ResponseWriter, r *http.Request, data any) error{
	maxBytes := 1048576; // 1Mb

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(Data)

	if err != nil {
		return err
	} 

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return err.New("Body must have a single JSON value")
	}

	return nil
}


func (App *config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error{
	out, err := json.Marshel(data)

	if err != nil {
		return err
	} 

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(out)

	if err != nil {
		return err
	} 

	return nil
}


func (App *config) errorJSON(w http.ResponseWriter, err error, status ...int) error{
	statusCode := StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.WriteJSON(w, statusCode, payload)
}
