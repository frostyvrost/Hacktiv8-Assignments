package models

import (
	"project-2/database"
	"project-2/pkg"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Photo struct {
	Id        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	Id        int
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SocialMedia struct {
	Id             int
	Name           string
	SocialMediaUrl string
	UserId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

var invalidToken = pkg.NewUnauthenticatedError("invalid token")

func (u *User) parseToken(tokenString string) (*jwt.Token, pkg.Error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidToken
		}

		return []byte(database.AppConfig().JwtSecretKey), nil
	})

	if err != nil {
		return nil, invalidToken
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claims jwt.MapClaims) pkg.Error {

	if id, ok := claims["id"].(float64); !ok {
		return invalidToken
	} else {
		u.Id = int(id)
	}

	if username, ok := claims["username"].(string); !ok {
		return invalidToken
	} else {
		u.Username = username
	}

	if email, ok := claims["email"].(string); !ok {
		return invalidToken
	} else {
		u.Email = email
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) pkg.Error {

	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidToken
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidToken
	}

	tokenString := splitToken[1]
	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidToken
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
		"email":    u.Email,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(database.AppConfig().JwtSecretKey))
	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()
	return u.signToken(claims)
}

func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashPassword)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
