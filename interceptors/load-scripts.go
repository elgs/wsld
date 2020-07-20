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

func loadScripts(app *wsl.WSL) ([]string, error) {
	err := app.LoadScripts("")
	scriptNames := make([]string, len(app.Scripts))

	i := 0
	for k := range app.Scripts {
		scriptNames[i] = k
		i++
	}
	return scriptNames, err
}

func (this *LoadScriptsInterceptor) Before(tx *sql.Tx, context map[string]interface{}) error {
	params := context["params"].(map[string]interface{})
	app := context["app"].(*wsl.WSL)
	if params["__user_mode"] == "root" {
		sn, err := loadScripts(app)
		scriptNames = sn
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No Access")
}

func (this *LoadScriptsInterceptor) BeforeEach(tx *sql.Tx, context map[string]interface{}, script *string, sqlParams []interface{}, scriptIndex int, cumulativeResults interface{}) (bool, error) {
	return false, nil
}

func (this *LoadScriptsInterceptor) AfterEach(tx *sql.Tx, context map[string]interface{}, result interface{}, allResults interface{}, scriptIndex int) error {

	// data, ok := result["data"].([]interface{})
	// if !ok {
	// 	return errors.New("No data is returned.")
	// }

	// for _, s := range scriptNames {
	// 	data = append(data, s)
	// }
	return nil
}
