package permcheck

import (
	"errors"
	"testing"
)

func TestSimplePermission_Satisfies(t *testing.T) {
	fallbackErr := errors.New("Permission Error")
	perm := NewPermission("test", fallbackErr)

	if perm.Satisfies([]string{"test", "example"}) != nil {
		t.Errorf("Satisfies was incorrect, expected nil, got error")
	}

	if perm.Satisfies([]string{"example", "sample"}) != fallbackErr {
		t.Errorf("Satisfies was incorrect, expected fallback error, got nil or wrong error")
	}

	var nilPerm *simplePermission = nil
	if nilPerm.Satisfies([]string{"example", "sample"}) != nil {
		t.Errorf("Satisfies was incorrect for nil permission, expected nil, got error")
	}
}

func TestEmptyPermission_Satisfies(t *testing.T) {
	perm := Empty()

	if perm.Satisfies([]string{"any", "permission"}) != nil {
		t.Errorf("Satisfies was incorrect, expected nil, got error")
	}
}

func TestOrPermission_Satisfies(t *testing.T) {
	fallbackErr1 := errors.New("Permission Error1")
	fallbackErr2 := errors.New("Permission Error2")
	perm1 := NewPermission("test", fallbackErr1)
	perm2 := NewPermission("example", fallbackErr2)

	orPerm := Or(perm1, perm2)

	if orPerm.Satisfies([]string{"test", "sample"}) != nil {
		t.Errorf("Satisfies was incorrect, expected nil, got error")
	}

	if orPerm.Satisfies([]string{"sample", "anotherSample"}) == nil {
		t.Errorf("Satisfies was incorrect, expected error, got nil")
	}

	if orPerm.Satisfies([]string{}) == nil {
		t.Errorf("Satisfies was incorrect, expected error, got nil")
	}
}

func TestAndPermission_Satisfies(t *testing.T) {
	fallbackErr1 := errors.New("Permission Error1")
	fallbackErr2 := errors.New("Permission Error2")
	perm1 := NewPermission("test", fallbackErr1)
	perm2 := NewPermission("example", fallbackErr2)

	andPerm := And(perm1, perm2)

	if andPerm.Satisfies([]string{"test", "example"}) != nil {
		t.Errorf("Satisfies was incorrect, expected nil, got error")
	}

	if andPerm.Satisfies([]string{"test", "sample"}) == nil {
		t.Errorf("Satisfies was incorrect, expected error, got nil")
	}

	if andPerm.Satisfies([]string{"sample", "example"}) == nil {
		t.Errorf("Satisfies was incorrect, expected nil, got error")
	}

	if andPerm.Satisfies([]string{"sample", "anotherSample"}) == nil {
		t.Errorf("Satisfies was incorrect, expected error, got nil")
	}
}
