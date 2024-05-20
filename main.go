package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var exercises []exercise

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

	fmt.Println("starting server at port:8000")
	log.Fatal(router.Run(":8000"))
}

func getExercises(c *gin.Context) {
	c.JSON(http.StatusOK, exercises)
}

func getExerciseByID(c *gin.Context) {
	id := c.Param("id")

	for _, exercise := range exercises {
		if exercise.ID == id {
			c.JSON(http.StatusOK, exercise)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
}

func createExercise(c *gin.Context) {
	var newExercise exercise
	if err := c.BindJSON(&newExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error to bind create exercise": err.Error()})
		return
	}
	exercises = append(exercises, newExercise)
	c.JSON(http.StatusCreated, newExercise)
}

func updateExercise(c *gin.Context) {
	id := c.Param("id")
	var updatedExercise exercise
	if err := c.BindJSON(&updatedExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error to bind update exercise": err.Error()})
		return
	}
	for i, exercise := range exercises {
		if exercise.ID == id {
			exercises[i] = updatedExercise
			c.JSON(http.StatusOK, updatedExercise)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
}

func deleteExercise(c *gin.Context) {
	id := c.Param("id")
	var newExercises []exercise

	for _, exercise := range exercises {
		if exercise.ID != id {
			newExercises = append(newExercises, exercise)
		}
	}

	if len(newExercises) != len(exercises) {
		exercises = newExercises
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
	}
}
