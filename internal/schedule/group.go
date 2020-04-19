package schedule

import (
	"fmt"
	"strconv"
	"time"
)

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

// Group represents a
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

// NewGroup ...
func NewGroup(v []string) *Group {
	// TODO errors and validation
	f, _ := strconv.Atoi(v[8])
	c, _ := strconv.Atoi(v[10])
	dLayout := "01-02-06"
	d, err := time.Parse(dLayout, v[7])
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	tLayout := "15:04"
	st, _ := time.Parse(tLayout, v[4])
	et, _ := time.Parse(tLayout, v[5])
	return &Group{
		Type:      types[v[1]],
		Teacher:   v[2],
		Weekday:   time.Weekday(weekdays[v[3]]),
		StartTime: st,
		EndTime:   et,
		Place:     v[6],
		StartDate: d,
		Frequency: f,
		Name:      v[9],
		Capacity:  c,
	}
}
