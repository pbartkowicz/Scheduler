package university

import (
	"errors"
	"strconv"
	"time"
)

var (
	// ErrWrongClassType is returned when a passed class type is incorrect.
	ErrWrongClassType = errors.New("incorrect class type, available types: C - Class, W - Lecture, L - Laboratory")
	// ErrWrongWeekday is returned when a passed weekday is incorrect.
	ErrWrongWeekday = errors.New("incorrect weekday, available weekdays: Monday, Tuesday, Wednesday, Thursday, Friday")
)

// GroupError represents an error struct returned when creating new Group.
type GroupError struct {
	Err error
}

func (e *GroupError) Error() string {
	return "failed to create group: " + e.Err.Error()
}

// ClassType defines type of a class.
type ClassType string

const (
	// Class - class.
	Class ClassType = "Class"
	// Lecture - lecture.
	Lecture ClassType = "Lecture"
	// Laboratory - laboratory.
	Laboratory ClassType = "Laboratory"
)

var types = map[string]ClassType{
	"C": Class,
	"W": Lecture,
	"L": Laboratory,
}

var weekdays = map[string]time.Weekday{
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
}

// Group represents a single students group for one subject.
type Group struct {
	Type             ClassType
	Teacher          string
	Weekday          time.Weekday
	StartTime        time.Time
	EndTime          time.Time
	Place            string
	StartDate        time.Time
	Frequency        int
	Name             string
	Capacity         int
	Students         []Student
	PriorityStudents []Student
}

// NewGroup creates new instance of Group.
// It returns GroupError when passed parameters are invalid.
// subjects:
// 0 - subject name
// 1 - class type [C - class, W - lecture, L - laboratory]
// 2 - teacher
// 3 - weekday [Monday, Tuesday, Wednesday, Thursday, Friday]
// 4 - start time, format: 14:00
// 5 - end time, format: 15:30
// 6 - place
// 7 - start date, format: 03-05-20 (5th of March 2020)
// 8 - frequency, format: number
// 9 - group name
// 10 - capacity, format: nubmer
func NewGroup(subjects []string) (*Group, error) {
	t := types[subjects[1]]
	if t == "" {
		return nil, &GroupError{Err: ErrWrongClassType}
	}

	w := weekdays[subjects[3]]
	if w == 0 {
		return nil, &GroupError{Err: ErrWrongWeekday}
	}

	st, err := time.Parse(timeLayout, subjects[4])
	if err != nil {
		return nil, &GroupError{Err: err}
	}
	et, err := time.Parse(timeLayout, subjects[5])
	if err != nil {
		return nil, &GroupError{Err: err}
	}

	d, err := time.Parse(dateLayout, subjects[7])
	if err != nil {
		return nil, &GroupError{Err: err}
	}

	f, err := strconv.Atoi(subjects[8])
	if err != nil {
		return nil, &GroupError{Err: err}
	}
	c, err := strconv.Atoi(subjects[10])
	if err != nil {
		return nil, &GroupError{Err: err}
	}

	return &Group{
		Type:      t,
		Teacher:   subjects[2],
		Weekday:   time.Weekday(w),
		StartTime: st,
		EndTime:   et,
		Place:     subjects[6],
		StartDate: d,
		Frequency: f,
		Name:      subjects[9],
		Capacity:  c,
	}, nil
}
