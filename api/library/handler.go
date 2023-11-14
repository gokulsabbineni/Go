package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {
	fmt.Println("Checking whether we are sending correct hhtp request and if it is good proceed with bizlogic")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := CreateBookLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
