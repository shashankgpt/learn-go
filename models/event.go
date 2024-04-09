package models

import (
	"fmt"
	"time"

	"codesnooper.com/api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserId      int64       `json:"user_id"`
}

var events []Event = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() []Event {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}
	return events
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT * FROM events where id = ?`
	row := db.DB.QueryRow(query, id)
	fmt.Print(row)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
