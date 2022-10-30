package helper

import (
	"database/sql"
	
	
	"github.com/ainiokta/task-5-vix-btpns-AINI-/database"
)

type Validation struct {
	conn *sql.DB
}
func NewValidation() *Validation {
	conn, err := database.DBConn()

	if err != nil {
		panic(err)
	}

	return &Validation {
		conn: conn,
	}
}
