package schedule

import (
	"fmt"
	"sort"

	"github.com/pbartkowicz/scheduler/internal/university"
)

// Enroll ... - move to university?
func Enroll(schedule *university.Schedule, students []*university.Student) {
	assign(schedule, students)
	// Sort subjects by number of conflicts
	sort.Sort(schedule)
	for _, s := range schedule.Subjects {
		fmt.Printf("\n%+v\n", s.Name)
		for _, g := range s.Groups {
			fmt.Printf("%+v: %+v\n", g.Name, len(g.Students)+len(g.PriorityStudents))
		}
	}
	resolve(schedule, students)

	for _, s := range schedule.Subjects {
		fmt.Printf("\n%+v\n", s.Name)
		for _, g := range s.Groups {
			fmt.Printf("%+v: %+v\n", g.Name, len(g.Students)+len(g.PriorityStudents))
		}
	}

	for _, st := range students {
		fmt.Printf("\n%s: %+v\n", st.Name, st.GetHappieness())
	}
}

// assign students to preferred groups
func assign(schedule *university.Schedule, students []*university.Student) {
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
			prefGroup := st.GetPreferredGroup(s.Name, gns)
			g := s.GetGroup(prefGroup)
			if st.Priority {
				g.PriorityStudents = append(g.PriorityStudents, st)
			} else {
				g.Students = append(g.Students, st)
			}
			st.Happieness[s.Name] = 100.0
		}
	}
}

func resolve(schedule *university.Schedule, students []*university.Student) {
	// Sort groups by number of conflicts [descending]
	for _, s := range schedule.Subjects {
		sort.Sort(s)
	}
	for _, s := range schedule.Subjects {
		if s.Conflicts() == 0 {
			for _, st := range students {
				st.SetFinalGroup(s)
			}
			continue
		}
		// Sort students within group by happieness [descending]
		for _, g := range s.Groups {
			sort.Sort(g)
		}
		sg := &StudentGroup{}
		for i, g := range s.Groups {
			c := g.Conflicts()
			if c <= 0 {
				continue
			}
			fmt.Printf("\nConflicts:%v", s.Conflicts())
			// Get students who likes other groups
			sgs := getStudents(i, true, s, g.Students)
			// Get students who can be moved to other groups and don't like them
			mSgs := getStudents(i, false, s, g.Students)

			fmt.Printf("\n%s:%s:%v:%v:%v\n", s.Name, g.Name, len(sgs), len(mSgs), c)

			for ; c > 0; c-- {
				// Move students who like other groups and can be moved
				if len(sgs) != 0 {
					sg, sgs = pop(sgs)
					sg.Group.Students = append(sg.Group.Students, sg.Student)
					g.RemoveStudent(sg.Student)
					continue
				}

				sg, mSgs = pop(mSgs)
				sg.Group.Students = append(sg.Group.Students, sg.Student)
				g.RemoveStudent(sg.Student)
				// Change student happieness
				sg.Student.CalculateHappieness(s.Name)
			}
		}
		// Set final groups for this subject
		for _, st := range students {
			st.SetFinalGroup(s)
		}
	}
	// TODO: Check for conflicts in students plans
}

// StudentGroup is used to store information about students who likes other groups and can be moved to them.
type StudentGroup struct {
	Student *university.Student
	Group   *university.Group
}

// getStudents returns students who can be moved to other groups and like or doesn't like being moved.
// TODO: Refactor it to one
func getStudents(i int, likes bool, s *university.Subject, students []*university.Student) (sgs []*StudentGroup) {
	for _, st := range students {
		if likes {
			if st.Likes(s.Name, s.Groups[i+1].Name) && st.CanMove(s.Name, s.Groups[i+1]) {
				sgs = append(sgs, &StudentGroup{Student: st, Group: s.Groups[i+1]})
			}
		} else {
			if !st.Likes(s.Name, s.Groups[i+1].Name) && st.CanMove(s.Name, s.Groups[i+1]) {
				sgs = append(sgs, &StudentGroup{Student: st, Group: s.Groups[i+1]})
			}
		}
	}
	return
}

// pop is used to remove first StudentGroup from a slice and return it.
func pop(sts []*StudentGroup) (*StudentGroup, []*StudentGroup) {
	first, sts := sts[0], sts[1:]
	return first, sts
}
