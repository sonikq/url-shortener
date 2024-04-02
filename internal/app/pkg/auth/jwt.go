package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

const (
	CookieName      = "token"
	SecretKey       = "secret"
	TokenExpiration = 60 * time.Minute
)

func generateJWTString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiration)),
		},
		UserID: uuid.NewString(),
	})

	return token.SignedString([]byte(SecretKey))
}

func generateCookie() (*http.Cookie, error) {
	token, err := generateJWTString()
	if err != nil {
		return nil, fmt.Errorf("jwt, generateCookie: %s", err.Error())
	}
	cookie := &http.Cookie{
		Name:  CookieName,
		Value: token,
		Path:  "/",
	}
	return cookie, nil
}

func GetUserToken(w http.ResponseWriter, r *http.Request) (string, error) {
	var (
		cookie *http.Cookie
		err    error
	)
	claims := &Claims{}

	cookie, _ = r.Cookie(CookieName)
	if cookie == nil {
		cookie, err = generateCookie()
		if err != nil {
			return "", err
		}
		http.SetCookie(w, cookie)
	} else {
		token, err := jwt.ParseWithClaims(cookie.Value, claims,
			func(j *jwt.Token) (interface{}, error) {
				if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, err
				}
				return []byte(SecretKey), nil
			})
		if claims.UserID == "" || !token.Valid || err != nil {
			return "", fmt.Errorf("cookie does not contain user id or cookie is invalid")
		}
	}

	token, err := jwt.ParseWithClaims(cookie.Value, claims,
		func(j *jwt.Token) (interface{}, error) {
			if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, err
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		cookie, err = generateCookie()
		http.SetCookie(w, cookie)
	}

	if !token.Valid {
		cookie, err = generateCookie()
		http.SetCookie(w, cookie)
	}

	return claims.UserID, nil
}
