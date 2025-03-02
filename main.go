package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	Duration string `json:"duration"`
	Result   string `json:"result"`
}

func slowQueryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		for i := 1; i <= 10; i++ {
			var result string
			err := db.QueryRow(fmt.Sprintf("SELECT SLEEP(%d)", i)).Scan(&result)
			if err != nil {
				http.Error(w, fmt.Sprintln(err), http.StatusInternalServerError)
				return
			}
		}

		duration := time.Since(start)

		response := Response{
			Duration: duration.String(),
			Result:   "Completed",
		}

		// Відправляємо JSON-відповідь
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/test", slowQueryHandler(db))

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
