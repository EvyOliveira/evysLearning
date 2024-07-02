package app_test

import (
	app "evys-learning/app/entity"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestCourse_IsValid(t *testing.T) {
	course := app.Course{}
	course.ID = uuid.NewV4().String()
	course.Name = "Analysis and systems development"
	course.Description = "Focus on developing systems design and development skills."

	_, err := course.IsValid()
	require.Nil(t, err)

	course.Name = ""
	_, err = course.IsValid()
	require.Equal(t, "empty course name", err.Error())

	course.Description = ""
	course.Name = "Analysis and systems development"
	_, err = course.IsValid()
	require.Equal(t, "empty course description", err.Error())

}

func TestCourse_GetID(t *testing.T) {
	course := app.Course{}
	course.ID = uuid.NewV4().String()
	course.Name = "Analysis and systems development"
	course.Description = "Focus on developing systems design and development skills."

	ID := course.GetID()
	require.Equal(t, course.ID, ID)
}

func TestCourse_GetName(t *testing.T) {
	course := app.Course{}
	course.ID = uuid.NewV4().String()
	course.Name = "Analysis and systems development"
	course.Description = "Focus on developing systems design and development skills."

	name := course.GetName()
	require.Equal(t, course.Name, name)
}

func TestCourse_GetDescription(t *testing.T) {
	course := app.Course{}
	course.ID = uuid.NewV4().String()
	course.Name = "Analysis and systems development"
	course.Description = "Focus on developing systems design and development skills."

	description := course.GetDescription()
	require.Equal(t, course.Description, description)
}
