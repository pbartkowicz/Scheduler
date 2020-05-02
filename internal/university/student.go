package university

import (
	"errors"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	// ErrWrongPriority is returned when priorities for one of a subjects does not start from 1.
	ErrWrongPriority = errors.New("incorrect priority value: priorities have to start from 1")
	// ErrWrongSubPriority is returned when priorities for one of a subjects are not consecutive with repetition.
	ErrWrongSubPriority = errors.New("incorrect priority for subject: priorities have to be consecutive with repetition")
)

// StudentError represents an error struct returned when creating new Student.
type StudentError struct {
	Err error
}

func (e *StudentError) Error() string {
	return "failed to create student: " + e.Err.Error()
}

// Student represents a university student and his or her preferences.
// Name - student name which is read from file name with preferences.
// Priority - if set to true, than student will receive the same schedule as in ChosenGroups.
// Happieness -
// Preferences - contains list of groups with priorities for each subject, one priority for group.
// - Priorities have to start from 1 (highest).
// - Priorities have to be consecutive and they can be repeated.
// FinalGroups - groups to which student is assigned after scheduling.
type Student struct {
	Name        string
	Priority    bool
	Happieness  float64
	Preferences map[SubjectGroup]int
	FinalGroups map[string]*Group
}

// SubjectGroup is used as a key in Preferences.
// It contains subject and group name.
type SubjectGroup struct {
	Subject string
	Group   string
}

// NewStudent creates new instance of Student.
// It returns StudentError when passed parameters are invalid.
// Passed parameters:
// n - filename which contains student name
// pref:
// 0 - subject name
// 1 - group name
// 2 - group priority
func NewStudent(pref [][]string, n string) (*Student, error) {
	s := &Student{
		Name:        strings.Replace(n, ".xlsx", "", -1),
		Happieness:  100,
		Preferences: make(map[SubjectGroup]int),
		FinalGroups: make(map[string]*Group),
	}
	for _, p := range pref {
		pr, err := strconv.Atoi(p[2])
		if err != nil {
			return nil, &StudentError{Err: err}
		}
		s.Preferences[SubjectGroup{p[0], p[1]}] = pr
	}
	return s, s.validate()
}

// validate is used to check if priorities for each subject are set correctly.
func (s *Student) validate() error {
	sub := make(map[string][]int)
	for k, v := range s.Preferences {
		sub[k.Subject] = append(sub[k.Subject], v)
	}
	for _, v := range sub {
		sort.Ints(v)
	}
	for _, v := range sub {
		if v[0] != 1 {
			return &StudentError{Err: ErrWrongPriority}
		}
		for i := 1; i < v[len(v)-1]; i++ {
			diff := v[i] - v[i-1]
			if diff > 1 {
				return &StudentError{Err: ErrWrongSubPriority}
			}
		}
	}
	return nil
}

func (s *Student) GetPrefredGroup(subject string, groups []string) (res string) {
	p := math.MaxInt64
	for _, g := range groups {
		v := s.Preferences[SubjectGroup{subject, g}]
		if p > v {
			p = v
			res = g
		}
	}
	return
}
