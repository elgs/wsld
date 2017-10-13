package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	headers map[string]string,
	wslApp *wsl.WSL) error {
	authHeader := headers["Authorization"]
	if authHeader != "" {
		s := strings.Split(authHeader, " ")
		if len(s) == 2 {
			tokenString := s[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				userId := token.Claims.(jwt.MapClaims)["user_id"]
				if userId == nil {
					return nil, errors.New("Invalid token")
				}

				sessionKey, err := getSessionKey(tx, userId.(string))
				if err != nil {
					return nil, err
				}

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
		}
	}
	return nil
}
func (this *AuthInterceptor) After(tx *sql.Tx, result *[]interface{}, wslApp *wsl.WSL) error {
	return nil
}
func (this *AuthInterceptor) OnError(err *error) error {
	return *err
}
