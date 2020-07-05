package interceptors

import (
	"database/sql"
	"errors"

	"github.com/elgs/wsl"
)

type LoadScriptsInterceptor struct {
	*wsl.DefaultInterceptor
}

var scriptNames []string

func loadScripts(config *wsl.Config) ([]string, error) {
	err := config.LoadScripts("")
	scriptNames := make([]string, len(config.App["scripts"].(map[string]interface{})))

	i := 0
	for k := range config.App["scripts"].(map[string]interface{}) {
		scriptNames[i] = k
		i++
	}
	return scriptNames, err
}

func (this *LoadScriptsInterceptor) Before(tx *sql.Tx, context map[string]interface{}) error {
	params := context["params"].(map[string]interface{})
	app := context["app"].(*wsl.WSL)
	if params["__user_mode"] == "root" {
		sn, err := loadScripts(app.Config)
		scriptNames = sn
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No Access")
}

func (this *LoadScriptsInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int) (bool, error) {
	return false, nil
}

func (this *LoadScriptsInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, scriptIndex int) error {

	// data, ok := result["data"].([]interface{})
	// if !ok {
	// 	return errors.New("No data is returned.")
	// }

	// for _, s := range scriptNames {
	// 	data = append(data, s)
	// }
	return nil
}
