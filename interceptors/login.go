package interceptors

import (
	"database/sql"

	"github.com/elgs/wsl"
)

type LoginInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *LoginInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, context map[string]interface{}, wslApp *wsl.WSL) error {

	// fmt.Println("Before")
	return nil
}

func (this *LoginInterceptor) After(tx *sql.Tx, params map[string]string, result interface{}, context map[string]interface{}, wslApp *wsl.WSL) error {

	// fmt.Println("After")
	return nil
}

func (this *LoginInterceptor) BeforeEach(tx *sql.Tx, script *string, sqlParams []interface{}, context map[string]interface{}, index int, wslApp *wsl.WSL) (bool, error) {

	// fmt.Println("BeforeEach")
	// fmt.Println(*script)
	// fmt.Println(sqlParams)
	// fmt.Println(context)
	// fmt.Println(index)
	return false, nil
}

func (this *LoginInterceptor) AfterEach(tx *sql.Tx, params map[string]string, result interface{}, context map[string]interface{}, index int, wslApp *wsl.WSL) error {

	// fmt.Println("AfterEach")
	// fmt.Println(params)
	// fmt.Println(result)
	// fmt.Println(context)
	// fmt.Println(index)
	// fmt.Println("====================================")

	// if session, ok := data[1].([]map[string]string); ok && len(session) > 0 {
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
