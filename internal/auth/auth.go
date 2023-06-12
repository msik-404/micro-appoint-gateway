package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var Secret = os.Getenv("SECRET")

func HashAndSalt(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func AreEqual(hashedPwd []byte, plainPwd []byte) (bool, error) {
	byteHash := hashedPwd
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CreateJWT(
    userID string, 
) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["exp"] = time.Now().Add(time.Hour).Unix()
    claims["user_id"] = userID

    tokenStr, err := token.SignedString([]byte(Secret))
    if err != nil {
        return "", err
    }
    return tokenStr, err
}
