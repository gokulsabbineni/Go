package api

import (
	"database/sql"
	"errors"
	"myproject/dataservice"
	"myproject/model"
)

type IBizLogic interface {
	CreateBookLogic(book model.Book) error
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *Bizlogic {
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) CreateBookLogic(book model.Book) error {
	if exists, err := dataservice.GetBook(bl.DB, book.Title, book.Author); err != nil {
		return err
	} else if exists {
		return errors.New("book already exists")
	}
	return dataservice.CreateBook(bl.DB, book)
}

//updatebooklogic
//deletebooklogic

// first api reponse should be verify the routes. routes nothing but checking what is parsed in the url
// handler will check whether we are parsing correct httprquest and if it is good it will send the data to businesslogic layer
// it will perform necessary operations and if everything is good it will call dataservice to perform necessary db operations
// dataservic e will perform necessary db operations

// ticket:

// validation in createbooklogic where you have to check whether the record is already present are not
