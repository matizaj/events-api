package models

import (
	"errors"
	"events-api/db"
	"events-api/utils"
	"fmt"
)

type User struct {
	Id       int64
	Username string
	Password string `binding:"required"`
	Email    string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Printf("Failed to prepare query %v", err)
		return err
	}
	defer stmt.Close()
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Printf("Failed to hash password %v", err)
		return err
	}
	result, err := stmt.Exec(u.Email, hash)
	if err != nil {
		fmt.Printf("Failed to execute query %v", err)
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to retrive user id %v", err)
		return err
	}
	u.Id = userId
	return nil
}
func (u User) Login() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	err = utils.CheckPassword(u.Password, retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	return nil
}
func updateUserData(u *User) {
	u.Username = "test"
	u.Password = "dupa"
	u.Id = 999
}
