package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	exercises exercise
	classes   class
	courses   course
	id        uint32
	err       error
	config    *configuration
	data      []dataList
)

type exercise struct {
	ID            uint32 `json:"id"`
	Question      string `json:"question"`
	Answers       string `json:"answers"`
	CorrectAnswer string `json:"correct_answer"`
	Subject       string `json:"subject"`
	Classes       *class `json:"classes"`
}

type class struct {
	ID     uint32  `json:"id"`
	Title  string  `json:"title"`
	Resume string  `json:"resume"`
	Text   string  `json:"text"`
	Course *course `json:"course"`
}

type course struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type dataList struct {
	Exercises []exercise `json:"exercises"`
	Classes   []class    `json:"classes"`
	Courses   []course   `json:"courses"`
}

type configuration struct {
	API apiConfiguration
	DB  dbConfiguration
}

type apiConfiguration struct {
	Port string
}

type dbConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func main() {
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

	fmt.Println("starting server at port:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func init() {
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func load() error {
	viper.SetConfigFile("config")
	viper.SetConfigFile("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	configuration := new(configuration)
	configuration.API = apiConfiguration{
		Port: viper.GetString("api.port"),
	}

	configuration.DB = dbConfiguration{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func getDB() dbConfiguration {
	return config.DB
}

func getServerPort() string {
	return config.API.Port
}

func openConnection() (*sql.DB, error) {
	conf := getDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disbale",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	databaseConnection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = databaseConnection.Ping()
	return databaseConnection, err
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("successful connection")

	return db
}

func getAll(w http.ResponseWriter, r *http.Request) {
	databaseConnection, err := openConnection()
	if err != nil {
		return
	}
	defer databaseConnection.Close()

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

	databaseConnection, err := openConnection()
	if err != nil {
		return
	}
	defer databaseConnection.Close()

	switch r.URL.Path {
	case "/exercises":
		row := databaseConnection.QueryRow(`SELECT * FROM exercises WHERE id=$1`, id)
		err = row.Scan(&exercises.ID, &exercises.Question, &exercises.Answers, &exercises.CorrectAnswer, &exercises.Subject, &exercises.Subject, &exercises.Classes)
	case "/classes":
		row := databaseConnection.QueryRow(`SELECT * FROM classes WHERE id=$1`, id)
		err = row.Scan(&classes.ID, &classes.Title, &classes.Resume, &classes.Text, &classes.Course)
	case "/courses":
		row := databaseConnection.QueryRow(`SELECT * FROM classes WHERE id=$1`, id)
		err = row.Scan(&courses.ID, &courses.Name, &courses.Description)
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	return
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	databaseConnection, err := openConnection()
	if err != nil {
		return
	}
	defer databaseConnection.Close()

	switch r.URL.Path {
	case "/exercises":
		row := `INSERT INTO exercises (questions, answers, correctAnswer, subject, classes) VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err = databaseConnection.QueryRow(row, exercises.Question, exercises.Answers, exercises.CorrectAnswer, exercises.Subject, exercises.Classes).Scan(&id)
	case "/classes":
		row := `INSERT INTO classes (title, resume, text, course) VALUES ($1, $2, $3, $4) RETURNING id`
		err = databaseConnection.QueryRow(row, classes.Title, classes.Resume, classes.Text, classes.Course).Scan(&id)
	case "/courses":
		row := `INSERT INTO courses (name, description) VALUES ($1, $2) RETURNING id`
		err = databaseConnection.QueryRow(row, courses.Name, courses.Description).Scan(&id)
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}
	return
}

func update(w http.ResponseWriter, r *http.Request) {
	_, err = updateItem(w, r)
	if err != nil {
		panic(err)
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) (int64, error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	databaseConnection, err := openConnection()
	if err != nil {
		return 0, err
	}
	defer databaseConnection.Close()

	params := mux.Vars(r)
	id := params["id"]

	switch r.URL.Path {
	case "/exercises" + id:
		row, err := databaseConnection.Exec(`UPDATE exercises SET question=$1, answers=$2, correct_answer=$3, subject=$4`,
			exercises.Question, exercises.Answers, exercises.CorrectAnswer, exercises.Subject)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case "/classes" + id:
		row, err := databaseConnection.Exec(`UPDATE classes SET title=$1, resume=$2, text=$3, course=$4`,
			classes.Title, classes.Resume, classes.Text, classes.Course)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case "/courses" + id:
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

	databaseConnection, err := openConnection()
	if err != nil {
		return 0, err
	}
	defer databaseConnection.Close()

	params := mux.Vars(r)
	id := params["id"]

	switch r.URL.Path {
	case "/exercises" + id:
		row, err := databaseConnection.Exec(`DELETE FROM exercises WHERE id=$1`, id)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case "/classes" + id:
		row, err := databaseConnection.Exec(`DELETE FROM classes WHERE id=$1`, id)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	case "/courses" + id:
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
