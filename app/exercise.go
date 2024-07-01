package app

import "errors"

type ExerciseInterface interface {
	IsValid() (bool, error)
	GetId() string
	GetQuestion() string
	GetAnswer() string
	GetCorrectAnswer() string
	GetSubject() string
}

type Exercise struct {
	ID            int64  `valid:"uuidv4"`
	Question      string `valid:"optional"`
	Answer        string `valid:"optional"`
	CorrectAnswer string `valid:"optional"`
	Subject       string `valid:"optional"`
	Classes       *Class `valid:"optional"`
}

func (e *Exercise) IsValid() (bool, error) {
	if e.Question == "" {
		return false, errors.New("invalid exercise question")
	}
	if e.Answer == "" {
		return false, errors.New("invalid exercise answer")
	}
	if e.CorrectAnswer == "" {
		return false, errors.New("invalid exercise correct answer")
	}
	if e.Subject == "" {
		return false, errors.New("invalid exercise subject")
	}
	return true, nil
}

func (e *Exercise) GetID() int64 {
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
