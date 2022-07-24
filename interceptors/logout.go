package interceptors

import (
	"database/sql"

	"github.com/elgs/wsl"
)

type LogoutInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LogoutInterceptor) After(tx *sql.Tx, context *wsl.Context, exportedResults any, cumulativeResults any) error {
	return context.App.Cache.Delete("session:" + context.SessionID)
}
