package models

import (

	"codesnooper.com/api/db"
	hash "codesnooper.com/api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	u.Password, err = hash.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	return nil
}

func Login(email string, password string) (string, error) {
	query := `SELECT * FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, email)
	var user User
	var token string
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return token, err
	}
	err = hash.CheckPasswordHash(password, user.Password)
	if err != nil {
		return token, err
	}
	token, err = hash.GenerateToken(user.Email, user.ID)
	if err != nil {
		return token, err
	}
	return token, nil
}
