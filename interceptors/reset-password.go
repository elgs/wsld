package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/gosqljson"
	"github.com/elgs/wsl"
)

type ResetPasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ResetPasswordInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int) (bool, error) {

	if context["session_id"] == "" {
		return false, errors.New("Invalid token.")
	}

	return false, nil
}

func (this *ResetPasswordInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, scriptIndex int) error {
	params := context["params"].(map[string]interface{})

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
