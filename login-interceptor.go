package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elgs/wsl"
)

type LoginInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LoginInterceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	context map[string]interface{},
	w http.ResponseWriter,
	r *http.Request, wslApp *wsl.WSL) error {
	params["case"] = "lower"
	return nil
}

func (this *LoginInterceptor) After(tx *sql.Tx, result *[]interface{},
	context map[string]interface{},
	w http.ResponseWriter,
	r *http.Request, wslApp *wsl.WSL) error {
	if u, ok := (*result)[0].([]map[string]string); ok && len(u) > 0 {
		log.Println("Login succeeded.")
		mapClaims := jwt.MapClaims{}
		for k, v := range u[0] {
			mapClaims[k] = v
		}

		expMinutes := wslApp.Config.App["session_expire_in_minutes"].(float64)
		mapClaims["exp"] = time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()
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
		w.Header().Add("token", tokenString)
	} else {
		log.Println("Login failed.")
		return nil
	}
	return nil
}
