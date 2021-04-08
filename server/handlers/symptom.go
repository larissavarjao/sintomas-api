package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/larissavarjao/sintomas-api/core/symptom"
)

func SymptomsHandlers(r *mux.Router, n *negroni.Negroni, s symptom.UseCase) {
	r.Handle("/symptoms", n.With(
		negroni.Wrap(getAllSymptoms(s)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/symptoms", n.With(
		negroni.Wrap(createSymptom(s)),
	)).Methods("POST")
}

func getAllSymptoms(service symptom.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		all, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error formating JSON"))
		}
	})
}

func createSymptom(service symptom.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var sy symptom.Symptom

		err := json.NewDecoder(r.Body).Decode(&sy)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		// TODO: VALIDACAO

		symptom, err := service.Create(&sy)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(symptom)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error formating JSON"))
		}

		w.WriteHeader(http.StatusCreated)
	})
}