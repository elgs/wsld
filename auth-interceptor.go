package main

import (
	"database/sql"
	"fmt"
	"strings"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/elgs/gojq"
	"github.com/elgs/wsl"
)

type AuthInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *AuthInterceptor) Before(
	tx *sql.Tx,
	script *string,
	params map[string]string,
	headers map[string]string,
	config *wsl.Config) error {
	authHeader := headers["Authorization"]
	if authHeader != "" {
		s := strings.Split(authHeader, " ")
		if len(s) == 2 {
			token := s[1]
			payload, _, err := jose.Decode(token, []byte(config.Web.JwtKey))

			if err == nil {
				// fmt.Printf("\npayload = %v\n", payload)
				parser, err := gojq.NewStringQuery(payload)
				if err != nil {
					return err
				}
				_sessionId, err := parser.Query("id")
				if err != nil {
					return err
				}
				_userId, err := parser.Query("user_id")
				if err != nil {
					return err
				}
				_userMode, err := parser.Query("mode")
				if err != nil {
					return err
				}
				sessionId := fmt.Sprintf("%v", _sessionId)
				userId := fmt.Sprintf("%v", _userId)
				userMode := fmt.Sprintf("%v", _userMode)

				params["__session_id"] = sessionId
				params["__user_id"] = userId
				params["__user_mode"] = userMode

				//and/or use headers
				// fmt.Printf("\nheaders = %v\n", headers)
			} else {
				return err
			}
		}
	}
	return nil
}
func (this *AuthInterceptor) After(tx *sql.Tx, result *[]interface{}, config *wsl.Config) error {
	return nil
}
func (this *AuthInterceptor) OnError(err *error) error {
	return *err
}