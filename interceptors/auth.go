package interceptors

import (
	"database/sql"
	"errors"
	"fmt"

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
SELECT CODE, VALUE FROM USER_FLAG WHERE USER_ID=?
`

var updateLastSeenQuery = `
UPDATE USER_SESSION 
SET LAST_SEEN_TIME=CONVERT_TZ(NOW(),'System','+0:0'),
LAST_SEEN_IP=?
WHERE ID=?
`

func (this *AuthInterceptor) getSession(tx *sql.Tx, context *wsl.Context) (map[string]string, error) {
	if session, err := context.App.Cache.GetMap("session:" + context.SessionID); err != nil && len(session) > 0 {
		context.Session = session
		return session, nil
	}

	dbResult, err := gosqljson.QueryToMap(tx, gosqljson.Lower, sessionQuery, context.SessionID)
	if err != nil {
		return nil, err
	}
	if len(dbResult) != 1 {
		return nil, errors.New("Session not found.")
	}
	session := dbResult[0]
	err = context.App.Cache.SetMap("session:"+context.SessionID, session, 0)
	if err != nil {
		return nil, err
	}

	userId := session["user_id"]
	userFlags, err := this.getUserFlags(tx, userId)
	if err != nil {
		return nil, err
	}
	context.Flags = userFlags

	return session, nil
}

func (this *AuthInterceptor) getUserFlags(tx *sql.Tx, userId string) (map[string]string, error) {
	flagMpas, err := gosqljson.QueryToMap(tx, gosqljson.Lower, flagsQuery, userId)
	if err != nil || len(flagMpas) != 1 {
		return nil, err
	}
	return flagMpas[0], nil
}

func (this *AuthInterceptor) updateLastSeen(db *sql.DB, sessionId string, ip string) {
	gosqljson.Exec(db, updateLastSeenQuery, ip, sessionId)
}

type AuthInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *AuthInterceptor) Before(tx *sql.Tx, context *wsl.Context) error {
	if context.AccessToken == "" {
		context.AuthLevel = wsl.AuthNoToken
		return nil
	}

	session, err := this.getSession(tx, context)
	if err != nil {
		context.AuthLevel = wsl.AuthFailed
		return nil
	}

	db := context.App.GetDB("main")
	go this.updateLastSeen(db, context.AccessToken, context.ClientIP)

	context.Params["__session_id"] = fmt.Sprintf("%v", session["session_id"])
	context.Params["__user_id"] = fmt.Sprintf("%v", session["user_id"])

	if session["mode"] == "root" {
		context.AuthLevel = wsl.AuthRootAuthorized
	} else {
		context.AuthLevel = wsl.AuthUserAuthorized
	}

	return nil
}
