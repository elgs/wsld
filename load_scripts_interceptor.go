package main

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type LoadScriptsInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LoadScriptsInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, fns map[string]func()) error {
	if params["_user_mode"] == "root" {
		fns["loadScripts"]()
		return nil
	}
	return errors.New("No Access")
}
