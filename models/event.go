package models

import (
	"events-api/db"
	"fmt"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events []Event

func (e Event) Save() error {
	fmt.Println("[DEBUG LOG]-> ", e)
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?,?,?,?,?);`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		fmt.Printf("Failed to prepare query %v", err)
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		fmt.Printf("Failed to execute query %v", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to get id %v", err)
		return err
	}
	e.Id = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	fmt.Printf("[DEBUG ROWS]-> %v", *rows)
	if err != nil {
		fmt.Printf("Failed to get events %v", err)
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			fmt.Printf("Failed to process events rows %v", err)
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		fmt.Printf("Failed to process events rows %v", err)
		return nil, err
	}
	return &e, nil
}
func (e Event) Update() error {
	query := `
		UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id =?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Printf("Failed to prepare query %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	if err != nil {
		fmt.Printf("Failed to execute query %v", err)
		return err
	}
	return nil
}

func (e Event) DeleteEvent() error {
	query := "DELETE FROM events Where id = ?"
	result, err := db.DB.Exec(query, e.Id)
	if err != nil {
		fmt.Printf("Failed to execute delete query %v", err)
		return err
	}
	rows, err := result.RowsAffected()
	fmt.Printf("[rows] %v", rows)
	return nil
}
