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

func _(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome To Signup")
	var student models.Student
	db := database.Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		fmt.Println("Error decoding response object", err)
	}

	query := "Insert into student (firstname,lastname,email,password) values ( '" + student.Firstname + "' ,'" + student.Lastname + "' ,'" + student.Email + "','" + student.Password + " )"
	querychk := "Select * from student where email = " + student.Email + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		fmt.Println("Result Check Error", err)
	}
	defer func(resultchk *sql.Rows) {
		err := resultchk.Close()
		if err != nil {

		}
	}(resultchk)
	var (
		chkstudent models.Student
	)
	for resultchk.Next() {
		err := resultchk.Scan(&chkstudent.Firstname, &chkstudent.Lastname, &chkstudent.Email, &chkstudent.Password)
		if err != nil {
			fmt.Println("Next Result check Error", err)
		}
	}
	if student.Email != chkstudent.Email {
		result, err := db.Exec(query)
		if err != nil {
			fmt.Println("Check Email inside Error 1", err)
		}
		fmt.Println(result)
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(student)
		if err != nil {
			return
		}
		resp := json.NewEncoder(w).Encode("User Successfully SignUp")
		fmt.Printf("This is the new encoded response: %+v\n", resp)
	} else {
		err := json.NewEncoder(w).Encode("This user already exist")
		if err != nil {
			return
		}
	}
}
