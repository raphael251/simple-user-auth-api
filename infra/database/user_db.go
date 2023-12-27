package database

import (
	"database/sql"

	"github.com/raphael251/simple-user-auth-api/internal/entity"
)

func InsertUser(db *sql.DB, user *entity.User) error {
	stmt, err := db.Prepare("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func FindAllUsers(db *sql.DB) ([]entity.User, error) {
	stmt, err := db.Query("SELECT id, email FROM users")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var users []entity.User

	for stmt.Next() {
		var user entity.User

		stmt.Scan(&user.ID, &user.Email)

		users = append(users, user)
	}

	return users, nil
}

func FindUserByEmail(db *sql.DB, email string) (*entity.User, error) {
	stmt, err := db.Prepare("SELECT id, email, password FROM users WHERE email = $1")

	if err != nil {
		return nil, err
	}

	var user = entity.User{}

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
