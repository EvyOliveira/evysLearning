package app_test

import (
	app "evys-learning/app/entity"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestExercise_IsValid(t *testing.T) {
	exercise := app.Exercise{}
	exercise.ID = uuid.NewV4().String()
	exercise.Question = "What is the best programming language?"
	exercise.Answer = "Python"
	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = "Programming Language"

	_, err := exercise.IsValid()
	require.Nil(t, err)

	exercise.Question = ""
	_, err = exercise.IsValid()
	require.Equal(t, "empty exercise question", err.Error())

	exercise.Question = "What is the best programming language?"
	exercise.Answer = ""
	_, err = exercise.IsValid()
	require.Equal(t, "empty exercise answer", err.Error())

	exercise.Answer = "Python"
	exercise.CorrectAnswer = ""
	_, err = exercise.IsValid()
	require.Equal(t, "empty exercise correct answer", err.Error())

	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = ""
	_, err = exercise.IsValid()
	require.Equal(t, "empty exercise subject", err.Error())
}

func TestExercise_GetID(t *testing.T) {
	exercise := app.Exercise{}
	exercise.ID = uuid.NewV4().String()
	exercise.Question = "What is the best programming language?"
	exercise.Answer = "Python"
	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = "Programming Language"

	ID := exercise.GetID()
	require.Equal(t, exercise.ID, ID)
}

func TestExercise_GetQuestion(t *testing.T) {
	exercise := app.Exercise{}
	exercise.ID = uuid.NewV4().String()
	exercise.Question = "What is the best programming language?"
	exercise.Answer = "Python"
	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = "Programming Language"

	question := exercise.GetQuestion()
	require.Equal(t, exercise.Question, question)
}

func TestExercise_GetAnswer(t *testing.T) {
	exercise := app.Exercise{}
	exercise.ID = uuid.NewV4().String()
	exercise.Question = "What is the best programming language?"
	exercise.Answer = "Python"
	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = "Programming Language"

	answer := exercise.GetAnswer()
	require.Equal(t, exercise.Answer, answer)
}

func TestExercise_GetCorrectAnswer(t *testing.T) {
	exercise := app.Exercise{}
	exercise.ID = uuid.NewV4().String()
	exercise.Question = "What is the best programming language?"
	exercise.Answer = "Python"
	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = "Programming Language"

	correctAnswer := exercise.GetCorrectAnswer()
	require.Equal(t, exercise.CorrectAnswer, correctAnswer)
}

func TestExercise_GetSubject(t *testing.T) {
	exercise := app.Exercise{}
	exercise.ID = uuid.NewV4().String()
	exercise.Question = "What is the best programming language?"
	exercise.Answer = "Python"
	exercise.CorrectAnswer = "It depends on the context, developers familiarity, software applicability and other factors."
	exercise.Subject = "Programming Language"

	subject := exercise.GetSubject()
	require.Equal(t, exercise.Subject, subject)
}
