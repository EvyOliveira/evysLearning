package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var exercises []exercise
var singleExercise exercise

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

type databaseConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func loadDatabaseConfig() *databaseConfiguration {
	config := &databaseConfiguration{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.Database == "" {
		log.Fatalf("missing required environment variables")
	}

	return config
}

func openDatabaseConnection(config *databaseConfiguration) (*sql.DB, error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	databaseConnection, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	err = databaseConnection.Ping()
	return databaseConnection, err
}

func main() {
	router := gin.Default()

	exercises = []exercise{
		{ID: "1", Question: "user1", Answers: "user1@example.com", CorrectAnswer: "fffff", Subject: "g"},
		{ID: "2", Question: "user2", Answers: "user2@example.com", CorrectAnswer: "hhhhhh", Subject: "u"},
	}

	router.GET("/exercises", getExercises)
	router.GET("/exercises/:id", getExerciseByID)
	router.POST("/exercises/:id}", createExercise)
	router.PUT("/exercises/:id", updateExercise)
	router.DELETE("/exercises/:id", deleteExercise)

	fmt.Println("starting server at port:8080")
	log.Fatal(router.Run(":8080"))
}

func getExercises(c *gin.Context) {
	config, err := loadDatabaseConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	databaseConnection, err := openDatabaseConnection(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer databaseConnection.Close()

	rows, err := databaseConnection.Query(`SELECT * FROM exercises`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&singleExercise.ID, &singleExercise.Question, &singleExercise.Answers, &singleExercise.CorrectAnswer, &singleExercise.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		exercises = append(exercises, singleExercise)
	}
	c.JSON(http.StatusOK, exercises)
}

func getExerciseByID(c *gin.Context) {
	config, err := loadDatabaseConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	databaseConnection, err := openDatabaseConnection(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer databaseConnection.Close()

	id := c.Param("id")

	row := databaseConnection.QueryRow(`SELECT * FROM exercises WHERE id = $1`, id)
	err = row.Scan(&singleExercise.ID, &singleExercise.Question, &singleExercise.Answers, &singleExercise.CorrectAnswer, &singleExercise.Subject)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, singleExercise)
}

func createExercise(c *gin.Context) {
	if err := c.BindJSON(&singleExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error to bind create exercise": err.Error()})
		return
	}

	config, err := loadDatabaseConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	databaseConnection, err := openDatabaseConnection(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer databaseConnection.Close()

	result, err := databaseConnection.Exec(`INSERT INTO exercises (question, answers, correct_answer, subject) VALUES ($1, $2, $3, $4)`, singleExercise.Question, singleExercise.Answers, singleExercise.CorrectAnswer, singleExercise.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to create exercise"})
		return
	}

	if rowsAffect != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to create exercise"})
		return
	}

	c.JSON(http.StatusCreated, singleExercise)
}

func updateExercise(c *gin.Context) {
	id := c.Param("id")

	if err := c.BindJSON(&singleExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := loadDatabaseConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	databaseConnection, err := openDatabaseConnection(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer databaseConnection.Close()

	result, err := databaseConnection.Exec(`UPDATE exercises SET question = $1, answers = $2, correct_answer = $3, subject = $4 WHERE id = $5`,
		singleExercise.Question, singleExercise.Answers, singleExercise.CorrectAnswer, singleExercise.Subject, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "exercise updated successfully"})
}

func deleteExercise(c *gin.Context) {
	id := c.Param("id")

	config, err := loadDatabaseConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	databaseConnection, err := openDatabaseConnection(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer databaseConnection.Close()

	result, err := databaseConnection.Exec(`DELETE FROM exercises WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
