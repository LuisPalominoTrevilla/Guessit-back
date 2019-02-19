package authentication

import (
	"fmt"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"

	"github.com/LuisPalominoTrevilla/Guessit-back/errors"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT Generates jwt token
func GenerateJWT(user *models.User) (string, error) {
	signInKey := []byte(os.Getenv("SECRET_JWT_KEY"))

	type CustomClaims struct {
		UserID primitive.ObjectID `json:"userId"`
		jwt.StandardClaims
	}

	claims := CustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "GuessIt",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signInKey)
	if err != nil {
		fmt.Println("There was an error ", err.Error())
		return "", err
	}

	return tokenString, nil
}

// VerifyJWT verifies jwt and sends back claims (user_id) or error
func VerifyJWT(token string) (string, error) {
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	if parsedToken.Valid {
		return claims["userId"].(string), nil
	}
	return "", errors.New("Token not valid")
}
