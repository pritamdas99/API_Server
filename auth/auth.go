package auth

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var secretKey []byte

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "pritam"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HasJWT(r *http.Request) (bool, error) {
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
			}
			return secretKey, nil
		})
		if err != nil {
			fmt.Println(err)
			return false, nil
		}
		if token.Valid {
			return true, nil
		}
	}
	return false, nil
}

func ValidateCookie(value string) (bool, error) {
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return secretKey, nil
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, nil
}
func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		//has cookies?
		for _, cookie := range request.Cookies() {
			if cookie.Name == "Token" {
				ok, err := ValidateCookie(cookie.Value)
				if err != nil {
					log.Println(err)
				}
				if ok {
					log.Println("serve because of cookies")
					next.ServeHTTP(response, request)
					return
				}
			}
		}
		json.NewEncoder(response).Encode(request.Cookies())
		//has token?
		ok, err := HasJWT(request)
		if err != nil {
			log.Fatal("there was a error find" +
				"ing token")
		}
		if ok {
			log.Println("serve because of has jwt token")
			next.ServeHTTP(response, request)
			return
		}
		ress := "NO access available for unknown user"
		json.NewEncoder(response).Encode(ress)
		response.WriteHeader(http.StatusUnauthorized) //unauthorized
		return
	}
}

func init() {
	secretKey = []byte("SecretKey")
}
