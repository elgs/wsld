package main

import (
	"database/sql"
	"encoding/json"
	"log"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/elgs/wsl"
)

type LoginInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LoginInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, config *wsl.Config) error {
	params["case"] = "lower"
	return nil
}

func (this *LoginInterceptor) After(tx *sql.Tx, result *[]interface{}, config *wsl.Config) error {
	if v, ok := (*result)[0].([]map[string]string); ok {
		if len(v) == 0 {
			log.Println("Login failed.")
		} else {
			log.Println("Login succeeded.")
			loginData, err := json.Marshal(v[0])
			if err != nil {
				return nil
			}
			token, err := jose.Sign(string(loginData), jose.HS256, []byte(config.Web.JwtKey))
			if err != nil {
				return nil
			}
			*result = append(*result, token)
		}
	} else {
		log.Println("Login failed.")
		return nil
	}
	return nil
}
