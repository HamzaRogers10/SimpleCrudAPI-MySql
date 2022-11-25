package routes

import (
	"fypHamza/controllers"
	"net/http"
)

func UpdateRoute() {

	http.HandleFunc("/update", controllers.UpdateStudentRec)
}
