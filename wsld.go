package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/elgs/wsl"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	version := flag.Bool("v", false, "prints version")
	confFile := flag.String("c", "/etc/myapp/app.json", "Configration file path.")
	flag.Parse()

	if *version {
		fmt.Println("Alpha 20171115")
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
	wsl.RegisterGlobalInterceptors(&AuthInterceptor{})

	wsl.RegisterQueryInterceptors("load-scripts", &LoadScriptsInterceptor{})
	wsl.RegisterQueryInterceptors("login", &LoginInterceptor{})
	wsl.RegisterQueryInterceptors("signup", &SignupInterceptor{})
	wsl.RegisterQueryInterceptors("forget-password-1", &ForgetPassword1Interceptor{})
	wsl.RegisterQueryInterceptors("forget-password-2", &ResetPasswordInterceptor{})
	wsl.RegisterQueryInterceptors("reset-password", &ResetPasswordInterceptor{})
	wsl.RegisterQueryInterceptors("change-password", &ChangePasswordInterceptor{})
}
