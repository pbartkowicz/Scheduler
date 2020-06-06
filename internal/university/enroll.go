package university

import (
	"fmt"
	"sort"
)

// Enroll is used to assign students and resolve conflicts in schedule
func (s *Schedule) Enroll(students []*Student) {
	s.assign(students)
	// Sort subjects by number of conflicts
	sort.Sort(s)
	s.resolve(students)
	// TODO: Check for conflicts in students plans
	printHappieness(students)
}

func printHappieness(students []*Student) {
	var happy float64
	var stLen int
	for _, st := range students {
		if !st.Priority {
			stLen++
			happy += st.GetHappieness()
		}
	}
	fmt.Printf("\nStudents happieness: %.2f\n", happy/float64(stLen))
}

// assign students to preferred groups
func (s *Schedule) assign(students []*Student) {
	for _, st := range students {
		for _, sub := range s.Subjects {
			for _, l := range sub.Lectures {
				l.Students = append(l.Students, st)
			}
			gns := sub.GetGroupsNames()
			// Subject has no groups
			if len(gns) == 0 {
				continue
			}
			prefGroup := st.GetPreferredGroup(sub.Name, gns)
			g := sub.GetGroup(prefGroup)
			if st.Priority {
				g.PriorityStudents = append(g.PriorityStudents, st)
			} else {
				g.Students = append(g.Students, st)
			}
			st.Happieness[sub.Name] = 100.0
		}
	}
}

func (s *Schedule) resolve(students []*Student) {
	// Sort groups by number of conflicts [descending]
	for _, sub := range s.Subjects {
		sort.Sort(sub)
	}
	for _, sub := range s.Subjects {
		if sub.Conflicts() == 0 {
			for _, st := range students {
				st.SetFinalGroup(sub)
			}
			continue
		}
		// Sort students within group by happieness [descending]
		for _, g := range sub.Groups {
			sort.Sort(g)
		}
		sg := &StudentGroup{}
		for i, g := range sub.Groups {
			c := g.Conflicts()
			if c <= 0 {
				continue
			}
			// Get students who likes other groups
			sgs := getStudents(i, true, sub, g.Students)
			// Get students who can be moved to other groups and don't like them
			mSgs := getStudents(i, false, sub, g.Students)

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
				sg.Student.CalculateHappieness(sub.Name)
			}
		}
		// Set final groups for this subject
		for _, st := range students {
			st.SetFinalGroup(sub)
		}
	}
}

// StudentGroup is used to store information about students who likes other groups and can be moved to them.
type StudentGroup struct {
	Student *Student
	Group   *Group
}

// getStudents returns students who can be moved to other groups and like or doesn't like being moved.
func getStudents(i int, likes bool, s *Subject, students []*Student) (sgs []*StudentGroup) {
	for _, st := range students {
		if likes && st.Likes(s.Name, s.Groups[i+1].Name) && st.CanMove(s.Name, s.Groups[i+1]) {
			sgs = append(sgs, &StudentGroup{Student: st, Group: s.Groups[i+1]})
		}
		if !likes && !st.Likes(s.Name, s.Groups[i+1].Name) && st.CanMove(s.Name, s.Groups[i+1]) {
			sgs = append(sgs, &StudentGroup{Student: st, Group: s.Groups[i+1]})
		}
	}
	return
}

// pop is used to remove first StudentGroup from a slice and return it.
func pop(sts []*StudentGroup) (*StudentGroup, []*StudentGroup) {
	first, sts := sts[0], sts[1:]
	return first, sts
}
