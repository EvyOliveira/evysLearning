package app

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type ExerciseInterface interface {
	IsValid() (bool, error)
	GetId() string
	GetQuestion() string
	GetAnswer() string
	GetCorrectAnswer() string
	GetSubject() string
}

type Exercise struct {
	ID            string `valid:"uuidv4"`
	Question      string `valid:"required"`
	Answer        string `valid:"required"`
	CorrectAnswer string `valid:"required"`
	Subject       string `valid:"required"`
	Classes       *Class `valid:"required"`
}

func NewExercise(question, answer, correctAnswer, subject string) *Exercise {
	exercise := Exercise{
		ID:            uuid.NewV4().String(),
		Question:      question,
		Answer:        answer,
		CorrectAnswer: correctAnswer,
		Subject:       subject,
	}
	return &exercise
}

func (e *Exercise) IsValid() (bool, error) {
	if e.Question == "" {
		return false, errors.New("empty exercise question")
	}
	if e.Answer == "" {
		return false, errors.New("empty exercise answer")
	}
	if e.CorrectAnswer == "" {
		return false, errors.New("empty exercise correct answer")
	}
	if e.Subject == "" {
		return false, errors.New("empty exercise subject")
	}
	return true, nil
}

func (e *Exercise) GetID() string {
	return e.ID
}

func (e *Exercise) GetQuestion() string {
	return e.Question
}

func (e *Exercise) GetAnswer() string {
	return e.Answer
}

func (e *Exercise) GetCorrectAnswer() string {
	return e.CorrectAnswer
}

func (e *Exercise) GetSubject() string {
	return e.Subject
}
