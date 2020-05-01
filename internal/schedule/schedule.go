// Package schedule ...
package schedule

const (
	timeLayout = "15:04"
	dateLayout = "01-02-06"
)

// Schedule represents a students' schedule for one semester.
type Schedule struct {
	Subjects []Subject
}
