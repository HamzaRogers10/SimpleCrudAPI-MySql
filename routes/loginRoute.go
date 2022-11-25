package routes

import (
	"fypHamza/controllers"
	"net/http"
)

func LoginRoute() {
	http.HandleFunc("/login", controllers.Login)

}
