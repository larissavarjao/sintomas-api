package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/larissavarjao/sintomas-api/api/pacient"
)

func PacientsHandlers(r *mux.Router, n *negroni.Negroni, service pacient.UseCase) {
	r.Handle("/pacients", n.With(
		negroni.Wrap(getAllPacients(service)),
	)).Methods("GET", "OPTIONS")
}

func getAllPacients(service pacient.UseCase) http.Handler {
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