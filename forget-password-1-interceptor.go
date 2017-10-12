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

func (this *ForgetPassword1Interceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, config *wsl.Config) error {
	vCode, err := gostrgen.RandGen(8, gostrgen.LowerUpperDigit, "", "lO") // exclude small L and big O
	if err != nil {
		return err
	}
	*script = strings.Replace(*script, "$rp$", vCode, 1)
	params["case"] = "lower"
	return nil
}

func (this *ForgetPassword1Interceptor) After(tx *sql.Tx, result *[]interface{}, config *wsl.Config) error {
	if len(*result) == 0 {
		return errors.New("Failed get user information")
	}
	if userData, ok := (*result)[0].([]map[string]string); ok {
		if len(userData) == 0 {
			return errors.New("Failed get user information")
		}
		username := userData[0]["username"]
		email := userData[0]["email"]
		status := userData[0]["status"]
		err := sendMail(
			config.Mail.MailHost,
			config.Mail.MailUsername,
			config.Mail.MailPassword,
			config.Mail.MailFrom,
			"subject"+username, status, email)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Failed get user information")
	}
	return nil
}
