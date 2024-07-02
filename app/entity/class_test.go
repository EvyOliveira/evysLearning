package app_test

import (
	app "evys-learning/app/entity"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestClass_IsValid(t *testing.T) {
	class := app.Class{}
	class.ID = uuid.NewV4().String()
	class.Title = "Linear Programming"
	class.Resume = "Algorithms and methods for the Linear Programming approach."
	class.Text = "Elective subject"

	_, err := class.IsValid()
	require.Nil(t, err)

	class.Title = ""
	_, err = class.IsValid()
	require.Equal(t, "empty class title", err.Error())

	class.Resume = ""
	class.Title = "Linear Programming"
	_, err = class.IsValid()
	require.Equal(t, "empty class resume", err.Error())

	class.Text = ""
	class.Resume = "Algorithms and methods for the Linear Programming approach."
	_, err = class.IsValid()
	require.Equal(t, "empty class text", err.Error())
}

func TestClass_GetID(t *testing.T) {
	class := app.Class{}
	class.ID = uuid.NewV4().String()
	class.Title = "Linear Programming"
	class.Resume = "Algorithms and methods for the Linear Programming approach."
	class.Text = "Elective subject"

	ID := class.GetID()
	require.Equal(t, class.ID, ID)
}

func TestClass_GetTitle(t *testing.T) {
	class := app.Class{}
	class.ID = uuid.NewV4().String()
	class.Title = "Linear Programming"
	class.Resume = "Algorithms and methods for the Linear Programming approach."
	class.Text = "Elective subject"

	title := class.GetTitle()
	require.Equal(t, class.Title, title)
}

func TestClass_GetResume(t *testing.T) {
	class := app.Class{}
	class.ID = uuid.NewV4().String()
	class.Title = "Linear Programming"
	class.Resume = "Algorithms and methods for the Linear Programming approach."
	class.Text = "Elective subject"

	resume := class.GetResume()
	require.Equal(t, class.Resume, resume)
}

func TestClass_GetText(t *testing.T) {
	class := app.Class{}
	class.ID = uuid.NewV4().String()
	class.Title = "Linear Programming"
	class.Resume = "Algorithms and methods for the Linear Programming approach."
	class.Text = "Elective subject"

	text := class.GetText()
	require.Equal(t, class.Text, text)
}
