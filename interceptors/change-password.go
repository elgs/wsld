package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/gosqljson"
	"github.com/elgs/wsl"
)

type ChangePasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ChangePasswordInterceptor) BeforeEach(tx *sql.Tx, script *string, sqlParams []interface{},
	context map[string]interface{}, index int,
	wslApp *wsl.WSL) (bool, error) {

	if context["session_id"] == "" {
		return false, errors.New("Invalid token.")
	}

	return false, nil
}

func (this *ChangePasswordInterceptor) AfterEach(tx *sql.Tx, params map[string]string, result interface{},
	context map[string]interface{}, index int,
	wslApp *wsl.WSL) error {

	sessionId := params["__session_id"]
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
