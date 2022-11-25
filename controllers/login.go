package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"fypHamza/database"
	"fypHamza/models"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	db := database.Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
	}

	query := "SELECT * FROM student WHERE email=" + student.Email + "'"
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
		if student.Password == newstudent.Password {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			err := json.NewEncoder(w).Encode(fmt.Sprintf("Successfully Login by: %s", student.Email))
			if err != nil {
				return
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Password incorrect try again with the correct password to login")
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Email incorrect Try again with the correct email to login")
	}
}
