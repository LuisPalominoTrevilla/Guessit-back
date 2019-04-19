package modules

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//RetrieveRatedFromCookie Returns all rated image ObjectIds from a cookie. If the cookie does not exist, or it doesn't have a good format, the array will be empty
func RetrieveRatedFromCookie(cookieName string, r *http.Request) []primitive.ObjectID {
	var ratedImages []primitive.ObjectID = []primitive.ObjectID{}
	ratedCookie, err := r.Cookie(cookieName)

	if err != nil {
		return ratedImages
	}

	value := ratedCookie.Value
	rawIds := strings.Split(value, ",")
	for i := range rawIds {
		oid, err := primitive.ObjectIDFromHex(rawIds[i])
		if err == nil {
			ratedImages = append(ratedImages, oid)
		}
	}
	return ratedImages
}

// AddCookieValue adds a value to a cookie. If cookie doesn't exist, then it is created
func AddCookieValue(cookieName string, newValue string, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookies())
	ratedCookie, err := r.Cookie(cookieName)
	var chips *http.Cookie

	if err != nil {
		chips = &http.Cookie{
			Name:   cookieName,
			Value:  newValue + ",",
			Path:   "/",
			MaxAge: 0,
		}
		http.SetCookie(w, chips)
	} else {
		chips = &http.Cookie{
			Name:   cookieName,
			Value:  ratedCookie.Value + newValue + ",",
			Path:   "/",
			MaxAge: 0,
		}
	}
	http.SetCookie(w, chips)
}
