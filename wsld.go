package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/elgs/wsl"
	"github.com/elgs/wsld/interceptors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	confPath := flag.String("c", "wsld.json", "configration file path")
	flag.Parse()
	confBytes, err := ioutil.ReadFile(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	config, err := wsl.NewConfig(confBytes)
	if err != nil {
		log.Fatalln(err)
	}

	app := wsl.NewApp(config)

	// app.RegisterQueryInterceptors("signup", &SignupInterceptor{})
	app.RegisterGlobalInterceptors(&interceptors.AuthInterceptor{})

	app.Start()
}
