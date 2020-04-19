package schedule

// Student represents a student.
// Happieness
type Student struct {
	Name       string
	Priority   bool
	Happieness float64
	Groups     []ChosenGroups
}

// ChosenGroups ...
type ChosenGroups struct {
	Subject         string
	GroupPriorities []GroupPriority
}

// GroupPriority ...
type GroupPriority struct {
	Group    string
	Priority int
}
