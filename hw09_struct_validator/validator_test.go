package hw09structvalidator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidateUser(t *testing.T) {
	t.Run("Successful validation with minimum values", func(t *testing.T) {
		ve := Validate(User{
			ID:     "012345678901234567890123456789012345",
			Name:   "N",
			Age:    18,
			Email:  "t@m.g",
			Role:   "stuff",
			Phones: []string{"89999999999"},
			meta:   json.RawMessage("Some meta"),
		})

		require.Equal(t, ve, nil)
	})

	t.Run("Successful validation with maximum values", func(t *testing.T) {
		ve := Validate(User{
			ID:     "012345678901234567890123456789012345",
			Name:   "Nick012345678901234567890123456789012345012345678901234567890123456789012345012345678901234567890",
			Age:    50,
			Email:  "test012345678901234567890123456789@mail.go0123450123456789012345678901234567890123450123456789",
			Role:   "admin",
			Phones: []string{"81234567890"},
			meta:   json.RawMessage("Some meta"),
		})

		require.Equal(t, ve, nil)
	})

	t.Run("Failed validation with minimum values", func(t *testing.T) {
		ve := Validate(User{
			ID:     "01234567890123456789012345678901234",
			Name:   "123",
			Age:    17,
			Email:  "@.ru",
			Role:   "stuf",
			Phones: []string{"8123456789"},
			meta:   json.RawMessage("Some meta"),
		})

		require.Equal(t, ve, ValidationErrors{
			ValidationError{Field: "ID", Err: ErrStringLen},
			ValidationError{Field: "Email", Err: ErrStringRegexp},
			ValidationError{Field: "Role", Err: ErrStringOutOfList},
			ValidationError{Field: "Phones", Err: ErrStringLen},
		})
	})

	t.Run("Failed validation with maximum values", func(t *testing.T) {
		ve := Validate(User{
			ID:     "0123456789012345678901234567890123456",
			Name:   "123",
			Age:    51,
			Email:  "test@mail.",
			Role:   "admin,stuff",
			Phones: []string{"812345678901"},
			meta:   json.RawMessage("Some meta"),
		})

		require.Equal(t, ve, ValidationErrors{
			ValidationError{Field: "ID", Err: ErrStringLen},
			ValidationError{Field: "Email", Err: ErrStringRegexp},
			ValidationError{Field: "Role", Err: ErrStringOutOfList},
			ValidationError{Field: "Phones", Err: ErrStringLen},
		})
	})
}

func TestValidateResponse(t *testing.T) {
	t.Run("Successful validation", func(t *testing.T) {
		require.Equal(t, Validate(Response{
			Code: 200,
			Body: "Some body",
		}), nil)
		require.Equal(t, Validate(Response{
			Code: 404,
			Body: "Some body",
		}), nil)
		require.Equal(t, Validate(Response{
			Code: 500,
			Body: "Some body",
		}), nil)
	})

	t.Run("Failed validation", func(t *testing.T) {
		require.Equal(t, Validate(Response{
			Code: 199,
			Body: "Some body",
		}), ValidationErrors{
			ValidationError{Field: "Code", Err: ErrIntOutOfList},
		})

		require.Equal(t, Validate(Response{
			Code: 201,
			Body: "Some body",
		}), ValidationErrors{
			ValidationError{Field: "Code", Err: ErrIntOutOfList},
		})

		require.Equal(t, Validate(Response{
			Code: 403,
			Body: "Some body",
		}), ValidationErrors{
			ValidationError{Field: "Code", Err: ErrIntOutOfList},
		})

		require.Equal(t, Validate(Response{
			Code: 405,
			Body: "Some body",
		}), ValidationErrors{
			ValidationError{Field: "Code", Err: ErrIntOutOfList},
		})

		require.Equal(t, Validate(Response{
			Code: 499,
			Body: "Some body",
		}), ValidationErrors{
			ValidationError{Field: "Code", Err: ErrIntOutOfList},
		})

		require.Equal(t, Validate(Response{
			Code: 501,
			Body: "Some body",
		}), ValidationErrors{
			ValidationError{Field: "Code", Err: ErrIntOutOfList},
		})
	})
}

func TestValidateApp(t *testing.T) {
	t.Run("Successful validation", func(t *testing.T) {
		ve := Validate(App{
			Version: "12345",
		})

		require.Equal(t, ve, nil)
	})

	t.Run("Failed validation with minimum values", func(t *testing.T) {
		ve := Validate(App{
			Version: "1234",
		})

		require.Equal(t, ve, ValidationErrors{
			ValidationError{Field: "Version", Err: ErrStringLen},
		})
	})

	t.Run("Failed validation with maximum values", func(t *testing.T) {
		ve := Validate(App{
			Version: "123456",
		})

		require.Equal(t, ve, ValidationErrors{
			ValidationError{Field: "Version", Err: ErrStringLen},
		})
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("Successful validation", func(t *testing.T) {
		ve := Validate(Token{
			Header:    []byte("Some Header"),
			Payload:   []byte("Some Payload"),
			Signature: []byte("Some Signature"),
		})

		require.Equal(t, ve, nil)
	})
}
