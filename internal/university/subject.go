package university

// Subject represets one subject.
// It contains all subject's groups.
type Subject struct {
	Name     string
	Lectures []*Group
	Groups   []*Group
}

func (s *Subject) GetGroupsNames() []string {
	var gn []string
	for _, g := range s.Groups {
		gn = append(gn, g.Name)
	}
	return gn
}

func (s *Subject) GetGroup(n string) *Group {
	for _, g := range s.Groups {
		// Group capacity is set to -1 when group has more than one term.
		if g.Name == n && g.Capacity != -1 {
			return g
		}
	}
	return nil
}

// Conflicts is used to calculate the number of conflicts within one subject.
func (s *Subject) Conflicts() (res int) {
	for _, g := range s.Groups {
		c := g.Conflicts()
		if c < 0 {
			res += -c
		}
	}
	return
}
