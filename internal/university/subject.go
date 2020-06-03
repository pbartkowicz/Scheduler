package university

// Subject represets one subject.
// It contains all subject's groups.
// It implements sort.Interface based on number of conflicts in groups list.
type Subject struct {
	Name     string
	Lectures []*Group
	Groups   []*Group
}

func (s *Subject) Len() int {
	return len(s.Groups)
}

func (s *Subject) Less(i, j int) bool {
	return s.Groups[i].Conflicts() > s.Groups[j].Conflicts()
}

func (s *Subject) Swap(i, j int) {
	s.Groups[i], s.Groups[j] = s.Groups[j], s.Groups[i]
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
		if g.Name == gn {
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
