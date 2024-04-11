package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type exercise struct {
	ID            string `json:"id"`
	Question      string `json:"question"`
	Answers       string `json:"answers"`
	CorrectAnswer string `json:"correct_answer"`
	Subject       string `json:"subject"`
	Classes       class  `json:"classes"`
}

type class struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Resume string `json:"resume"`
	Text   string `json:"text"`
	Course course `json:"course"`
}

type course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var exercises []exercise

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/exercises", getExercises).Methods("GET")
	route.HandleFunc("/exercises/{id}", getExercise).Methods("GET")
	route.HandleFunc("/exercises", createExercise).Methods("POST")
	route.HandleFunc("/exercises/", updateExercise).Methods("PUT")
	route.HandleFunc("/exercises/", deleteExercise).Methods("DELETE")

	fmt.Println("starting server at port:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(exercises)
	if err != nil {
		panic(err)
	}
}

func getExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)

	for _, item := range exercises {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var exercise exercise

	_ = json.NewDecoder(r.Body).Decode(&exercise)
	exercise.ID = strconv.Itoa(rand.Intn(100000000))

	exercises = append(exercises, exercise)
	json.NewEncoder(w).Encode(exercise)
}

func updateExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)

	for index, item := range exercises {
		if item.ID == params["id"] {
			exercises = append(exercises[:index], exercises[index+1:]...)

			var exercise exercise
			_ = json.NewDecoder(r.Body).Decode(&exercise)
			exercise.ID = params["id"]
			exercises = append(exercises, exercise)
			json.NewEncoder(w).Encode(exercise)
			return
		}
	}
}

func deleteExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)

	for index, item := range exercises {
		if item.ID == params["id"] {
			exercises = append(exercises[:index], exercises[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(exercises)
}
