package schedule

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pbartkowicz/scheduler/internal/university"
)

func Test_pop(t *testing.T) {
	type args struct {
		sts []*StudentGroup
	}
	tests := []struct {
		name      string
		args      args
		want      *StudentGroup
		wantSlice []*StudentGroup
	}{
		{
			name: "Successfully pops StudentGroup",
			args: args{
				sts: []*StudentGroup{
					{
						Student: &university.Student{
							Name: "a",
						},
					},
					{
						Student: &university.Student{
							Name: "b",
						},
					},
					{
						Student: &university.Student{
							Name: "c",
						},
					},
				},
			},
			want: &StudentGroup{
				Student: &university.Student{
					Name: "a",
				},
			},
			wantSlice: []*StudentGroup{
				{
					Student: &university.Student{
						Name: "b",
					},
				},
				{
					Student: &university.Student{
						Name: "c",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotSlice := pop(tt.args.sts)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("pop() got = %v, want %v", got, tt.want)
			}
			if !cmp.Equal(gotSlice, tt.wantSlice) {
				t.Errorf("pop() got = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}

func Test_getStudents(t *testing.T) {
	type args struct {
		i        int
		likes    bool
		s        *university.Subject
		students []*university.Student
	}
	tests := []struct {
		name    string
		args    args
		wantSgs []*StudentGroup
	}{
		{
			name: "Return students who like other groups",
			args: args{
				i:     0,
				likes: true,
				s: &university.Subject{
					Name: "Math",
					Groups: []*university.Group{
						{
							Capacity: 5,
							Name:     "1",
						},
						{
							Capacity: -1,
							Name:     "1",
						},
						{
							Capacity: 5,
							Name:     "2",
						},
						{
							Capacity: 5,
							Name:     "3",
						},
					},
				},
				students: []*university.Student{
					{
						Name: "a",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 1,
							{
								Subject: "Math",
								Group:   "3",
							}: 1,
						},
					},
					{
						Name: "b",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 2,
							{
								Subject: "Math",
								Group:   "3",
							}: 3,
						},
					},
					{
						Name: "c",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 2,
							{
								Subject: "Math",
								Group:   "3",
							}: 1,
						},
					},
				},
			},
			wantSgs: []*StudentGroup{
				{
					Student: &university.Student{
						Name: "a",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 1,
							{
								Subject: "Math",
								Group:   "3",
							}: 1,
						},
					},
					Group: &university.Group{
						Capacity: 5,
						Name:     "2",
					},
				},
				{
					Student: &university.Student{
						Name: "a",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 1,
							{
								Subject: "Math",
								Group:   "3",
							}: 1,
						},
					},
					Group: &university.Group{
						Capacity: 5,
						Name:     "3",
					},
				},
				{
					Student: &university.Student{
						Name: "c",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 2,
							{
								Subject: "Math",
								Group:   "3",
							}: 1,
						},
					},
					Group: &university.Group{
						Capacity: 5,
						Name:     "3",
					},
				},
			},
		},
		{
			name: "Return students who can be moved to other groups",
			args: args{
				i: 0,
				s: &university.Subject{
					Name: "Math",
					Groups: []*university.Group{
						{
							Capacity: 5,
							Name:     "1",
						},
						{
							Capacity: -1,
							Name:     "1",
						},
						{
							Capacity: 5,
							Name:     "2",
						},
						{
							Capacity: 5,
							Name:     "3",
						},
					},
				},
				students: []*university.Student{
					{
						Name: "a",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 1,
							{
								Subject: "Math",
								Group:   "3",
							}: 2,
						},
					},
					{
						Name: "b",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 2,
							{
								Subject: "Math",
								Group:   "3",
							}: 3,
						},
					},
				},
			},
			wantSgs: []*StudentGroup{
				{
					Student: &university.Student{
						Name: "a",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 1,
							{
								Subject: "Math",
								Group:   "3",
							}: 2,
						},
					},
					Group: &university.Group{
						Capacity: 5,
						Name:     "3",
					},
				},
				{
					Student: &university.Student{
						Name: "b",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 2,
							{
								Subject: "Math",
								Group:   "3",
							}: 3,
						},
					},
					Group: &university.Group{
						Capacity: 5,
						Name:     "2",
					},
				},
				{
					Student: &university.Student{
						Name: "b",
						Preferences: map[university.SubjectGroup]int{
							{
								Subject: "Math",
								Group:   "1",
							}: 1,
							{
								Subject: "Math",
								Group:   "2",
							}: 2,
							{
								Subject: "Math",
								Group:   "3",
							}: 3,
						},
					},
					Group: &university.Group{
						Capacity: 5,
						Name:     "3",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSgs := getStudents(tt.args.i, tt.args.likes, tt.args.s, tt.args.students); !cmp.Equal(gotSgs, tt.wantSgs) {
				t.Errorf("getStudents() = %v, want %v", gotSgs, tt.wantSgs)
			}
		})
	}
}
