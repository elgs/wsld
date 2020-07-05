package interceptors

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/elgs/gosqljson"
	"github.com/elgs/jwt"
	"github.com/elgs/wsl"
)

var userKeys = make(map[string]string)
var userSessionIds = make(map[string]bool)

type AuthInterceptor struct {
	*wsl.DefaultInterceptor
}

func getSessionKey(tx *sql.Tx, userId string) (string, error) {
	if val, ok := userKeys[userId]; ok {
		return val, nil
	}
	dbResult, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT PASSWORD FROM USER WHERE ID=?", userId)
	if err != nil {
		return "", err
	}
	if len(dbResult) != 1 {
		return "", errors.New("User not found.")
	}
	sessionKey := dbResult[0]["password"]
	if sessionKey == "" {
		return "", errors.New("Session key is empty.")
	}
	userKeys[userId] = sessionKey
	return sessionKey, nil
}

func getSessionId(tx *sql.Tx, sessionId string) (bool, error) {
	if val, ok := userSessionIds[sessionId]; ok {
		return val, nil
	}
	dbResult, err := gosqljson.QueryTxToMap(tx, "lower", "SELECT ID FROM USER_SESSION WHERE ID=? AND NOT JSON_CONTAINS_PATH(STATUS, 'one', '$.pfv');", sessionId)
	if err != nil {
		return false, err
	}
	if len(dbResult) != 1 {
		return false, errors.New("Session not found.")
	}
	userSessionIds[sessionId] = true
	return true, nil
}

func (this *AuthInterceptor) Before(tx *sql.Tx, context map[string]interface{}) error {

	params := context["params"].(map[string]interface{})

	if tokenString, ok := context["access_token"].(string); ok {
		claims, err := jwt.Decode(tokenString)
		userId := claims["user_id"]
		sessionId := claims["id"].(string)
		hasSessionId, err := getSessionId(tx, sessionId)
		if !hasSessionId || err != nil {
			return errors.New("Invalid token.")
		}

		sessionKey, err := getSessionKey(tx, userId.(string))
		verified, err := jwt.Verify(tokenString, sessionKey)
		if !verified || err != nil {
			return errors.New("Invalid key.")
		}

		params["__session_id"] = fmt.Sprintf("%v", sessionId)
		params["__user_id"] = fmt.Sprintf("%v", userId)
		params["__user_mode"] = fmt.Sprintf("%v", claims["mode"])

		// signify token is verified, token interceptors could check context["session_id"] for whether token is verified, or ignore for public apis.
		context["session_id"] = sessionId
	}
	return nil
}

func (this *AuthInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int) (bool, error) {
	return false, nil
}
func (this *AuthInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, scriptIndex int) error {
	return nil
}
func (this *AuthInterceptor) OnError(err *error) error {
	return *err
}
