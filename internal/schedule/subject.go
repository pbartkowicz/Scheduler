package schedule

// Subject represets one subject.
// It contains all subject's groups.
type Subject struct {
	Name     string
	Lectures []Group
	Groups   []Group
}
