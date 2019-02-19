package authentication

import (
	"fmt"
	"net/http"
	"strings"
)

// AccessControl is a middleware that first checks for a token in the request Header and then if found and valid, puts it into the Header and continues execution
func AccessControl(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.Fields(r.Header.Get("Authorization"))
		if len(auth) > 1 && auth[0] == "Bearer" {
			// Found token
			userID, err := VerifyJWT(auth[1])

			if err != nil {
				w.WriteHeader(401)
				// TODO: Change this to return a JSON object
				fmt.Fprintf(w, "Invalid token")
			} else {
				r.Header.Add("uid", userID)
				next.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(401)
			// TODO: Change this to return a JSON object
			fmt.Fprintf(w, "Invalid token")
		}
	})
}
