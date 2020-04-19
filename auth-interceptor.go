package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/elgs/gosqljson"
	"github.com/elgs/jwt"
	"github.com/elgs/wsl"
)

var userKeys = make(map[string]string)

type AuthInterceptor struct {
	*wsl.DefaultInterceptor
}

func getSessionKey(tx *sql.Tx, userId string) (string, error) {
	if val, ok := userKeys[userId]; ok {
		return val, nil
	} else {
		dbResult, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT PASSWORD FROM USER WHERE ID=?", userId)
		if err != nil {
			return "", err
		}
		if len(dbResult) != 1 {
			return "", errors.New("User not found.")
		}
		sessionKey := dbResult[0]["password"]
		if sessionKey == "" {
			return "", errors.New("Session key is empty.")
		}
		userKeys[userId] = sessionKey
		return sessionKey, nil
	}
}

func (this *AuthInterceptor) Before(
	tx *sql.Tx,
	script *string,
	params map[string]string,
	context map[string]interface{},
	wslApp *wsl.WSL) error {
	if tokenString, ok := context["Authorization"].(string); ok {
		// needToRenewToken := false

		claims, err := jwt.Decode(tokenString)
		userId := claims["user_id"]
		sessionKey, err := getSessionKey(tx, userId.(string))
		verified, err := jwt.Verify(tokenString, sessionKey)
		if !verified || err != nil {
			return errors.New("Invalid token.")
		}

		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	exp := token.Claims.(jwt.MapClaims)["exp"]
		// 	expInSeconds := int64(exp.(float64)) - time.Now().Unix()
		// 	sessionSeconds := 60 * int64(wslApp.Config.App["session_expire_in_minutes"].(float64))
		// 	needToRenewToken = sessionSeconds > expInSeconds*2

		// 	userId := token.Claims.(jwt.MapClaims)["user_id"]
		// 	if userId == nil {
		// 		return nil, errors.New("Invalid token")
		// 	}

		// 	_sessionKey, err := getSessionKey(tx, userId.(string))
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	sessionKey = _sessionKey
		// 	return sessionKey, nil
		// })
		// if err != nil {
		// 	return err
		// }

		params["$$session_id"] = fmt.Sprintf("%v", claims["id"])
		params["$$user_id"] = fmt.Sprintf("%v", claims["user_id"])
		params["$$user_mode"] = fmt.Sprintf("%v", claims["mode"])

		// if needToRenewToken {
		// 	// fmt.Println("Renew token")
		// 	// fmt.Println(sessionKey)

		// 	expMinutes := wslApp.Config.App["session_expire_in_minutes"].(float64)
		// 	token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()

		// 	tokenString, err := token.SignedString([]byte(sessionKey))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	context["token"] = tokenString

		// 	// fmt.Println(tokenString)
		// }
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
