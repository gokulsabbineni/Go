package dataservice

import (
	"database/sql"
	"myproject/model"
)

func CreateBook(db *sql.DB, book model.Book) error {
	query := `INSERT INTO books (title, author, year) VALUES (?,?,?)`
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(db *sql.DB, title string, author string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM books WHERE title=? and author=?)`
	err := db.QueryRow(query, title, author).Scan(&exists)
	return exists, err
}
