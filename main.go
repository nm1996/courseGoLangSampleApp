package main

import (
	"fmt"
	"log"
	"net/http"
	handler "simpleGoLang/handlers"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	db := GetDBConnection()
	defer db.Close()

	http.Handle("/api", corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login":
			handler.GetUserByNameAndPasswordHandler(w, r, db)
		case "/register":
			handler.InsertUserHandler(w, r, db)
		case "/course":
			handler.GetCoursesHandler(w, r, db)
		default:
			http.NotFound(w, r)
			return
		}
	})))

	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
