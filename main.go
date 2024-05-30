package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type exercise struct {
	ID            int64  `json:"id"`
	Question      string `json:"question"`
	Answer        string `json:"answer"`
	CorrectAnswer string `json:"correct_answer"`
	Subject       string `json:"subject"`
	Classes       *class `json:"classes"`
}

type class struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Resume string  `json:"resume"`
	Text   string  `json:"text"`
	Course *course `json:"course"`
}

type course struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewExercise(id int64, question, answer, correctAnswer, subject string) (*exercise, error) {
	exercise := &exercise{
		ID:            id,
		Question:      question,
		Answer:        answer,
		CorrectAnswer: correctAnswer,
		Subject:       subject,
	}
	err := exercise.isValid()
	if err != nil {
		return nil, err
	}
	return exercise, nil
}

func (e *exercise) isValid() error {
	if e.ID < 0 {
		return errors.New("invalid exercise id")
	}
	if e.Question == "" {
		return errors.New("invalid exercise question")
	}
	if e.Answer == "" {
		return errors.New("invalid exercise answer")
	}
	if e.CorrectAnswer == "" {
		return errors.New("invalid exercise correct answer")
	}
	if e.Subject == "" {
		return errors.New("invalid exercise subject")
	}
	return nil
}

func databaseConnection() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connection := fmt.Sprint("postgres", "host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	databaseConnection, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return databaseConnection, err
}

func main() {
	router := gin.Default()

	router.GET("/exercises", getExercises)
	router.GET("/classes", getClasses)
	router.GET("/courses", getCourses)

	router.GET("/exercises/:id", getExerciseByID)
	router.GET("/classes/:id", getClassByID)
	router.GET("/courses/:id", getCourseByID)

	router.POST("/exercises/:id}", createExercise)
	router.POST("/classes/:id}", createClass)
	router.POST("/courses/:id}", createCourse)

	router.PUT("/exercises/:id", updateExercise)
	router.PUT("/classes/:id", updateClass)
	router.PUT("/courses/:id", updateCourse)

	router.DELETE("/exercises/:id", deleteExercise)
	router.DELETE("/classes/:id", deleteClass)
	router.DELETE("/courses/:id", deleteCourse)

	fmt.Println("starting server at port:8080")
	log.Fatal(router.Run(":8080"))
}

func getExercises(c *gin.Context) {
	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM exercises")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var allExercises []exercise
	for rows.Next() {
		var ex exercise
		err := rows.Scan(&ex.ID, &ex.Question, &ex.Answer, &ex.CorrectAnswer, &ex.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allExercises = append(allExercises, ex)
	}
	c.JSON(http.StatusOK, allExercises)
}

func getClasses(c *gin.Context) {
	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM classes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var allClasses []class
	for rows.Next() {
		var cl class
		err := rows.Scan(&cl.ID, &cl.Title, &cl.Resume, &cl.Text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allClasses = append(allClasses, cl)
	}
	c.JSON(http.StatusOK, allClasses)
}

func getCourses(c *gin.Context) {
	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM courses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var allCourses []course
	for rows.Next() {
		var co course
		err := rows.Scan(&co.ID, &co.Name, &co.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allCourses = append(allCourses, co)
	}
	c.JSON(http.StatusOK, allCourses)
}

func getExerciseByID(c *gin.Context) {
	exerciseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise ID"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM exercises WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var exercise exercise
	row := stmt.QueryRow(exerciseID)
	err = row.Scan(&exercise.ID, &exercise.Question, &exercise.Answer, &exercise.CorrectAnswer, &exercise.Subject)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func getClassByID(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid class ID"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM classes WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var class class
	row := stmt.QueryRow(classID)
	err = row.Scan(&class.ID, &class.Title, &class.Resume, &class.Text)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "class not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func getCourseByID(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM courses WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var course course
	row := stmt.QueryRow(courseID)
	err = row.Scan(&course.ID, &course.Name, &course.Description)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func createExercise(c *gin.Context) {
	var exercise exercise
	err := json.NewDecoder(c.Request.Body).Decode(&exercise)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise data"})
		return
	}

	exercise.ID = int64(uuid.New().ID())

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO exercises (question, answer, correct_answer, subject) VALUES ($1, $2, $3, $4)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(exercise.Question, exercise.Answer, exercise.CorrectAnswer, exercise.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, exercise)
}

func createClass(c *gin.Context) {
	var class class
	err := json.NewDecoder(c.Request.Body).Decode(&class)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid class data"})
		return
	}

	class.ID = int64(uuid.New().ID())

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO classes (title, resume, text) VALUES ($1, $2, $3)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(class.Title, class.Resume, class.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

func createCourse(c *gin.Context) {
	var course course
	err := json.NewDecoder(c.Request.Body).Decode(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course data"})
		return
	}

	course.ID = int64(uuid.New().ID())

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO courses (name, description) VALUES ($1, $2)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(course.Name, course.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func updateExercise(c *gin.Context) {
	exerciseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise ID"})
		return
	}

	var updatedExercise exercise
	err = json.NewDecoder(c.Request.Body).Decode(&updatedExercise)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise data"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE exercises SET question = $1, answer = $2, correct_answer = $3, subject = $4 WHERE id = $5")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(updatedExercise.Question, updatedExercise.Answer, updatedExercise.CorrectAnswer, updatedExercise.Subject, exerciseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating exercise: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedExercise)
}

func updateClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid class ID"})
		return
	}

	var updatedClass class
	err = json.NewDecoder(c.Request.Body).Decode(&updatedClass)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid class data"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE classes SET title = $1, resume = $2, text = $3 WHERE id = $4")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(updatedClass.Title, updatedClass.Resume, updatedClass.Text, classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating class: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedClass)
}

func updateCourse(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	var updatedCourse course
	err = json.NewDecoder(c.Request.Body).Decode(&updatedCourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course data"})
		return
	}

	updatedCourse.ID = int64(courseID)

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE courses SET name = $1, description = $2 WHERE id = $3")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(updatedCourse.Name, updatedCourse.Description, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating course: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCourse)
}

func deleteExercise(c *gin.Context) {
	exerciseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise id"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM exercises WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(strconv.Itoa(exerciseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting exercise: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "exercise deleted"})
}

func deleteClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid class id"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM classes WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(strconv.Itoa(classID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting class: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "class deleted"})
}

func deleteCourse(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}

	db, err := databaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM courses WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = db.Exec(strconv.Itoa(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting course: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course deleted"})
}
