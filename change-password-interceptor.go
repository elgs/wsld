package main

import (
	"database/sql"
	"errors"

	"github.com/elgs/gosqljson"
	"github.com/elgs/wsl"
)

type ChangePasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ChangePasswordInterceptor) Before(tx *sql.Tx, script *string, params map[string]string,
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	sessionId := params["__session_id"]
	context["session_id"] = sessionId
	return nil
}

func (this *ChangePasswordInterceptor) After(tx *sql.Tx, result map[string]interface{},
	context map[string]interface{},
	wslApp *wsl.WSL) error {

	sessionId := context["session_id"]
	userData, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT USER_ID FROM USER_SESSION WHERE ID=?", sessionId)
	if err != nil {
		return err
	}
	if len(userData) != 1 {
		return errors.New("Failed to find user.")
	}
	userId := userData[0]["user_id"]
	delete(userKeys, userId)
	return nil
}
