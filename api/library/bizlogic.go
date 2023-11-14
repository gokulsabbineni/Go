package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"myproject/dataservice"
	"myproject/model"
	"net/http"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// author and title
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}

	if exists, err := dataservice.GetBook(db, book.Title, book.Author); err != nil {
		return err
	} else if exists {
		http.Error(w, "Book already exists", http.StatusBadRequest)
		return errors.New("book already exists")
	}
	return dataservice.CreateBook(db, w, r)
}

//updatebooklogic
//deletebooklogic

// first api reponse should be verify the routes. routes nothing but checking what is parsed in the url
// handler will check whether we are parsing correct httprquest and if it is good it will send the data to businesslogic layer
// it will perform necessary operations and if everything is good it will call dataservice to perform necessary db operations
// dataservic e will perform necessary db operations

// ticket:

// validation in createbooklogic where you have to check whether the record is already present are not
