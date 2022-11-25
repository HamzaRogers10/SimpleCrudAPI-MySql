package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"fypHamza/database"
	"fypHamza/models"
	"net/http"
)

func DeleteStudnetRec(w http.ResponseWriter, r *http.Request) {
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
		query := "DELETE from student WHERE email='" + student.Email + "'"
		fmt.Println(query)
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("This is execution error: %v", err)
		}
		//fmt.Println(result)
		output := fmt.Sprintf("%s user deleted", student.Email)
		defer db.Close()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(output)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		err := json.NewEncoder(w).Encode("Email is not exist, enter valid email")
		if err != nil {
			return
		}
	}
}
