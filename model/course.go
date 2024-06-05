package model

import (
	"database/sql"
	"encoding/json"
)

type Course struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Areas       []string `json:"areas"`
	Description string   `json:"description"`
}

func GetAllCourses(db *sql.DB) ([]Course, error) {
	var courses []Course

	rows, err := db.Query("SELECT id, name, price, areas, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course Course
		var areasJSON []byte
		if err := rows.Scan(&course.ID, &course.Name, &course.Price, &areasJSON, &course.Description); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(areasJSON, &course.Areas); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func GetCoursesByNameKeyword(keyword string, db *sql.DB) ([]Course, error) {
	var courses []Course

	rows, err := db.Query("SELECT id, name, price, areas, description FROM courses WHERE name LIKE ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course Course
		var areasJSON []byte
		if err := rows.Scan(&course.ID, &course.Name, &course.Price, &areasJSON, &course.Description); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(areasJSON, &course.Areas); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
