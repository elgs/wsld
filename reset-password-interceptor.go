package main

import (
	"database/sql"
	"errors"

	"github.com/elgs/gosqljson"
	"github.com/elgs/wsl"
)

type ResetPasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ResetPasswordInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, context map[string]interface{}, wslApp *wsl.WSL) error {

	if context["session_id"] == "" {
		return errors.New("Invalid token.")
	}

	return nil
}

func (this *ResetPasswordInterceptor) After(
	tx *sql.Tx,
	params map[string]string,
	result map[string]interface{},
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	username := params["_0"]
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
