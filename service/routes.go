package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func InitiallizeRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/getwords", GetWordCount).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func GetWordCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	resp, err := getMostUsedWords(string(input))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(resp)
}
