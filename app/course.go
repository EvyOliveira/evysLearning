package app

import "errors"

type CourseInterface interface {
	IsValid() (bool, error)
	GetId() string
	GetName() string
	GetDescription() string
}

type Course struct {
	ID          int64  `valid:"uuidv4"`
	Name        string `valid:"optional"`
	Description string `valid:"optional"`
}

func (c *Course) IsValid() (bool, error) {
	if c.Name == "" {
		return false, errors.New("invalid course name")
	}
	if c.Description == "" {
		return false, errors.New("invalid course description")
	}
	return true, nil
}

func (c *Course) GetID() int64 {
	return c.ID
}

func (c *Course) GetName() string {
	return c.Name
}

func (c *Course) GetDescription() string {
	return c.Description
}
