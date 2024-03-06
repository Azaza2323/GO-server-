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

func (u *UserModel) GetByID(userID int) (*User, error) {
	stmt := `SELECT id, name, email, role FROM users WHERE id = ?`
	user := &User{}
	err := u.DB.QueryRow(stmt, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserModel) Update(id int, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `UPDATE users SET name = ?, email = ?, hashed_password = ? WHERE id = ?`
	_, err = u.DB.Exec(stmt, name, email, hashedPassword, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) Delete(id int) error {
	stmt := `DELETE FROM users WHERE id = ?`
	_, err := u.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) Authenticate(email, password string) (int, string, error) {
	var id int
	var hashedPassword []byte
	var role string

	stmt := "SELECT id, hashed_password, role FROM users WHERE email = ?"
	row := u.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword, &role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", ErrInvalidCredentials
		}
		return 0, "", err
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, "", ErrInvalidCredentials
		}
		return 0, "", err
	}
	return id, role, nil
}
func (u *UserModel) GetName(id int) (string, error) {
	stmt := `SELECT name FROM users WHERE id=?`
	row := u.DB.QueryRow(stmt, id)
	var name string
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
