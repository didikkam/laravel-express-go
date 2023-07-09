package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	EmailVerifiedAt sql.NullString `json:"email_verified_at"`
	Password        string         `json:"password"`
	RememberToken   string         `json:"remember_token"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}

func main() {
	// Konfigurasi koneksi database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/laravel_node_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat handler untuk endpoint "/users"
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// Mengeksekusi query untuk mengambil data pengguna dari database
		rows, err := db.Query("SELECT id, name, email, email_verified_at, password, remember_token, created_at, updated_at FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.EmailVerifiedAt, &user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengubah data pengguna menjadi format JSON dan mengirimkannya sebagai response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// Menjalankan server HTTP di port 8080
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
