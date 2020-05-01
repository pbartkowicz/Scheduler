// Package tools provides mechanisms used in testing.
package tools

import "reflect"

// CompareErrors is used in tests to compare actual error with expected error.
func CompareErrors(a error, e error) bool {
	ta := reflect.TypeOf(a)
	te := reflect.TypeOf(e)
	if ta != te {
		return false
	}
	// Types are equal, but errors can be nil.
	if ta == nil {
		return true
	}
	return a.Error() == e.Error()
}
