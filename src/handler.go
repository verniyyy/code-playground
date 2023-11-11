package src

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Handler ...
type Handler struct{}

// RunGo ...
func (h Handler) RunGo(w http.ResponseWriter, r *http.Request) {
	// reqBody := r.Body
	// defer reqBody.Close()

	src, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	defer src.Close()

	app := NewApplication(
		NewCodeRunner(),
	)
	out, err := app.RunGo(r.Context(), Input{
		Src: src,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.ResponseJSON(w, http.StatusOK, out)
}

// RunRuby ...
func (h Handler) RunRuby(w http.ResponseWriter, r *http.Request) {
	// reqBody := r.Body
	// defer reqBody.Close()

	src, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	defer src.Close()

	app := NewApplication(
		NewCodeRunner(),
	)
	out, err := app.RunRuby(r.Context(), Input{
		Src: src,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.ResponseJSON(w, http.StatusOK, out)
}

// ResponseJSON ...
func (h Handler) ResponseJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.Encode(payload)
}

// DecodeJSON ...
func DecodeJSON[T any](data io.Reader) (T, error) {
	var v T
	err := json.NewDecoder(data).Decode(&v)
	if err != nil {
		var emptyReturn T
		return emptyReturn, err
	}
	return v, nil
}
