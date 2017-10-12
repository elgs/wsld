package main

import (
	"flag"
	"log"

	"github.com/elgs/wsl"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	confFile := flag.String("c", "/etc/myapp/app.json", "Configration file path.")
	flag.Parse()

	wsld, err := wsl.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	// Register a bunch of interceptors
	// wsl.RegisterGlobalInterceptors(&DummyInterceptor{Message: "Global"})
	wsl.RegisterGlobalInterceptors(&AuthInterceptor{})

	wsl.RegisterQueryInterceptors("load-scripts", &LoadScriptsInterceptor{})
	wsl.RegisterQueryInterceptors("login", &LoginInterceptor{})
	wsl.RegisterQueryInterceptors("signup", &SignupInterceptor{})
	wsl.RegisterQueryInterceptors("forget-password-1", &ForgetPassword1Interceptor{})

	wsld.Start()
	wsl.Hook()
}
