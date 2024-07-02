package app

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type ClassInterface interface {
	IsValid() (bool, error)
	GetId() string
	GetTitle() string
	GetResume() string
	GetText() string
}

type Class struct {
	ID     string  `valid:"uuidv4"`
	Title  string  `valid:"required"`
	Resume string  `valid:"required"`
	Text   string  `valid:"required"`
	Course *Course `valid:"optional"`
}

func NewClass(title, resume, text string) *Class {
	class := Class{
		ID:     uuid.NewV4().String(),
		Title:  title,
		Resume: resume,
		Text:   text,
	}
	return &class
}

func (c *Class) IsValid() (bool, error) {
	if c.Title == "" {
		return false, errors.New("empty class title")
	}
	if c.Resume == "" {
		return false, errors.New("empty class resume")
	}
	if c.Text == "" {
		return false, errors.New("empty class text")
	}
	return true, nil
}

func (c *Class) GetID() string {
	return c.ID
}

func (c *Class) GetTitle() string {
	return c.Title
}

func (c *Class) GetResume() string {
	return c.Resume
}

func (c *Class) GetText() string {
	return c.Text
}
