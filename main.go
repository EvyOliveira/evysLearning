package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type exercise struct {
	ID            uint   `json:"id"`
	Question      string `json:"question"`
	Answers       string `json:"answers"`
	CorrectAnswer string `json:"correct_answer"`
	Subject       string `json:"subject"`
	//Classes       class  `json:"classes"`
}

type class struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Resume string `json:"resume"`
	Text   string `json:"text"`
	Course course `json:"course"`
}

type course struct {
	ID          uint `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var exercises = []exercise{
	{ID: 1, Question: "2", Answers: "D", CorrectAnswer: "E", Subject: "Math"},
	{ID: 20, Question: "2", Answers: "A", CorrectAnswer: "C", Subject: "Biology"},
}


func main() {
	http.HandleFunc("/", getAll)
	fmt.Println("api is on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Not success", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var response []exercise

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("erro ao recuperar dados", err.Error())
		return
	}

	fmt.Println(response)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(exercises)
	if err != nil {
		panic(err)
	}
}
