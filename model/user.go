package model

import (
	"database/sql"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func GetUserByNameAndPassword(name, password string, db *sql.DB) (User, error) {
	var user User

	err := db.QueryRow("SELECT id, name, surname, email FROM users WHERE name = ? AND password = ?", name, password).Scan(
		&user.ID, &user.Name, &user.Surname, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertUser(user User, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users (name, surname, email, password) VALUES (?, ?, ?, ?)",
		user.Name, user.Surname, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
