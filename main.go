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

type getID interface {
	getID() string
}

func (e *exercise) getID() string {
	return e.ID
}

func (c *class) getID() string {
	return c.ID
}

func (c *course) getID() string {
	return c.ID
}

var exercises []exercise
var classes []class
var courses []course

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", getAll).Methods("GET")
	route.HandleFunc("/exercises/{id}", get).Methods("GET")
	route.HandleFunc("/exercises/{id}", create).Methods("POST")
	route.HandleFunc("/exercises/{id}", update).Methods("PUT")
	route.HandleFunc("/exercises/{id}", delete).Methods("DELETE")

	route.HandleFunc("/classes/{id}", get).Methods("GET")
	route.HandleFunc("/classes/{id}", create).Methods("POST")
	route.HandleFunc("/classes/{id}", update).Methods("PUT")
	route.HandleFunc("/classes/{id}", delete).Methods("DELETE")

	route.HandleFunc("/courses/{id}", get).Methods("GET")
	route.HandleFunc("/classes/{id}", create).Methods("POST")
	route.HandleFunc("/classes/{id}", update).Methods("PUT")
	route.HandleFunc("/classes/{id}", delete).Methods("DELETE")

	fmt.Println("starting server at port:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var allData dataList
	switch r.URL.Path {
	case "/":
		allData.Exercises = exercises
		allData.Classes = classes
		allData.Courses = courses
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(allData)
	if err != nil {
		panic(err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var items []getID
	switch r.URL.Path {
	case "/exercises":
		for _, item := range exercises {
			items = append(items, &item)
		}
	case "/classes":
		for _, item := range classes {
			items = append(items, &item)
		}
	case "/courses":
		for _, item := range courses {
			items = append(items, &item)
		}
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	params := mux.Vars(r)

	for _, item := range exercises {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var item interface{}
	switch r.URL.Path {
	case "/exercises":
		item = &exercise{}
	case "/classes":
		item = &class{}
	case "/courses":
		item = &course{}
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	if item == nil {
		http.Error(w, "item not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := strconv.Itoa(rand.Intn(100000000))
	switch item.(type) {
	case *exercise:
		item.(*exercise).ID = id
		exercises = append(exercises, *item.(*exercise))
	case *class:
		item.(*class).ID = id
		classes = append(classes, *item.(*class))
	case *course:
		item.(*course).ID = id
		courses = append(courses, *item.(*course))
	}

	json.NewEncoder(w).Encode(item)
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)
	id := params["id"]

	var item interface{}
	switch r.URL.Path {
	case "/exercises" + id:
		item = &exercise{}
	case "/classes" + id:
		item = &class{}
	case "/courses" + id:
		item = &course{}
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found := false
	for index, existingItem := range exercises {
		if existingItem.ID == params["id"] {
			switch item.(type) {
			case *exercise:
				itemToUpdate := item.(*exercise)
				itemToUpdate.Question = item.(*exercise).Question
				itemToUpdate.Answers = item.(*exercise).Answers
				itemToUpdate.CorrectAnswer = item.(*exercise).CorrectAnswer
				itemToUpdate.Subject = item.(*exercise).Subject
				itemToUpdate.Classes = item.(*exercise).Classes
			case *class:
				itemToUpdate := item.(*class)
				itemToUpdate.Title = item.(*class).Title
				itemToUpdate.Resume = item.(*class).Resume
				itemToUpdate.Text = item.(*class).Text
				itemToUpdate.Course = item.(*class).Course
			case *course:
				itemToUpdate := item.(*course)
				itemToUpdate.Name = item.(*course).Name
				itemToUpdate.Description = item.(*course).Description
			}
			exercises[index] = existingItem
			item = true
			break
		}
	}
	if !found {
		http.Error(w, "item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)
	id := params["id"]

	var itemFound bool
	switch r.URL.Path {
	case "/exercises" + id:
		exercises = deleteExercise(exercises, id)
		itemFound = true
	case "/classes" + id:
		classes = deleteClass(classes, id)
		itemFound = true
	case "/courses" + id:
		courses = deleteCourse(courses, id)
		itemFound = true
	default:
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	if !itemFound {
		http.Error(w, "item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteExercise(exercises []exercise, id string) []exercise {
	newList := make([]exercise, 0)
	for _, item := range exercises {
		if item.ID != id {
			newList = append(newList, item)
		}
	}
	return newList
}

func deleteClass(classes []class, id string) []class {
	newList := make([]class, 0)
	for _, item := range classes {
		if item.ID != id {
			newList = append(newList, item)
		}
	}
	return newList
}

func deleteCourse(courses []course, id string) []course {
	newList := make([]course, 0)
	for _, item := range courses {
		if item.ID != id {
			newList = append(newList, item)
		}
	}
	return newList
}
