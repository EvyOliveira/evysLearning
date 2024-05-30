package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{}
	assert.Error(t, exercise.isValid(), "invalid exercise id")
}

func TestGivenAnEmptyQuestion_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345}
	assert.Error(t, exercise.isValid(), "invalid exercise question")
}

func TestGivenAnEmptyAnswer_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345, Question: "What is the best programming language?"}
	assert.Error(t, exercise.isValid(), "invalid exercise answer")
}

func TestGivenAnEmptyCorrectAnswer_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345, Question: "What is the best programming language?", Answer: "Python"}
	assert.Error(t, exercise.isValid(), "invalid exercise correct answer")
}

func TestGivenAnEmptySubject_WhenCreateANewExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := exercise{ID: 12345, Question: "What is the best programming language?", Answer: "Python", CorrectAnswer: "It depends on the context, developers familiarity, software applicability and other factors."}
	assert.Error(t, exercise.isValid(), "invalid exercise subject")
}
