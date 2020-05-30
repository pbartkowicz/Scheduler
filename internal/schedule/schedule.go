package schedule

import (
	"github.com/pbartkowicz/scheduler/internal/university"
)

// Enroll ... - move to university?
func Enroll(schedule *university.Schedule, students []*university.Student) error {
	// Assign students to preferred groups
	for _, st := range students {
		for _, s := range schedule.Subjects {
			for _, l := range s.Lectures {
				l.Students = append(l.Students, st)
			}
			gns := s.GetGroupsNames()
			// Subject has no groups
			if len(gns) == 0 {
				continue
			}
			prefGroup := st.GetPrefredGroup(s.Name, gns)
			g := s.GetGroup(prefGroup)
			if st.Priority {
				g.PriorityStudents = append(g.PriorityStudents, st)
			} else {
				g.Students = append(g.Students, st)
			}
		}
	}
	return nil
}
