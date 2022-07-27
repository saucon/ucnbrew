package boilerplate

import (
	"log"
	"os"
)

func createMain(pkgname string, appname string) (err error) {
	f, err := os.Create("cmd/" + appname + "/main.go")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	_, err = f.WriteString(`// This file is generated using ucnbrew tool. 
// Check out for more info "https://github.com/saucon/ucnbrew"
package main

import (
	"` + pkgname + `/configs/env"
	"` + pkgname + `/router"

	"log"
)


func main() {
	env.NewEnv(".env")
	cfg := env.Config

	router := router.NewRouter()
	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatal("Error running router : ",err)
	}
}`)

	if err != nil {
		return err
	}

	return nil
}
