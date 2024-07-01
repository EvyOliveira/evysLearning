package app

import "errors"

type ClassInterface interface {
	IsValid() (bool, error)
	GetId() string
	GetTitle() string
	GetResume() string
	GetText() string
}

type Class struct {
	ID     int64   `valid:"uuidv4"`
	Title  string  `valid:"required"`
	Resume string  `valid:"required"`
	Text   string  `valid:"required"`
	Course *Course `valid:"optional"`
}

func (c *Class) IsValid() (bool, error) {
	if c.Title == "" {
		return false, errors.New("invalid class title")
	}
	if c.Resume == "" {
		return false, errors.New("invalid class resume")
	}
	if c.Text == "" {
		return false, errors.New("invalid class text")
	}
	return true, nil
}

func (c *Class) GetID() int64 {
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
