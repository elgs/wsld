package main

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/elgs/gosqljson"
	"github.com/elgs/gostrgen"
	"github.com/elgs/wsl"
)

type ResetPasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ResetPasswordInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, context map[string]interface{}, wslApp *wsl.WSL) error {

	sessionKey, err := gostrgen.RandGen(20, gostrgen.All, "", "")
	if err != nil {
		return err
	}
	*script = strings.Replace(*script, "$session-key$", sessionKey, 1)
	username := params["_0"]
	context["username"] = username
	return nil
}

func (this *ResetPasswordInterceptor) After(tx *sql.Tx, result map[string]interface{},
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	username := context["username"]
	userData, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT ID FROM USER WHERE USERNAME=? OR EMAIL=?", username, username)
	if err != nil {
		return err
	}
	if len(userData) != 1 {
		return errors.New("Failed to find user.")
	}
	userId := userData[0]["id"]
	delete(userKeys, userId)
	return nil
}
