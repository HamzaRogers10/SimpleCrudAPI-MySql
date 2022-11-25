package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"fypHamza/database"
	"fypHamza/models"
	"net/http"
)

func UpdateStudentRec(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	//fmt.Println("Welcome to Update Student Rec")
	db := database.Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
	}

	query := "SELECT * FROM student WHERE email='" + student.Email + "'"
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
	if student.Email == newstudent.Email {
		query := "UPDATE lost_listing.student SET firstname=" + student.Firstname + "', lastname='" + student.Lastname + "', password='" + student.Password + "' WHERE email= '" + student.Email + "'"
		//fmt.Println(query)
		result, err := db.Query(query)
		if err != nil {
			fmt.Println("This is execution error: ", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			err := json.NewEncoder(w).Encode("This is execution error:")
			if err != nil {
				return
			}
			json.NewEncoder(w).Encode(err)
		} else {
			fmt.Println(result)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode("Successfully Update student record")
			json.NewEncoder(w).Encode(student)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode("Email is not exist, enter valid email")
	}
}
