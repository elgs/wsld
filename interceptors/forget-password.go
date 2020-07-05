package interceptors

import (
	"database/sql"
	"strings"

	"github.com/elgs/gostrgen"
	"github.com/elgs/wsl"
)

type ForgetPasswordInterceptor struct {
	*wsl.DefaultInterceptor
}

func (this *ForgetPasswordInterceptor) BeforeEach(tx *sql.Tx, script *string, sqlParams []interface{},
	context map[string]interface{}, index int,
	wslApp *wsl.WSL) (bool, error) {
	vCode, err := gostrgen.RandGen(8, gostrgen.LowerUpperDigit, "", "lO") // exclude small L and big O
	if err != nil {
		return false, err
	}
	*script = strings.Replace(*script, "$recovering-password$", vCode, 1)
	return false, nil
}

func (this *ForgetPasswordInterceptor) AfterEach(tx *sql.Tx, params map[string]string, result interface{},
	context map[string]interface{}, index int,
	wslApp *wsl.WSL) error {

	// if userData, ok := data[0].([]map[string]string); ok {
	// 	if len(userData) == 0 {
	// 		return errors.New("Failed get user information")
	// 	}
	// 	email := userData[0]["email"]
	// 	vCode := userData[0]["v_code"]
	// 	err := wslApp.SendMail(
	// 		wslApp.Config.App["mail_from"].(string), "Password Reset Verification Code", vCode, email)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	delete(userData[0], "v_code")
	// } else {
	// 	return errors.New("Failed get user information")
	// }
	return nil
}
