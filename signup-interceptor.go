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

func (this *SignupInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, config *wsl.Config) error {
	vCode, err := gostrgen.RandGen(8, gostrgen.LowerUpperDigit, "", "lO") // exclude small L and big O
	if err != nil {
		return err
	}
	*script = strings.Replace(*script, "$pfv$", vCode, 1)
	params["case"] = "lower"
	return nil
}

func (this *SignupInterceptor) After(tx *sql.Tx, result *[]interface{}, config *wsl.Config) error {
	if len(*result) == 0 {
		return errors.New("Failed to sign up.")
	}
	if userData, ok := (*result)[0].([]map[string]string); ok {
		if len(userData) == 0 {
			return errors.New("Failed to sign up.")
		}
		email := userData[0]["email"]
		vCode := userData[0]["v_code"]
		err := sendMail(
			config.Mail.MailHost,
			config.Mail.MailUsername,
			config.Mail.MailPassword,
			config.Mail.MailFrom,
			"New Account Verification Code", vCode, email)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Failed to sign up.")
	}
	return nil
}
