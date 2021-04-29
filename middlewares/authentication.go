package middlewares

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

func ValidateAuthentication() negroni.Handler {
  return negroni.HandlerFunc(func (w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("Validation authentication")

		next(w, r)
	})
}