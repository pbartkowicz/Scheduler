package university

import (
	"errors"
	"fmt"
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
	Name string
	Err  error
}

func (e *StudentError) Error() string {
	return fmt.Sprintf("failed to create student [%s]: %s", e.Name, e.Err.Error())
}

// Student represents a university student and their preferences.
// Name - student name which is read from file name with preferences.
// Priority - if set to true, than student will receive the same schedule as in ChosenGroups.
// Happieness - reflects how much final schedule is similar to their preferences.
// It contains map in which a key is subject name and value is calculated happieness.
// Preferences - contains list of groups with priorities for each subject, one priority for group.
// - Priorities have to start from 1 (highest).
// - Priorities have to be consecutive and they can be repeated.
// FinalGroups - groups to which student is assigned after scheduling.
type Student struct {
	Name        string
	Priority    bool
	Preferences map[SubjectGroup]int
	Happieness  map[string]float64
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
		Preferences: make(map[SubjectGroup]int),
		Happieness:  make(map[string]float64),
		FinalGroups: make(map[string]*Group),
	}
	for _, p := range pref {
		pr, err := strconv.Atoi(p[2])
		if err != nil {
			return nil, &StudentError{Err: err, Name: s.Name}
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
			return &StudentError{Err: ErrWrongPriority, Name: s.Name}
		}
		for i := 1; i < v[len(v)-1]; i++ {
			diff := v[i] - v[i-1]
			if diff > 1 {
				return &StudentError{Err: ErrWrongSubPriority, Name: s.Name}
			}
		}
	}
	return nil
}

// GetPreferredGroup returns a name of the group to which student wants to be assigned the most.
func (s *Student) GetPreferredGroup(subject string, groups []string) (res string) {
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

// SetFinalGroup sets a group to which student was assigned.
func (s *Student) SetFinalGroup(sub *Subject) {
	s.FinalGroups[sub.Name] = sub.GetStudentGroup(s.Name)
}

// Likes checks if a student set the highest priority to a passed group.
func (s *Student) Likes(sub, gn string) bool {
	return s.Preferences[SubjectGroup{sub, gn}] == 1
}

// CanMove checks if a student can be moved to the other group.
func (s *Student) CanMove(sub string, g *Group) bool {
	for _, fg := range s.FinalGroups {
		if fg != nil && g.Collide(fg) {
			return false
		}
	}
	return true
}

// GetHappieness is used to retrieve student's happieness
func (s *Student) GetHappieness() (res float64) {
	for _, v := range s.Happieness {
		res += v
	}
	return res / float64(len(s.Happieness))
}

// CalculateHappieness is used to count student's happieness for a subject.
// It's based on their preferences.
func (s *Student) CalculateHappieness(sn string) {
	// Count number of distinct priorities within one subject
	var priorities []int
	for k, v := range s.Preferences {
		if k.Subject == sn {
			priorities = append(priorities, v)
		}
	}
	encountered := map[int]bool{}
	distinct := []int{}
	for _, p := range priorities {
		if !encountered[p] {
			encountered[p] = true
			distinct = append(distinct, p)
		}
	}

	if len(distinct) == 1 {
		s.Happieness[sn] = 100.0
		return
	}
	s.Happieness[sn] = (1.0 / (float64(len(distinct)))) * 100.0
}

// Save creates a slice with groups which were chosen for a student.
func (s *Student) Save() [][]string {
	var i int
	var gLen int
	for _, v := range s.FinalGroups {
		if v != nil {
			gLen++
		}
	}
	res := make([][]string, gLen)
	for k, v := range s.FinalGroups {
		if v == nil {
			continue
		}
		r := make([]string, 2)
		r[0] = k
		r[1] = v.Name
		res[i] = r
		i++
	}
	return res
}
