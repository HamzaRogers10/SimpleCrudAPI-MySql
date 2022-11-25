package routes

import (
	"fypHamza/controllers"
	"net/http"
)

func SignupRouting() {
	http.HandleFunc("/signup", controllers.AddStudentRec)
}
