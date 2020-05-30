package schedule

import (
	"fmt"
	"sort"

	"github.com/pbartkowicz/scheduler/internal/university"
)

// Enroll ... - move to university?
func Enroll(schedule *university.Schedule, students []*university.Student) {
	assign(schedule, students)
	for _, s := range schedule.Subjects {
		fmt.Printf("\n%+v\n", s.Name)
		for _, g := range s.Groups {
			fmt.Printf("%+v: %+v\n", g.Name, len(g.Students)+len(g.PriorityStudents))
		}
	}
	// Sort subjects by number of conflicts
	sort.Sort(schedule)
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
	for _, s := range schedule.Subjects {
		if s.Conflicts() == 0 {
			for _, st := range students {
				st.SetFinalGroup(s)
			}
			continue
		}
		// Sort groups by number of conflicts [descending]
		sort.Sort(s)
		sg := &StudentGroup{}
		for i, g := range s.Groups {
			if g.Capacity == -1 {
				continue
			}
			// Sort students within group by happieness [ascending]
			sort.Sort(g)
			// Get students who likes other groups
			sgs := getStudents(i, s, g.Students)
			// Get students who can be moved to other groups
			mSgs := getStudentsToMove(i, s, g.Students)
			c := g.Conflicts()
			if c <= 0 {
				continue
			}
			fmt.Printf("\n%s:%s:%v:%v:%v", s.Name, g.Name, len(sgs), len(mSgs), c)
			for ; c > 0; c-- {
				// Move students who like other groups and can be moved
				if len(sgs) != 0 {
					sg, sgs = pop(sgs)
					sg.group.Students = append(sg.group.Students, sg.student)
					g.RemoveStudent(sg.student)
					continue
				}

				sg, mSgs = pop(mSgs)
				sg.group.Students = append(sg.group.Students, sg.student)
				g.RemoveStudent(sg.student)
				// Change student happieness
				sg.student.CalculateHappieness(s.Name)
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
	student *university.Student
	group   *university.Group
}

// getStudents returns students who like other groups and can be moved to them.
func getStudents(i int, s *university.Subject, students []*university.Student) (sgs []*StudentGroup) {
	for ; i < len(s.Groups); i++ {
		for _, st := range students {
			if s.Groups[i].Capacity != -1 && st.Likes(s.Name, s.Groups[i].Name) && st.CanMove(s.Name, s.Groups[i]) {
				sgs = append(sgs, &StudentGroup{student: st, group: s.Groups[i]})
			}
		}
	}
	return
}

// getStudentsToMove returns students who can be moved to other groups.
func getStudentsToMove(i int, s *university.Subject, students []*university.Student) (sgs []*StudentGroup) {
	for ; i < len(s.Groups); i++ {
		for _, st := range students {
			if s.Groups[i].Capacity != -1 && !st.Likes(s.Name, s.Groups[i].Name) && st.CanMove(s.Name, s.Groups[i]) {
				sgs = append(sgs, &StudentGroup{student: st, group: s.Groups[i]})
			}
		}
	}
	return
}

// pop is used to remove first StudentGroup from a slice and return it.
func pop(sts []*StudentGroup) (*StudentGroup, []*StudentGroup) {
	first := sts[0]
	sts[len(sts)-1], sts[0] = sts[0], sts[len(sts)-1]
	return first, sts[:len(sts)-1]
}
