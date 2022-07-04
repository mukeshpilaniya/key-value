package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type Payload struct {
	Message string `json:"message,omitempty"`
	Error   bool   `json:"error,omitempty"`
}
type KeyStore struct {
	Key   string
	Value string
}

func (app *Application) get(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	var p Payload
	if key == "" {
		p.Error = true
		p.Message = "key string is empty"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		message, err := app.db.GetValue(key)
		if err != nil {
			p.Error = true
			p.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			p.Error = false
			p.Message = message
			w.WriteHeader(http.StatusOK)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	out, _ := json.MarshalIndent(p, "", "\t")
	w.Write([]byte(out))
	return
}

func (app *Application) set(w http.ResponseWriter, r *http.Request) {
	var pair = KeyStore{}
	var p Payload
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&pair)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	if pair.Key == "" {
		p.Error = true
		p.Message = "key string is empty"
		w.WriteHeader(http.StatusBadRequest)
	} else if pair.Value == "" {
		p.Error = true
		p.Message = "value string is empty"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		app.db.AddKey(pair.Key, pair.Value)
		p.Message = "kay value successfully stored"
		p.Error = false
		w.WriteHeader(http.StatusOK)
	}
	out, _ := json.MarshalIndent(p, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(out))
	return
}

func (app *Application) search(w http.ResponseWriter, r *http.Request) {
	prefixString := r.URL.Query().Get("prefix")
	suffixString := r.URL.Query().Get("suffix")
	var keys string
	if prefixString != "" {
		k := app.db.GetAllKeys()
		for _, v := range k {
			if strings.HasPrefix(v, prefixString) {
				keys = keys + " " + v
			}
		}
	} else if suffixString != "" {
		k := app.db.GetAllKeys()
		for _, v := range k {
			if strings.HasSuffix(v, suffixString) {
				keys = keys + " " + v
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(keys))
	return
}
