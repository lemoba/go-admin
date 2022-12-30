package controller

import (
	"encoding/json"
	"github.com/lemoba/go-admin/model"
	"log"
	"net/http"
)

func registerJsonRoutes() {
	http.HandleFunc("/json", handleJson)
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		company := model.Company{}
		err := dec.Decode(&company)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		err = enc.Encode(company)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
