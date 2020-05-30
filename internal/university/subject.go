package university

// Subject represets one subject.
// It contains all subject's groups.
type Subject struct {
	Name     string
	Lectures []*Group
	Groups   []*Group
}

// GetGroupsNames returns names of all groups within a subject.
func (s *Subject) GetGroupsNames() []string {
	var gn []string
	for _, g := range s.Groups {
		gn = append(gn, g.Name)
	}
	return gn
}

// GetGroup returns a group which name is the same as passed.
// It returns nil if a group was not found.
func (s *Subject) GetGroup(gn string) *Group {
	for _, g := range s.Groups {
		// Group capacity is set to -1 when group has more than one term.
		if g.Name == gn && g.Capacity != -1 {
			return g
		}
	}
	return nil
}

// Conflicts is used to calculate the number of conflicts within one subject.
func (s *Subject) Conflicts() (res int) {
	for _, g := range s.Groups {
		c := g.Conflicts()
		if c > 0 {
			res += c
		}
	}
	return
}

// GetStudentGroup returns a group to which a student was assigned.
// It receives student's name.
func (s *Subject) GetStudentGroup(sn string) *Group {
	for _, g := range s.Groups {
		for _, st := range g.PriorityStudents {
			if st.Name == sn {
				return g
			}
		}
		for _, st := range g.Students {
			if st.Name == sn {
				return g
			}
		}
	}
	return nil
}
