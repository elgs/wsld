package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/elgs/wsl"
	"github.com/elgs/wsld/interceptors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	version := flag.Bool("v", false, "prints version")
	confFile := flag.String("c", "wsld_full.json", "Configration file path.")
	flag.Parse()

	if *version {
		fmt.Println("Alpha 20200412")
		os.Exit(0)
	}

	wsld, err := wsl.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	// Register a bunch of interceptors
	registerInterceptors()

	wsld.Start()
	wsl.Hook()
}

func registerInterceptors() {
	wsl.RegisterGlobalInterceptors(&interceptors.AuthInterceptor{})

	wsl.RegisterQueryInterceptors("load-scripts", &interceptors.LoadScriptsInterceptor{})
	wsl.RegisterQueryInterceptors("login", &interceptors.LoginInterceptor{})
	wsl.RegisterQueryInterceptors("logout", &interceptors.LogoutInterceptor{})
	wsl.RegisterQueryInterceptors("session", &interceptors.SessionInterceptor{})
	wsl.RegisterQueryInterceptors("forget-password-send-code", &interceptors.ForgetPasswordInterceptor{})
	wsl.RegisterQueryInterceptors("forget-password-verify-code", &interceptors.ResetPasswordInterceptor{})
	wsl.RegisterQueryInterceptors("reset-password", &interceptors.ResetPasswordInterceptor{})
	wsl.RegisterQueryInterceptors("change-password", &interceptors.ChangePasswordInterceptor{})
}
