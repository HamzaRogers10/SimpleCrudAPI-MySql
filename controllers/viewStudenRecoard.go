package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"fypHamza/database"
	"fypHamza/models"
	"net/http"
)

func ViewStudentRec(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	//fmt.Println("welcome to View record")
	db := database.Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
	}

	query := "SELECT * FROM student WHERE email= " + student.Email + "'"
	result, err := db.Query(query)
	if err != nil {
		fmt.Println("This is execution error: ", err)
	}
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {

		}
	}(result)

	var newstudent models.Student
	for result.Next() {
		err := result.Scan(&newstudent.Firstname, &newstudent.Lastname, &newstudent.Email, &newstudent.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	newstudent.Password = "*************"
	if student.Email == newstudent.Email {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(newstudent)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode("User not found")
	}

}
