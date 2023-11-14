package api

import (
	"database/sql"
	"encoding/json"
	"myproject/model"
	"net/http"
)

type Handler struct {
	biz IBizLogic
}

func NewHandler(db *sql.DB) Handler {
	return Handler{biz: NewBizLogic(db)}
}

func (h Handler) CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}
    
		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := h.biz.CreateBookLogic(book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
