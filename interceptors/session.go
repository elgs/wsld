package interceptors

import (
	"database/sql"

	"github.com/elgs/wsl"
)

type SessionInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *SessionInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, scriptIndex int) error {

	// data, ok := result["data"].([]interface{})
	// if !ok {
	// 	return errors.New("no_data")
	// }

	// if session, ok := data[0].([]map[string]string); ok && len(session) > 0 {
	// 	result["session_id"] = session[0]["session_id"]
	// 	delete(result, "data")
	// 	return nil
	// }

	// 	vCode := u[0]["v_code"]
	// 	if vCode != "" {
	// 		fmt.Printf("User not verified (%v)", u[0]["username"])

	// 		email := u[0]["email"]
	// 		err := wslApp.SendMail(wslApp.Config.App["mail_from"].(string), "New Account Verification Code", vCode, email)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		result["data"] = "user_not_verified"

	// 		return nil
	// 	}

	// 	log.Printf("Login succeeded (%v)", u[0]["username"])

	// 	mapClaims := map[string]interface{}{
	// 		"user_id":  u[0]["user_id"],
	// 		"id":       u[0]["id"],
	// 		"mode":     u[0]["mode"],
	// 		"username": u[0]["username"],
	// 		"email":    u[0]["email"],
	// 	}

	// 	userId := u[0]["user_id"]
	// 	sessionKey, err := getSessionKey(tx, userId)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	expMinutes := wslApp.Config.App["session_expire_in_minutes"].(float64)
	// 	token, err := jwt.Encode(mapClaims, int(expMinutes)*60, sessionKey)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	result["token"] = token
	// 	delete(result, "data")
	// } else {
	// 	log.Println("Login failed.")
	// 	return errors.New("login_failed")
	// }
	return nil
}
