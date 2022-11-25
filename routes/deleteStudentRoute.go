package routes

import (
	"fypHamza/controllers"
	"net/http"
)

func RemoveRouting() {

	http.HandleFunc("/delete", controllers.DeleteStudnetRec)
}
