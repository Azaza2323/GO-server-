package mysql

import (
	"database/sql"
)

type UserBooksModel struct {
	DB *sql.DB
}

func (ub *UserBooksModel) Add(userID, bookID int) error {
	stmt := `INSERT INTO user_books (user_id, book_id) VALUES (?, ?)`
	_, err := ub.DB.Exec(stmt, userID, bookID)
	if err != nil {
		return err
	}
	return nil
}

func (ub *UserBooksModel) Delete(userID, bookID int) error {
	stmt := `DELETE FROM user_books WHERE user_id = ? AND book_id = ?`
	_, err := ub.DB.Exec(stmt, userID, bookID)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserBooksModel) GetBooksByUserID(userID int) ([]*Book, error) {
	stmt := `SELECT b.id, b.name, b.author, b.description, b.category, b.image FROM books b INNER JOIN user_books ub ON b.id = ub.book_id WHERE ub.user_id = ?`

	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*Book{}

	for rows.Next() {
		b := &Book{}
		err := rows.Scan(&b.ID, &b.Name, &b.Author, &b.Description, &b.Category, &b.Image)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
