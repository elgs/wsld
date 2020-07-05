package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type LogoutInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LogoutInterceptor) BeforeEach(tx *sql.Tx, script *string, sqlParams []interface{},
	context map[string]interface{}, index int,
	wslApp *wsl.WSL) (bool, error) {

	if context["session_id"] == "" {
		return false, errors.New("Invalid token.")
	}

	return false, nil
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
