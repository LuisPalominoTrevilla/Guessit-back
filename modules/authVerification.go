package modules

import (
	"net/http"
	"strings"

	"github.com/LuisPalominoTrevilla/Guessit-back/authentication"
)

// IsAuthed returns if current user is authenticated. If it is, it also returns the userId as a string
func IsAuthed(r *http.Request) (bool, string) {
	auth := strings.Fields(r.Header.Get("Authorization"))
	if len(auth) > 1 && auth[0] == "Bearer" {
		claims, err := authentication.VerifyJWT(auth[1])

		if err == nil {
			userID := claims["userId"].(string)
			return true, userID
		}
	}
	return false, ""
}
