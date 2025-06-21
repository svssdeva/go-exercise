package models

import (
	"time"

	"deva.com/backend/v2/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events []Event = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id)
	 VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Could not prepare statement: " + err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		panic("Could not execute statement: " + err.Error())
	}
	id, errId := result.LastInsertId()
	e.ID = id
	return errId
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
