package mysql

import (
	"database/sql"
	"errors"
)

type UserBooksModel struct {
	DB *sql.DB
}

func (ub *UserBooksModel) Add(userID, bookID int, rated int, feedback string) error {
	stmt := `INSERT INTO user_books (user_id, book_id, rated, feedback) VALUES (?, ?, ?, ?)`
	_, err := ub.DB.Exec(stmt, userID, bookID, rated, feedback)
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

func (ub *UserBooksModel) GetBooksByUserID(userID int) ([]*Book, error) {
	stmt := `SELECT b.id, b.name, b.author, b.description, b.category, b.image, ub.rated, ub.feedback 
             FROM books b 
             INNER JOIN user_books ub ON b.id = ub.book_id 
             WHERE ub.user_id = ?`

	rows, err := ub.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*Book{}

	for rows.Next() {
		b := &Book{}
		var rated int
		var feedback string

		// Adjust the rows.Scan to include the rated and feedback variables
		err := rows.Scan(&b.ID, &b.Name, &b.Author, &b.Description, &b.Category, &b.Image, &rated, &feedback)
		if err != nil {
			return nil, err
		}
		b.Rated = rated
		b.Feedback = feedback

		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
func (ub *UserBooksModel) GetReviews(bookID int) ([]ReviewResponse, error) {
	query := `SELECT u.name, ub.rated, ub.feedback 
              FROM user_books ub
              JOIN users u ON ub.user_id = u.id
              WHERE ub.book_id = ?`

	rows, err := ub.DB.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []ReviewResponse
	for rows.Next() {
		var r ReviewResponse
		if err := rows.Scan(&r.Username, &r.Rated, &r.Feedback); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}

	return reviews, nil
}

func (ub *UserBooksModel) Update(bookId, userId, rated int, feedback string) error {
	stmt := `UPDATE user_books SET rated = ?, feedback = ? WHERE book_id = ? AND user_id = ?`
	result, err := ub.DB.Exec(stmt, rated, feedback, bookId, userId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	return nil
}
