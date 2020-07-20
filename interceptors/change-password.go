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

func (this *ChangePasswordInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int, cumulativeResults interface{}) (bool, error) {

	if context["session_id"] == "" {
		return false, errors.New("Invalid token.")
	}

	return false, nil
}

func (this *ChangePasswordInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, allResults interface{}, scriptIndex int) error {
	params := context["params"].(map[string]interface{})

	sessionId := params["__session_id"]
	userData, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT USER_ID FROM USER_SESSION WHERE ID=?", sessionId)
	if err != nil {
		return err
	}
	if len(userData) != 1 {
		return errors.New("Failed to find user.")
	}
	// userId := userData[0]["user_id"]
	// delete(userKeys, userId)
	return nil
}
