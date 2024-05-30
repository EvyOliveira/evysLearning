package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{}
	assert.Error(t, exercise.exerciseFieldValidator(), "invalid exercise id")
}

func TestGivenAnEmptyQuestion_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345}
	assert.Error(t, exercise.exerciseFieldValidator(), "invalid exercise question")
}

func TestGivenAnEmptyAnswer_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345, Question: "What is the best programming language?"}
	assert.Error(t, exercise.exerciseFieldValidator(), "invalid exercise answer")
}

func TestGivenAnEmptyCorrectAnswer_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345, Question: "What is the best programming language?", Answer: "Python"}
	assert.Error(t, exercise.exerciseFieldValidator(), "invalid exercise correct answer")
}

func TestGivenAnEmptySubject_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345, Question: "What is the best programming language?", Answer: "Python", CorrectAnswer: "It depends on the context, developers familiarity, software applicability and other factors."}
	assert.Error(t, exercise.exerciseFieldValidator(), "invalid exercise subject")
}

func TestGivenAnEmptyId_WhenCreateANewClass_ThenShouldReceiveAnError(t *testing.T) {
	class := class{}
	assert.Error(t, class.classFieldValidator(), "invalid class id")
}

func TestGivenAnEmptyTitle_WhenCreateANewClass_ThenShouldReceiveAnError(t *testing.T) {
	class := class{ID: 67890}
	assert.Error(t, class.classFieldValidator(), "invalid class title")
}

func TestGivenAnEmptyResume_WhenCreateANewClass_ThenShouldReceiveAnError(t *testing.T) {
	class := class{ID: 67890, Title: "Linear Programming"}
	assert.Error(t, class.classFieldValidator(), "invalid class resume")
}

func TestGivenAnEmptyText_WhenCreateANewClass_ThenShouldReceiveAnError(t *testing.T) {
	class := class{ID: 67890, Title: "Linear Programming", Resume: "Algorithms and methods for the Linear Programming approach."}
	assert.Error(t, class.classFieldValidator(), "invalid class text")
}
