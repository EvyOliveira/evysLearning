package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	exercisesPath = "/exercises"
	classesPath   = "/classes"
	coursesPath   = "/courses"
)

var (
	exercises exercise
	classes   class
	courses   course
	id        uint32
	err       error
	data      []dataList
)

type exercise struct {
	ID            string `json:"id"`
	Question      string `json:"question"`
	Answers       string `json:"answers"`
	CorrectAnswer string `json:"correct_answer"`
	Subject       string `json:"subject"`
	Classes       *class `json:"classes"`
}

type class struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Resume string  `json:"resume"`
	Text   string  `json:"text"`
	Course *course `json:"course"`
}

type course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type dataList struct {
	Exercises []exercise `json:"exercises"`
	Classes   []class    `json:"classes"`
	Courses   []course   `json:"courses"`
}

func NewExercise(question, answers, correctAnswer, subject string) *exercise {
	return &exercise{
		ID:            uuid.New().String(),
		Question:      question,
		Answers:       answers,
		CorrectAnswer: correctAnswer,
		Subject:       subject,
	}
}

func NewClass(title, resume, text string) *class {
	return &class{
		ID:     uuid.New().String(),
		Title:  title,
		Resume: resume,
		Text:   text,
	}
}

func NewCourse(name, description string) *course {
	return &course{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}

func main() {
	databaseConnection()
	route := mux.NewRouter()

	route.HandleFunc("/", getAll).Methods("GET")
	route.HandleFunc("/exercises/{id}", getById).Methods("GET")
	route.HandleFunc("/exercises/{id}", create).Methods("POST")
	route.HandleFunc("/exercises/{id}", update).Methods("PUT")
	route.HandleFunc("/exercises/{id}", delete).Methods("DELETE")

	route.HandleFunc("/classes/{id}", getById).Methods("GET")
	route.HandleFunc("/classes/{id}", create).Methods("POST")
	route.HandleFunc("/classes/{id}", update).Methods("PUT")
	route.HandleFunc("/classes/{id}", delete).Methods("DELETE")

	route.HandleFunc("/courses/{id}", getById).Methods("GET")
	route.HandleFunc("/classes/{id}", create).Methods("POST")
	route.HandleFunc("/classes/{id}", update).Methods("PUT")
	route.HandleFunc("/classes/{id}", delete).Methods("DELETE")

	fmt.Println("starting server at port:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func databaseConnection() (*sql.DB, error) {
	databaseConnection, err := sql.Open("postgres", "root:root@tcp(localhost:8000)/goexpert")
	if err != nil {
		log.Fatal("there is an error to connect to database.")
	}
	defer databaseConnection.Close()
	return databaseConnection, nil
}

func getAll(w http.ResponseWriter, r *http.Request) {
	databaseConnection, err := databaseConnection()
	if err != nil {
		return
	}

	rows, err := databaseConnection.Query(`SELECT * FROM datalist`)
	if err != nil {
		return
	}

	for rows.Next() {
		var allData dataList
		err = rows.Scan(&allData.Exercises, &allData.Classes, &allData.Courses)
		if err != nil {
			continue
		}
		data = append(data, allData)
	}
}

func getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	databaseConnection, err := databaseConnection()
	if err != nil {
		return
	}

	switch r.URL.Path {
	case exercisesPath:
		row := databaseConnection.QueryRow(`SELECT * FROM exercises WHERE id=$1`, id)
		_ = row.Scan(&exercises.ID, &exercises.Question, &exercises.Answers, &exercises.CorrectAnswer, &exercises.Subject, &exercises.Subject, &exercises.Classes)
		return
	case classesPath:
		row := databaseConnection.QueryRow(`SELECT * FROM classes WHERE id=$1`, id)
		_ = row.Scan(&classes.ID, &classes.Title, &classes.Resume, &classes.Text, &classes.Course)
		return
	case coursesPath:
		row := databaseConnection.QueryRow(`SELECT * FROM classes WHERE id=$1`, id)
		_ = row.Scan(&courses.ID, &courses.Name, &courses.Description)
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	databaseConnection, err := databaseConnection()
	if err != nil {
		return
	}

	switch r.URL.Path {
	case exercisesPath:
		row := `INSERT INTO exercises (questions, answers, correctAnswer, subject, classes) VALUES ($1, $2, $3, $4, $5) RETURNING id`
		_ = databaseConnection.QueryRow(row, exercises.Question, exercises.Answers, exercises.CorrectAnswer, exercises.Subject, exercises.Classes).Scan(&id)
		return
	case classesPath:
		row := `INSERT INTO classes (title, resume, text, course) VALUES ($1, $2, $3, $4) RETURNING id`
		_ = databaseConnection.QueryRow(row, classes.Title, classes.Resume, classes.Text, classes.Course).Scan(&id)
		return
	case coursesPath:
		row := `INSERT INTO courses (name, description) VALUES ($1, $2) RETURNING id`
		_ = databaseConnection.QueryRow(row, courses.Name, courses.Description).Scan(&id)
		return
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	_, err = updateItem(w, r)
	if err != nil {
		panic(err)
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) (int64, error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	databaseConnection, err := databaseConnection()
	if err != nil {
		return 0, err
	}

	params := mux.Vars(r)
	id := params["id"]

	switch r.URL.Path {
	case exercisesPath + id:
		row, err := databaseConnection.Exec(`UPDATE exercises SET question=$1, answers=$2, correct_answer=$3, subject=$4`,
			exercises.Question, exercises.Answers, exercises.CorrectAnswer, exercises.Subject)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case classesPath + id:
		row, err := databaseConnection.Exec(`UPDATE classes SET title=$1, resume=$2, text=$3, course=$4`,
			classes.Title, classes.Resume, classes.Text, classes.Course)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case coursesPath + id:
		row, err := databaseConnection.Exec(`UPDATE courses SET name=$1, description=$2`,
			courses.Name, courses.Description)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return 0, nil
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	_, err = deleteItem(w, r)
	if err != nil {
		panic(err)
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) (int64, error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	databaseConnection, err := databaseConnection()
	if err != nil {
		return 0, err
	}

	params := mux.Vars(r)
	id := params["id"]

	switch r.URL.Path {
	case exercisesPath + id:
		row, err := databaseConnection.Exec(`DELETE FROM exercises WHERE id=$1`, id)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case classesPath + id:
		row, err := databaseConnection.Exec(`DELETE FROM classes WHERE id=$1`, id)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case coursesPath + id:
		row, err := databaseConnection.Exec(`DELETE FROM courses WHERE id=$1`, id)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return 0, nil
	}
}
