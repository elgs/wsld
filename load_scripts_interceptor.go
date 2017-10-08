package main

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type LoadScriptsInterceptor struct {
	*wsl.DefaultInterceptor
}

var scriptNames []string

func (this *LoadScriptsInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, ii *wsl.InterceptorInterface) error {
	if params["_user_mode"] == "root" {
		sn, err := ii.LoadScripts()
		scriptNames = sn
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No Access")
}

func (this *LoadScriptsInterceptor) After(tx *sql.Tx, result *[]interface{}, ii *wsl.InterceptorInterface) error {
	for _, s := range scriptNames {
		*result = append(*result, s)
	}
	return nil
}
