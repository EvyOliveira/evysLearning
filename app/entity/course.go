package app

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type CourseInterface interface {
	IsValid() (bool, error)
	GetId() string
	GetName() string
	GetDescription() string
}

type Course struct {
	ID          string `valid:"uuidv4"`
	Name        string `valid:"required"`
	Description string `valid:"required"`
}

func NewCourse(name, description string) *Course {
	course := Course{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: description,
	}
	return &course
}

func (c *Course) IsValid() (bool, error) {
	if c.Name == "" {
		return false, errors.New("empty course name")
	}
	if c.Description == "" {
		return false, errors.New("empty course description")
	}
	return true, nil
}

func (c *Course) GetID() string {
	return c.ID
}

func (c *Course) GetName() string {
	return c.Name
}

func (c *Course) GetDescription() string {
	return c.Description
}
