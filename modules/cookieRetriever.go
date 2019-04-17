package modules

import (
	"net/http"

	_ "github.com/mongodb/mongo-go-driver/bson/primitive"
)

func RetrieveRatedFromCookie(cookieName string, r *http.Request) (string, error) {
	ratedCookie, err := r.Cookie(cookieName)

	if err != nil {
		return "", err
	}

	value := ratedCookie.Value
	return value, nil
}
