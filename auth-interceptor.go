package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elgs/gosqljson"
	"github.com/elgs/wsl"
)

var userKeys = make(map[string]string)

type AuthInterceptor struct {
	*wsl.DefaultInterceptor
}

func getSessionKey(tx *sql.Tx, userId string) ([]byte, error) {
	var sessionKey string
	if val, ok := userKeys[userId]; ok {
		sessionKey = val
	} else {
		dbResult, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT SESSION_KEY FROM USER WHERE ID=?", userId)
		if err != nil {
			return nil, err
		}
		if len(dbResult) != 1 {
			return nil, errors.New("User not found.")
		}
		sessionKey = dbResult[0]["session_key"]
		if sessionKey == "" {
			return nil, errors.New("Session key is empty.")
		}
		userKeys[userId] = sessionKey
	}
	return []byte(sessionKey), nil
}

func (this *AuthInterceptor) Before(
	tx *sql.Tx,
	script *string,
	params map[string]string,
	context map[string]interface{},
	wslApp *wsl.WSL) error {
	authHeader := context["Authorization"]
	if authHeader != nil {
		s := strings.Split(authHeader.(string), " ")
		if len(s) == 2 {
			tokenString := s[1]
			needToRenewToken := false
			var sessionKey []byte
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				exp := token.Claims.(jwt.MapClaims)["exp"]
				expInSeconds := int64(exp.(float64)) - time.Now().Unix()
				sessionSeconds := 60 * int64(wslApp.Config.App["session_expire_in_minutes"].(float64))
				needToRenewToken = sessionSeconds > expInSeconds*2

				userId := token.Claims.(jwt.MapClaims)["user_id"]
				if userId == nil {
					return nil, errors.New("Invalid token")
				}

				_sessionKey, err := getSessionKey(tx, userId.(string))
				if err != nil {
					return nil, err
				}
				sessionKey = _sessionKey
				return sessionKey, nil
			})
			if err != nil {
				return err
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				params["__session_id"] = fmt.Sprintf("%v", claims["id"])
				params["__user_id"] = fmt.Sprintf("%v", claims["user_id"])
				params["__user_mode"] = fmt.Sprintf("%v", claims["mode"])
			} else {
				return errors.New("Authentication failed.")
			}

			if needToRenewToken {
				// fmt.Println("Renew token")
				// fmt.Println(sessionKey)

				expMinutes := wslApp.Config.App["session_expire_in_minutes"].(float64)
				token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()

				tokenString, err := token.SignedString([]byte(sessionKey))
				if err != nil {
					return err
				}
				context["token"] = tokenString

				// fmt.Println(tokenString)
			}
		}
	}
	return nil
}
func (this *AuthInterceptor) After(tx *sql.Tx, result map[string]interface{},
	context map[string]interface{}, wslApp *wsl.WSL) error {
	return nil
}
func (this *AuthInterceptor) OnError(err *error) error {
	return *err
}
