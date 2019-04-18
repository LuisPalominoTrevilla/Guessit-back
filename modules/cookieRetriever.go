package modules

import (
	"net/http"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

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
