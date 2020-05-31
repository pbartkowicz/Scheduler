package university

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortSubject(t *testing.T) {
	tests := []struct {
		name string
		sub  *Subject
		want *Subject
	}{
		{
			name: "Successfully sort groups by number of conflicts",
			sub: &Subject{
				Name: "Subject1",
				Groups: []*Group{
					{
						Name:     "1",
						Capacity: 2,
						Students: []*Student{},
					},
					{
						Name:     "2",
						Capacity: 1,
						Students: []*Student{
							{
								Name: "e",
							},
						},
					},
					{
						Name:     "3",
						Capacity: 2,
						Students: []*Student{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
							{
								Name: "d",
							},
						},
					},
				},
			},
			want: &Subject{
				Name: "Subject1",
				Groups: []*Group{
					{
						Name:     "3",
						Capacity: 2,
						Students: []*Student{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
							{
								Name: "d",
							},
						},
					},
					{
						Name:     "2",
						Capacity: 1,
						Students: []*Student{
							{
								Name: "e",
							},
						},
					},
					{
						Name:     "1",
						Capacity: 2,
						Students: []*Student{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.sub)
			if !cmp.Equal(tt.sub, tt.want) {
				t.Errorf("sort.Sort(Subject) got = %v, want %v", tt.sub, tt.want)
			}
		})
	}
}

func TestSubject_GetGroupsNames(t *testing.T) {
	tests := []struct {
		name string
		s    *Subject
		want []string
	}{
		{
			name: "Successfully retrieves groups names",
			s: &Subject{
				Groups: []*Group{
					{
						Name:     "1",
						Capacity: 10,
					},
					{
						Name:     "1",
						Capacity: -1,
					},
					{
						Name:     "2",
						Capacity: 3,
					},
					{
						Name:     "4",
						Capacity: -1,
					},
				},
			},
			want: []string{
				"1",
				"2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GetGroupsNames()
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Subject.GetGroupsNames() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_GetGroup(t *testing.T) {
	type args struct {
		gn string
	}
	tests := []struct {
		name string
		args args
		s    *Subject
		want *Group
	}{
		{
			name: "Returns nil because group doesn't exist",
			s:    &Subject{},
		},
		{
			name: "Successfully returns group",
			args: args{
				gn: "1",
			},
			s: &Subject{
				Groups: []*Group{
					{
						Name: "1a",
					},
					{
						Name: "1",
					},
				},
			},
			want: &Group{
				Name: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GetGroup(tt.args.gn)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Subject.GetGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_Conflicts(t *testing.T) {
	tests := []struct {
		name string
		s    *Subject
		want int
	}{
		{
			name: "Successfully returns number of conflicts",
			s: &Subject{
				Groups: []*Group{
					{
						Name:     "1",
						Capacity: 2,
						Students: []*Student{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
							{
								Name: "d",
							},
						},
					},
					{
						Name:     "2",
						Capacity: 2,
						Students: []*Student{},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Conflicts()
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Subject.Conflicts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_GetStudentGroup(t *testing.T) {
	type args struct {
		sn string
	}
	tests := []struct {
		name string
		args args
		s    *Subject
		want *Group
	}{
		{
			name: "Successfully returns group for priority student",
			args: args{
				sn: "student",
			},
			s: &Subject{
				Groups: []*Group{
					{
						Name: "1",
						PriorityStudents: []*Student{
							{
								Name: "student",
							},
							{
								Name: "student2",
							},
						},
					},
					{
						Name: "2",
					},
				},
			},
			want: &Group{
				Name: "1",
				PriorityStudents: []*Student{
					{
						Name: "student",
					},
					{
						Name: "student2",
					},
				},
			},
		},
		{
			name: "Returns nil because student is not assigned",
			args: args{
				sn: "student",
			},
			s: &Subject{},
		},
		{
			name: "Successfully returns group for student",
			args: args{
				sn: "student",
			},
			s: &Subject{
				Groups: []*Group{
					{
						Name: "1",
						Students: []*Student{
							{
								Name: "student",
							},
							{
								Name: "student2",
							},
						},
					},
					{
						Name: "2",
					},
				},
			},
			want: &Group{
				Name: "1",
				Students: []*Student{
					{
						Name: "student",
					},
					{
						Name: "student2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GetStudentGroup(tt.args.sn)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Subject.GetStudentGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}
