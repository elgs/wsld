package main

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/elgs/gostrgen"
	"github.com/elgs/wsl"
)

type SignupInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *SignupInterceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	context map[string]interface{},
	wslApp *wsl.WSL) error {
	vCode, err := gostrgen.RandGen(8, gostrgen.LowerUpperDigit, "", "lO") // exclude small L and big O
	if err != nil {
		return err
	}
	*script = strings.Replace(*script, "$pending-for-verification$", vCode, 1)

	params["case"] = "lower"
	return nil
}

func (this *SignupInterceptor) After(
	tx *sql.Tx,
	params map[string]string,
	result map[string]interface{},
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	data, ok := result["data"].([]interface{})
	if !ok {
		return errors.New("No data is returned.")
	}

	if len(data) == 0 {
		return errors.New("Failed to sign up.")
	}
	if userData, ok := data[0].([]map[string]string); ok {
		if len(userData) == 0 {
			return errors.New("Failed to sign up.")
		}
		email := userData[0]["email"]
		vCode := userData[0]["v_code"]
		err := wslApp.SendMail(
			wslApp.Config.App["mail_from"].(string), "New Account Verification Code", vCode, email)
		if err != nil {
			return err
		}
		delete(userData[0], "v_code")
	} else {
		return errors.New("Failed to sign up.")
	}
	return nil
}
