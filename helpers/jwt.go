package helpers

import (
	"nami_navigation_log/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var privateKeys = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWToken(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKeys)
}
