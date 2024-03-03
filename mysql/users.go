package mysql

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, role) VALUES (?, ?, ?, 'user')`
	_, err = u.DB.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM users WHERE email = ?"
	row := u.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		}
		return 0, err
	}

	return id, nil
}

func (u *UserModel) GetRole(id int) (string, error) {
	stmt := `SELECT role FROM users WHERE id=?`
	row := u.DB.QueryRow(stmt, id)
	var role string
	err := row.Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}
