// Package university ...
package university

const (
	timeLayout = "15:04"
	dateLayout = "01-02-06"
)

// Schedule represents schedule for one semester.
type Schedule struct {
	Subjects []*Subject
}

// NewSchedule creates new instance of Schedule.
// It returns GroupError when passed parameters are invalid.
// It receives slice of groups - see NewGroup for description of parameters.
func NewSchedule(groups [][]string) (*Schedule, error) {
	s := &Schedule{}
	for _, g := range groups {
		ng, err := NewGroup(g)
		if err != nil {
			return nil, err
		}
		sub := s.GetSubject(g[0])
		if sub == nil {
			sub = &Subject{
				Name: g[0],
			}
			s.Subjects = append(s.Subjects, sub)
		}
		if ng.Type == Lecture {
			sub.Lectures = append(sub.Lectures, ng)
		} else {
			sub.Groups = append(sub.Groups, ng)
		}
	}
	return s, nil
}

// GetSubject returns a Subject with a passed name.
func (s *Schedule) GetSubject(n string) *Subject {
	for _, sub := range s.Subjects {
		if sub.Name == n {
			return sub
		}
	}
	return nil
}
