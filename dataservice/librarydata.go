package dataservice

import (
	"database/sql"
	"encoding/json"
	"myproject/model"
	"net/http"
)

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}

	query := `INSERT INTO books (title, author, year) VALUES (?,?,?)`
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}

func GetBook(db *sql.DB, title string, author string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM books WHERE title=? and author=?)`
	err := db.QueryRow(query, title, author).Scan(&exists)
	return exists, err
}
