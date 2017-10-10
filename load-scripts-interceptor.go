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

func loadScripts(config *wsl.Config) ([]string, error) {
	err := config.LoadScripts()
	scriptNames := make([]string, len(config.Db.Scripts))

	i := 0
	for k := range config.Db.Scripts {
		scriptNames[i] = k
		i++
	}
	return scriptNames, err
}

func (this *LoadScriptsInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, config *wsl.Config) error {
	if params["_user_mode"] == "root" {
		sn, err := loadScripts(config)
		scriptNames = sn
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No Access")
}

func (this *LoadScriptsInterceptor) After(tx *sql.Tx, result *[]interface{}, config *wsl.Config) error {
	for _, s := range scriptNames {
		*result = append(*result, s)
	}
	return nil
}