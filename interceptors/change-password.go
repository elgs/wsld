package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type ChangePasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ChangePasswordInterceptor) Before(tx *sql.Tx, context map[string]interface{}) error {

	if context["session_id"] == "" {
		return errors.New("invalid_token")
	}

	return nil
}

func (this *ChangePasswordInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int, cumulativeResults []interface{}) (bool, error) {
	if scriptIndex == 2 {
		if cumulativeResults[1] == int64(0) {
			// if password is not changed, skip deleting other sessions
			return true, nil
		}
	}
	return false, nil
}
