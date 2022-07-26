package boilerplate

import (
	"log"
	"os"
)

func createRouter() (err error) {
	f, err := os.Create("router/router.go")
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
package router

import (
	"github.com/gin-gonic/gin"
)

// setup gin's router
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"responseCode": "40404", "responseMessage": "Invalid Path"})
	})

	return router
}
`)

	if err != nil {
		return err
	}

	return nil
}
