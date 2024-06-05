package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simpleGoLang/model"
)

func GetCoursesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	queryParams := r.URL.Query()
	keyword := queryParams.Get("keyword")

	if keyword != "" {
		// Keyword is present, perform logic for filtering by keyword
		courses, err := model.GetCoursesByNameKeyword(keyword, db)
		if err != nil {
			http.Error(w, "Failed to fetch courses by keywrod", http.StatusInternalServerError)
			return
		}

		// Encode courses as JSON and send response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(courses); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	} else {
		// Keyword is not present, fetch all courses
		courses, err := model.GetAllCourses(db)
		if err != nil {
			http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
			return
		}

		// Encode courses as JSON and send response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(courses); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
