package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type LogoutInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LogoutInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int) (bool, error) {

	if context["session_id"] == "" {
		return false, errors.New("Invalid token.")
	}

	return false, nil
}

func (this *LogoutInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, scriptIndex int) error {
	params := context["params"].(map[string]interface{})
	userId := params["user_id"].(string)
	delete(userKeys, userId)
	delete(userSessionIds, context["session_id"].(string))
	return nil
}
