package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/elgs/wsl"
	"github.com/elgs/wsld/interceptors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr"
)

func main() {
	version := flag.Bool("v", false, "prints version")
	confFile := flag.String("c", "wsld_full.json", "Configration file path.")
	flag.Parse()

	if *version {
		fmt.Println("Alpha 20200719")
		os.Exit(0)
	}

	wsld, err := wsl.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	loadBuiltInScripts(wsld)
	registerInterceptors(wsld)

	wsld.Start()
	wsl.Hook()
}

func registerInterceptors(wsld *wsl.WSL) {
	wsld.RegisterGlobalInterceptors(&interceptors.AuthInterceptor{})

	wsld.RegisterQueryInterceptors("signup", &interceptors.SignupInterceptor{})
	wsld.RegisterQueryInterceptors("login", &interceptors.LoginInterceptor{})
	wsld.RegisterQueryInterceptors("verify-user", &interceptors.VerifyUserInterceptor{})
	wsld.RegisterQueryInterceptors("logout", &interceptors.LogoutInterceptor{})
	wsld.RegisterQueryInterceptors("session", &interceptors.SessionInterceptor{})
	wsld.RegisterQueryInterceptors("forget-password-send-code", &interceptors.ForgetPasswordSendCodeInterceptor{})
	wsld.RegisterQueryInterceptors("reset-password", &interceptors.ResetPasswordInterceptor{})
	wsld.RegisterQueryInterceptors("change-password", &interceptors.ChangePasswordInterceptor{})
}

var scriptNames = []string{
	"init",
	"signup",
	"login",
	"verify-user",
	"session",
	"logout",
	"change-password",
	"reset-password",
	"forget-password-send-code",
	"forget-password-reset-password",
}

func loadBuiltInScripts(wsld *wsl.WSL) {
	scriptsBox := packr.NewBox("./scripts")
	for _, scriptName := range scriptNames {
		if scriptString, err := scriptsBox.FindString(fmt.Sprint(scriptName, ".sql")); err == nil {
			wsld.Scripts[scriptName] = scriptString
		}
	}
}
