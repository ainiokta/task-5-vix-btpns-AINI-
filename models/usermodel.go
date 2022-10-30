package models

import (
	"database/sql"

	"github.com/ainiokta/task-5-vix-btpns-AINI-/app"
	"github.com/ainiokta/task-5-vix-btpns-AINI-/database"
)

type Usermodel struct {
	db *sql.DB
}

func NewUsermodel() *Usermodel {
	conn, err := database.DBConn()

	if err != nil {
		panic(err)
	}

	return &Usermodel{
		db: conn,
	}
}

func (u Usermodel) Where(users *app.Users, fieldName, fieldValue string) error {

	row, err := u.db.Query("SELECT * FROM `users` where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&users.Id, &users.Nama, &users.Email, &users.Username, &users.Password)
	}

	return nil
}
