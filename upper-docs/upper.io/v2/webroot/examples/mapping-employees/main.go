package main

import (
	"log"

	"upper.io/db.v2"            // Imports the main db package.
	"upper.io/db.v2/postgresql" // Imports the postgresql adapter.
)

// Employee represents an employee.
type Employee struct {
	ID       int    `db:"id,omitempty"`
	LastName string `db:"last_name"`
}

var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	req := sess.Collection("employees").Find().OrderBy("last_name")

	log.Println("A list of employees:")

	for {
		var employee Employee
		err := req.Next(&employee)
		if err != nil {
			if err == db.ErrNoMoreRows {
				break
			}
			log.Fatal(err)
		}
		log.Printf("#%d: %s\n", employee.ID, employee.LastName)
	}
}
