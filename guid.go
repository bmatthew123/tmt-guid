package main

import (
	"fmt"
	eden "github.com/byu-oit-ssengineering/tmt-eden"
	"github.com/satori/go.uuid"
)

func NewGuid(c *eden.Context) {
	guid := uuid.NewV4()
	c.Respond(200, eden.Response{"OK", guid})
}

func main() {
	r := eden.New()

	r.GET("/guid", NewGuid)

	if err := r.Run(":8000"); err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}
}
