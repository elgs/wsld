package main

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/elgs/gostrgen"
	"github.com/elgs/wsl"
)

type ForgetPasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ForgetPasswordInterceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	context map[string]interface{},
	wslApp *wsl.WSL) error {
	vCode, err := gostrgen.RandGen(8, gostrgen.LowerUpperDigit, "", "lO") // exclude small L and big O
	if err != nil {
		return err
	}
	*script = strings.Replace(*script, "$recovering-password$", vCode, 1)
	params["case"] = "lower"
	return nil
}

func (this *ForgetPasswordInterceptor) After(tx *sql.Tx, result map[string]interface{},
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	data, ok := result["data"].([]interface{})
	if !ok {
		return errors.New("No data is returned.")
	}

	if len(data) == 0 {
		return errors.New("Failed get user information")
	}
	if userData, ok := data[0].([]map[string]string); ok {
		if len(userData) == 0 {
			return errors.New("Failed get user information")
		}
		email := userData[0]["email"]
		vCode := userData[0]["v_code"]
		err := wslApp.SendMail(
			wslApp.Config.App["mail_from"].(string), "Password Reset Verification Code", vCode, email)
		if err != nil {
			return err
		}
		delete(userData[0], "v_code")
	} else {
		return errors.New("Failed get user information")
	}
	return nil
}