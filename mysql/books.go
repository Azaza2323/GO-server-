package mysql

import (
	"database/sql"
)

type BookModel struct {
	DB *sql.DB
}

func (b *BookModel) Get() ([]*Book, error) {
	stmt := `SELECT * FROM books LIMIT 10`
	result, err := b.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	var bookList []*Book
	for result.Next() {
		n := &Book{}
		err := result.Scan(&n.ID, &n.Name, &n.Author, &n.Description, &n.Category, &n.Image)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, n)
	}
	return bookList, nil
}
func (b *BookModel) GetID(id int) ([]*Book, error) {
	stmt := `SELECT * FROM books WHERE id = ?`
	result, _ := b.DB.Query(stmt, id)
	var bookList []*Book
	for result.Next() {
		n := &Book{}
		err := result.Scan(&n.ID, &n.Name, &n.Author, &n.Description, &n.Category, &n.Image)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, n)
	}
	return bookList, nil
}

func (b *BookModel) Insert(name, author, description, category, image string) error {
	stmt := `INSERT INTO books (name, author, description, category, image) VALUES (?, ?, ?, ?, ?)`
	_, err := b.DB.Exec(stmt, name, author, description, category, image)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookModel) Delete(id int) error {
	stmt := `DELETE FROM books WHERE id=?`
	_, err := b.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookModel) GetByCategory(category string) ([]*Book, error) {
	stmt := `SELECT * FROM books WHERE category = ?`
	rows, err := b.DB.Query(stmt, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		book := &Book{}
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.Category, &book.Image)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
