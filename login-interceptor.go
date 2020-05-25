package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/elgs/jwt"
	"github.com/elgs/wsl"
)

type LoginInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LoginInterceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	context map[string]interface{},
	wslApp *wsl.WSL) error {
	params["case"] = "lower"
	return nil
}

func (this *LoginInterceptor) After(tx *sql.Tx, result map[string]interface{},
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	data, ok := result["data"].([]interface{})
	if !ok {
		return errors.New("No data is returned.")
	}

	if u, ok := data[0].([]map[string]string); ok && len(u) > 0 {
		log.Printf("Login succeeded (%v)", u[0]["username"])

		mapClaims := map[string]interface{}{
			"user_id": u[0]["user_id"],
			"id":      u[0]["id"],
			"mode":    u[0]["mode"],
		}

		userId := u[0]["user_id"]
		sessionKey, err := getSessionKey(tx, userId)
		if err != nil {
			return err
		}

		expMinutes := wslApp.Config.App["session_expire_in_minutes"].(float64)
		token, err := jwt.Encode(mapClaims, int(expMinutes)*60, sessionKey)
		if err != nil {
			return err
		}

		result["token"] = token
	} else {
		log.Println("Login failed.")
		return nil
	}
	return nil
}
