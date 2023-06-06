// Package permcheck provides a simple way to manage and verify permissions.
package permcheck

import (
	"fmt"
)

// Permission is an interface that describes a permission which can be
// satisfied by a set of strings.
type Permission interface {
	// Satisfies checks if the permission can be satisfied by the given
	// permissions and returns an error if it can't be satisfied.
	Satisfies([]string) error
}

// NewPermission creates a new simplePermission with a given value and a fallback error.
// This function returns a Permission that can be checked for satisfaction.
func NewPermission(p string, fallbackErr error) Permission {
	return &simplePermission{
		value:       p,
		fallbackErr: fallbackErr,
	}
}

// simplePermission is a struct implementing the Permission interface.
// It carries a value which will be used to check against a set of permissions.
type simplePermission struct {
	value       string
	fallbackErr error
}

// Satisfies checks if the simplePermission is satisfied by any permission
// in the given permission set.
func (p *simplePermission) Satisfies(perms []string) error {
	if p == nil {
		return nil
	}

	for _, perm := range perms {
		if perm == p.value {
			return nil
		}
	}

	return p.fallbackErr
}

// Empty creates an emptyPermission.
// This function returns a Permission that is always satisfied.
func Empty() Permission {
	return &emptyPermission{}
}

// emptyPermission is a struct implementing the Permission interface.
// It is always satisfied.
type emptyPermission struct {
	value string
}

// Satisfies checks if the emptyPermission is satisfied by any permission
// in the given permission set. As it's always satisfied, it always returns nil.
func (p *emptyPermission) Satisfies(perms []string) error {
	return nil
}

// Or creates a new orPermission with two underlying permissions.
// This function returns a Permission that is satisfied if either of the underlying permissions is satisfied.
func Or(s1, s2 Permission) Permission {
	return &orPermission{
		s1: s1,
		s2: s2,
	}
}

// orPermission is a struct implementing the Permission interface.
// It carries two Permissions and is satisfied if either of them is satisfied.
type orPermission struct {
	s1 Permission
	s2 Permission
}

// Satisfies checks if the orPermission is satisfied by any permission
// in the given permission set.
func (p *orPermission) Satisfies(perms []string) error {
	var err1, err2 error

	err1 = p.s1.Satisfies(perms)
	err2 = p.s2.Satisfies(perms)

	if err1 != nil && err2 != nil {
		return fmt.Errorf("%w or %w", err1, err2)
	}

	return nil
}

// And creates a new andPermission with two underlying permissions.
// This function returns a Permission that is satisfied only if both underlying permissions are satisfied.
func And(s1, s2 Permission) Permission {
	return &andPermission{
		s1: s1,
		s2: s2,
	}
}

// andPermission is a struct implementing the Permission interface.
// It carries two Permissions and is satisfied if both of them are satisfied.
type andPermission struct {
	s1 Permission
	s2 Permission
}

// Satisfies checks if the andPermission is satisfied by any permission
// in the given permission set.
func (p *andPermission) Satisfies(perms []string) error {
	var err1, err2 error

	err1 = p.s1.Satisfies(perms)
	err2 = p.s2.Satisfies(perms)

	if err1 == nil && err2 == nil {
		return nil
	}

	if err1 != nil && err2 != nil {
		return fmt.Errorf("%w and %w", err1, err2)
	}

	if err1 != nil {
		return err1
	}

	return err2
}
