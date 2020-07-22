package main

import (
	"flag"
	"log"

	"github.com/elgs/wsl"
	"github.com/elgs/wsl/interceptors"
	"github.com/elgs/wsl/scripts"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	confFile := flag.String("c", "/etc/wsld.json", "configration file path")
	flag.Parse()

	wsld, err := wsl.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	// optionally load built in user management interceptors and scripts
	scripts.LoadBuiltInScripts(wsld)
	interceptors.RegisterBuiltInInterceptors(wsld)

	// done manully
	// wsld.RegisterGlobalInterceptors(&interceptors.AuthInterceptor{})
	// wsld.RegisterQueryInterceptors("signup", &interceptors.SignupInterceptor{})
	// ...

	// wsld.Scripts["init"] = scripts.Init
	// wsld.Scripts["signup"] = scripts.Signup
	// ...

	wsld.Start()
	wsl.Hook()
}
