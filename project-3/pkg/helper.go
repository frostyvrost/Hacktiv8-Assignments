package pkg

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "seed"

func GenerateToken(id uint, email string, role string) (string, Error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", InternalServerError("Failed to generate token")
	}

	return signedToken, nil
}

func VerifyToken(context *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := context.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	verifiedToken, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, errResponse
	}

	return verifiedToken, nil
}

func HashPass(p string) (string, Error) {
	salt := 12
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return "", InternalServerError("Failed to hash the password")
	}

	return string(hash), nil
}

func ComparePass(h string, p string) bool {
	pass, hash := []byte(p), []byte(h)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}

func GetIdParam(context *gin.Context, idName string) (uint, Error) {
	id, err := strconv.Atoi(context.Param(idName))

	if err != nil {
		return uint(0), BadRequest("Invalid id params")
	}

	return uint(id), nil
}
