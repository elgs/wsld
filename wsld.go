package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/elgs/wsl"
	"github.com/elgs/wsl/interceptors"
	"github.com/elgs/wsl/scripts"
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

	// optionally load built in user management interceptors and scripts, and jobs
	scripts.LoadBuiltInScripts(app)
	interceptors.RegisterBuiltInInterceptors(app)

	app.Start()
}
