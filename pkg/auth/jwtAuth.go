package auth

import (
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

func init() {
	os.Setenv("JWT_SECRET", "bolaJabeNah123445")
}

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWTAndSetCookie(w http.ResponseWriter, user *models.User) error {
	str, err := GenerateJWT(user)
	if err != nil {
		return err
	}
	// set token on cookies storage
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: str})
	return nil
}

func GenerateJWT(user *models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    user.Email,
		Username: user.FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err = token.SignedString(jwtKey)
	return
}

func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("jwt")
		if err != nil {
			fmt.Println(err)
			return
		}
		token := cookie.Value

		if err := ValidateToken(token); err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
		} else {
			next.ServeHTTP(writer, request)
		}
	})
}

func GetUserFromXCookieJWT(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		fmt.Println(err)
		return
	}
	cookie.Expires = time.Now()
	http.SetCookie(w, cookie)
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			jwtKey := []byte(os.Getenv("JWT_SECRET"))
			return jwtKey, nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return fmt.Errorf("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return fmt.Errorf("token expired")
	}
	return nil
}
