package authentication

import (
	"fmt"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT Generates jwt token
func GenerateJWT(user *models.User) (string, error) {
	signInKey := []byte(os.Getenv("SECRET_JWT_KEY"))

	type CustomClaims struct {
		UserID   primitive.ObjectID `json:"userId"`
		Username string             `json:"username"`
		jwt.StandardClaims
	}

	claims := CustomClaims{
		user.ID,
		user.Username,
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
