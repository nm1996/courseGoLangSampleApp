package main

import (
	"fmt"
	"log"
	"net/http"
	handler "simpleGoLang/handlers"
)

func main() {
	db := GetDBConnection()

	defer db.Close()

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		handler.GetUserByNameAndPasswordHandler(w, r, db)
	})

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		handler.InsertUserHandler(w, r, db)
	})

	http.HandleFunc("/api/course", func(w http.ResponseWriter, r *http.Request) {
		handler.GetCoursesHandler(w, r, db)
	})

	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
