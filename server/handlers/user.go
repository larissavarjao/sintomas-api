package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/larissavarjao/sintomas-api/core/user"
)



func UsersHandlers(r *mux.Router, n *negroni.Negroni, service user.UseCase) {
	r.Handle("/users", n.With(
		negroni.Wrap(getAllUsers(service)),
	)).Methods("GET", "OPTIONS")
}

func getAllUsers(service user.UseCase) http.Handler {
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