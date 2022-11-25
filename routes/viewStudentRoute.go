package routes

import (
	"fypHamza/controllers"
	"net/http"
)

func ViewRouting() {

	http.HandleFunc("/view", controllers.ViewStudentRec)
}
