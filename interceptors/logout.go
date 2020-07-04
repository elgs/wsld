package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type LogoutInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LogoutInterceptor) BeforeEach(tx *sql.Tx, script *string, params map[string]string,
	context map[string]interface{}, index int,
	wslApp *wsl.WSL) error {

	if context["session_id"] == "" {
		return errors.New("Invalid token.")
	}

	return nil
}

func (this *LogoutInterceptor) AfterEach(
	tx *sql.Tx,
	params map[string]string,
	result interface{},
	context map[string]interface{},
	index int,
	wslApp *wsl.WSL) error {

	userId := params["user_id"]
	delete(userKeys, userId)
	delete(userSessionIds, context["session_id"].(string))
	return nil
}
