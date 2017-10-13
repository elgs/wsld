package main

import (
	"database/sql"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elgs/wsl"
)

type LoginInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LoginInterceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	headers map[string]string, wslApp *wsl.WSL) error {
	params["case"] = "lower"
	return nil
}

func (this *LoginInterceptor) After(tx *sql.Tx, result *[]interface{}, wslApp *wsl.WSL) error {
	if u, ok := (*result)[0].([]map[string]string); ok && len(u) > 0 {
		log.Println("Login succeeded.")
		mapClaims := jwt.MapClaims{}
		for k, v := range u[0] {
			mapClaims[k] = v
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

		userId := u[0]["user_id"]
		sessionKey, err := getSessionKey(tx, userId)
		if err != nil {
			return err
		}

		tokenString, err := token.SignedString([]byte(sessionKey))
		if err != nil {
			return err
		}
		*result = append(*result, tokenString)
	} else {
		log.Println("Login failed.")
		return nil
	}
	return nil
}
