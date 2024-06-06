package auth

import (
	"fmt"
	"github.com/sonikq/url-shortener/internal/app/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims -
type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

// Константы для пакета auth.
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

// SetUserCookie -
func SetUserCookie(w http.ResponseWriter) error {
	cookie, err := generateCookie()
	if err != nil {
		return err
	}
	http.SetCookie(w, cookie)
	return nil
}

// VerifyUserToken -
func VerifyUserToken(w http.ResponseWriter, r *http.Request) (string, error) {
	var (
		cookie *http.Cookie
		err    error
	)
	claims := &Claims{}

	cookie, err = r.Cookie(CookieName)
	if cookie == nil || err != nil {
		return "", err
	}

	token, parseCookieErr := parseCookie(cookie.Value, claims)
	if parseCookieErr != nil {
		return "", parseCookieErr
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid cookie")
	}

	if claims.UserID == "" {
		return "", fmt.Errorf("user_id is empty - invalid")
	}

	return claims.UserID, nil
}

// GetUserToken -
func GetUserToken(w http.ResponseWriter, r *http.Request) (string, error) {
	var (
		cookie *http.Cookie
		err    error
	)
	claims := &Claims{}

	cookie, err = r.Cookie(CookieName)
	if cookie == nil || err != nil {
		cookie, err = generateCookie()
		if err != nil {
			return "", err
		}
		http.SetCookie(w, cookie)
	}

	token, parseCookieErr := parseCookie(cookie.Value, claims)
	if parseCookieErr != nil {
		cookie, err = generateCookie()
		if err != nil {
			return "", models.ErrGenerateCookie
		}
		http.SetCookie(w, cookie)
	}

	if !token.Valid {
		cookie, err = generateCookie()
		if err != nil {
			return "", models.ErrGenerateCookie
		}
		http.SetCookie(w, cookie)
	}

	return claims.UserID, nil
}

func parseCookie(value string, claim *Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(value, claim,
		func(j *jwt.Token) (interface{}, error) {
			if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		return nil, err
	}
	return token, nil
}
