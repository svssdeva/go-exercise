package models

import (
	"deva.com/backend/v2/db"
	"deva.com/backend/v2/utils"
)

type User struct {
	ID        int64  `json:"id"`                          // Unique identifier for the user
	Name      string `json:"name" binding:"required"`     // Name of the user
	Email     string `json:"email" binding:"required"`    // Email address of the user
	Password  string `json:"password" binding:"required"` // Password for the user account
	CreatedAt string `json:"created_at"`                  // Timestamp of when the user was created
}

func (u *User) Save() error {
	query := `
INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Could not prepare statement: " + err.Error())
	}
	defer stmt.Close()

	hashedPassword, _ := utils.HashPassword(u.Password)
	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		panic("Could not execute statement: " + err.Error())
	}
	id, errId := result.LastInsertId()
	u.ID = id
	return errId
}
