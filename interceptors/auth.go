package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/gosqljson"
	"github.com/elgs/wsl"
)

var sessionQuery = `
SELECT
USER.ID AS USER_ID,
USER.USERNAME,
USER.EMAIL,
USER.MODE,
USER.CREATED_TIME,
USER_SESSION.ID AS SESSION_ID,
USER_SESSION.LOGIN_TIME,
USER_SESSION.LOGIN_IP
FROM USER INNER JOIN USER_SESSION ON USER.ID=USER_SESSION.USER_ID 
WHERE USER_SESSION.ID=?
`

var flagsQuery = `
SELECT CODE, VALUE, PRIVATE FROM USER_FLAG WHERE USER_ID=?
`

var updateLastSeenQuery = `
UPDATE USER_SESSION 
SET LAST_SEEN_TIME=CONVERT_TZ(NOW(),'System','+0:0'),
LAST_SEEN_IP=?
WHERE ID=?
`

var Sessions = make(map[string]map[string]any)

func (this *AuthInterceptor) getSession(tx *sql.Tx, sessionId string) (map[string]any, error) {
	if val, ok := Sessions[sessionId]; ok {
		return val, nil
	}

	dbResult, err := gosqljson.QueryToMap(tx, gosqljson.Lower, sessionQuery, sessionId)
	if err != nil {
		return nil, err
	}
	if len(dbResult) != 1 {
		return nil, errors.New("session_not_found")
	}
	Sessions[sessionId], err = wsl.ConvertMap[string, any](dbResult[0])
	if err != nil {
		return nil, err
	}

	userId := dbResult[0]["user_id"]
	userFlags, err := this.getUserFlags(tx, userId)
	if err != nil {
		return nil, err
	}

	userFlagsMap := make(map[string]any)
	for _, flag := range userFlags {
		private := flag["private"]
		if private == "1" {
			userFlagsMap[flag["code"]] = ""
		} else {
			userFlagsMap[flag["code"]] = flag["value"]
		}
	}

	Sessions[sessionId]["flags"] = userFlagsMap

	return Sessions[sessionId], nil
}

func (this *AuthInterceptor) getUserFlags(tx *sql.Tx, userId string) ([]map[string]string, error) {
	return gosqljson.QueryToMap(tx, gosqljson.Lower, flagsQuery, userId)
}

func (this *AuthInterceptor) updateLastSeen(db *sql.DB, sessionId string, ip string) {
	gosqljson.Exec(db, updateLastSeenQuery, ip, sessionId)
}

type AuthInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *AuthInterceptor) Before(tx *sql.Tx, context *wsl.Context) error {

	if context.AccessToken != "" {

		session, err := this.getSession(tx, context.AccessToken)
		if err != nil {
			return err
		}
		db := context.App.GetDB("main")
		go this.updateLastSeen(db, context.AccessToken, context.ClientIP)

		_ = session
		// params["__session_id"] = fmt.Sprintf("%v", session["session_id"])
		// params["__user_id"] = fmt.Sprintf("%v", session["user_id"])
		// params["__user_mode"] = fmt.Sprintf("%v", session["mode"])

		// context["session_id"] = session["session_id"]
		// context["session"] = session
		// context["user_id"] = session["user_id"]
		// context["user_mode"] = session["mode"]
	}
	return nil
}
