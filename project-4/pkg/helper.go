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

func GenerateToken(id int, email string, role string) (string, Error) {
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

func VerifyToken(context *gin.Context) (interface{}, Error) {
	headerToken := context.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, Unauthorized("You are not logged in. Please log in to access this resource.")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("you are not logged in. Please log in to access this resource")
		}

		return []byte(secretKey), nil
	})

	verifiedToken, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, Unauthorized("You are not logged in. Please log in to access this resource.")
	}

	return verifiedToken, nil
}

func GetIdParam(context *gin.Context, paramName string) (int, Error) {
	id, err := strconv.Atoi(context.Param(paramName))

	if err != nil {
		return int(0), BadRequest("Invalid id params")
	}

	return int(id), nil
}

func HashPassword(p string) (string, Error) {
	salt := 12
	password := []byte(p)
	hashedPass, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return "", InternalServerError("Failed to hash the password")
	}

	return string(hashedPass), nil
}

func ComparePassword(p string, h string) bool {
	password, hashedPassword := []byte(p), []byte(h)

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err == nil
}
