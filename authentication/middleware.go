package authentication

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/LuisPalominoTrevilla/Guessit-back/redis"
)

// Middleware serves as a wrapper for middleware resources
type Middleware struct {
	RedisClient *redis.Client
}

// AccessControl is a middleware that first checks for a token in the request Header and then if found and valid, puts it into the Header and continues execution
func (mid *Middleware) AccessControl(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.Fields(r.Header.Get("Authorization"))
		if len(auth) > 1 && auth[0] == "Bearer" {
			// Found token

			// Check that token is not in redis
			loggedOut, _ := mid.RedisClient.ExistsKey("blacklist:" + auth[1])

			if loggedOut {
				w.WriteHeader(401)
				// TODO: Change this to return a JSON object
				fmt.Fprintf(w, "Invalid token")
				return
			}

			claims, err := VerifyJWT(auth[1])

			if err != nil {
				w.WriteHeader(401)
				// TODO: Change this to return a JSON object
				fmt.Fprintf(w, "Invalid token")
			} else {
				userID := claims["userId"].(string)
				exp := claims["exp"].(float64)
				expS := strconv.FormatFloat(exp, 'E', -1, 64)
				r.Header.Add("uid", userID)
				r.Header.Add("exp", expS)
				r.Header.Add("token", auth[1])
				next.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(401)
			// TODO: Change this to return a JSON object
			fmt.Fprintf(w, "Invalid token")
		}
	})
}
