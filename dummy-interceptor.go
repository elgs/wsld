package main

import (
	"database/sql"
	"log"

	"github.com/elgs/wsl"
)

type DummyInterceptor struct {
	*wsl.DefaultInterceptor
	Message string
}

func (this *DummyInterceptor) Before(tx *sql.Tx, script *string, params map[string]string, headers map[string]string, config *wsl.Config) error {
	log.Println("Before", this.Message)
	// log.Println(headers)
	// log.Println(params["__client_ip"])
	return nil
}
func (this *DummyInterceptor) After(tx *sql.Tx, result *[]interface{}, config *wsl.Config) error {
	log.Println("After", this.Message)
	return nil
}
func (this *DummyInterceptor) OnError(err *error) error {
	log.Println("error: ", this.Message, *err)
	return nil
}