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
		if g.Name == n {
			return g
		}
	}
	return nil
}
