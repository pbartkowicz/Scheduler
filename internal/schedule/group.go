package schedule

import "time"

// ClassType defines
type ClassType string

const (
	// Class ...
	Class ClassType = "Class"
	// Lecture ...
	Lecture ClassType = "Lecture"
	// Laboratory ...
	Laboratory ClassType = "Laboratory"
)

// Group represents a
type Group struct {
	Name             string
	Place            string
	Teacher          string
	Type             ClassType
	Capacity         int
	Frequency        int
	StartTime        time.Time
	EndTime          time.Time
	Weekday          time.Weekday
	StartDate        time.Time
	Students         []Student
	PriorityStudents []Student
}
