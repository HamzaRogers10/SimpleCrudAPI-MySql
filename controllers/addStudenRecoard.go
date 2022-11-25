package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"fypHamza/database"
	"fypHamza/models"
	"net/http"
)

func AddStudentRec(w http.ResponseWriter, r *http.Request) {
	var student models.Student
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
	if student.Email != newstudent.Email {
		query := "INSERT into student ( firstname,lastname,email,password ) values ( '" + student.Firstname + "' ,'" + student.Lastname + "','" + student.Email + "','" + student.Password + "')"
		//	fmt.Println(query)
		result, err := db.Exec(query)
		if err != nil {
			fmt.Printf("This is execution error: %v\n", err)
		}
		//	fmt.Println(result)
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(fmt.Sprintf("Successfully SignUp: %s", result))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Email Already Exist")
	}
}
