package main

import (
	"fypHamza/database"
	"fypHamza/routes"
	"net/http"
)

func main() {
	database.Sqlclient()

	routes.SignupRouting()
	routes.ViewRouting()
	routes.UpdateRoute()
	routes.RemoveRouting()
	routes.LoginRoute()

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		return
	}

}
