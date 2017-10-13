package main

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/elgs/gostrgen"
	"github.com/elgs/wsl"
)

type ForgetPassword1Interceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ForgetPassword1Interceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	headers map[string]string, wslApp *wsl.WSL) error {
	vCode, err := gostrgen.RandGen(8, gostrgen.LowerUpperDigit, "", "lO") // exclude small L and big O
	if err != nil {
		return err
	}
	*script = strings.Replace(*script, "$rp$", vCode, 1)
	params["case"] = "lower"
	return nil
}

func (this *ForgetPassword1Interceptor) After(tx *sql.Tx, result *[]interface{}, wslApp *wsl.WSL) error {
	if len(*result) == 0 {
		return errors.New("Failed get user information")
	}
	if userData, ok := (*result)[0].([]map[string]string); ok {
		if len(userData) == 0 {
			return errors.New("Failed get user information")
		}
		email := userData[0]["email"]
		vCode := userData[0]["v_code"]
		err := wslApp.SendMail(
			"config.Mail.MailFrom", "Password Reset Verification Code", vCode, email)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Failed get user information")
	}
	return nil
}
