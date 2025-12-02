package errors

import (
	"errors"
	"testing"
)

func TestAtoError(t *testing.T) {
	err := &AtoError{
		Message: "test error",
	}

	if err.Error() != "test error" {
		t.Errorf("Expected 'test error', got '%s'", err.Error())
	}
}

func TestAtoErrorWithCause(t *testing.T) {
	cause := errors.New("underlying error")
	err := &AtoError{
		Message: "test error",
		Cause:   cause,
	}

	expected := "test error: underlying error"
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}

	if !errors.Is(err, cause) {
		t.Error("Expected error to wrap cause")
	}
}

func TestNewUserException(t *testing.T) {
	err := NewUserException("user error")
	if err.Message != "user error" {
		t.Errorf("Expected 'user error', got '%s'", err.Message)
	}
}

func TestNewUserBadParameterError(t *testing.T) {
	err := NewUserBadParameterError("bad param", "param1")
	if err.Message != "bad param" {
		t.Errorf("Expected 'bad param', got '%s'", err.Message)
	}
	if err.Parameter != "param1" {
		t.Errorf("Expected 'param1', got '%s'", err.Parameter)
	}
}

func TestNewUserNoProjectException(t *testing.T) {
	err := NewUserNoProjectException("no project")
	if err.Message != "no project" {
		t.Errorf("Expected 'no project', got '%s'", err.Message)
	}
}

func TestNewUserResourceException(t *testing.T) {
	err := NewUserResourceException("resource error")
	if err.Message != "resource error" {
		t.Errorf("Expected 'resource error', got '%s'", err.Message)
	}
}

func TestNewInternalException(t *testing.T) {
	err := NewInternalException("internal error")
	if err.Message != "internal error" {
		t.Errorf("Expected 'internal error', got '%s'", err.Message)
	}
}
