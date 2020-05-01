// Package university ...
package university

const (
	timeLayout = "15:04"
	dateLayout = "01-02-06"
)

// Schedule represents a students' schedule for one semester.
type Schedule struct {
	Subjects []*Subject
}

func NewSchedule() *Schedule {
	return nil
}

func (s *Schedule) GetSubject(n string) *Subject {
	for _, sub := range s.Subjects {
		if sub.Name == n {
			return sub
		}
	}
	return nil
}
