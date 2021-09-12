package main

import (
	"fmt"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/app/api"

	"github.com/danish45007/hex-architecture-golang/internal/ports"
)

func main() {
	//ports

	//core(business-logic)
	var core ports.ArithmeticPort
	//application
	var app ports.APIPort
	//framework
	var db ports.DBPort
	app = api.NewAdapter(core, db)
	res, err := app.GetAddition(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
