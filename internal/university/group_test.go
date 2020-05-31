package university

import (
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pbartkowicz/scheduler/test/tools"
)

func TestNewGroup(t *testing.T) {
	type args struct {
		subjects []string
	}
	tests := []struct {
		name string
		args args
		want *Group
		err  error
	}{
		{
			name: "Incorrect class type",
			args: args{
				[]string{
					"Programming",
					"Wrong",
				},
			},
			err: &GroupError{
				Err: ErrWrongClassType,
			},
		},
		{
			name: "Incorrect weekday",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"wrong",
				},
			},
			err: &GroupError{
				Err: ErrWrongWeekday,
			},
		},
		{
			name: "Incorrect time format for start time",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"Thursday",
					"14:00:00",
				},
			},
			err: &GroupError{
				Err: &time.ParseError{
					Value:   "14:00:00",
					Message: ": extra text: :00",
				},
			},
		},
		{
			name: "Incorrect time format for end time",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"Thursday",
					"14:00",
					"15:30:00",
				},
			},
			err: &GroupError{
				Err: &time.ParseError{
					Value:   "15:30:00",
					Message: ": extra text: :00",
				},
			},
		},
		{
			name: "Incorrect date format for start date",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"Thursday",
					"14:00",
					"15:30",
					"C-2 313",
					"03-05-2020",
				},
			},
			err: &GroupError{
				Err: &time.ParseError{
					Value:   "03-05-2020",
					Message: ": extra text: 20",
				},
			},
		},
		{
			name: "Incorrect frequency",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"Thursday",
					"14:00",
					"15:30",
					"C-2 313",
					"03-05-20",
					"wrong",
				},
			},
			err: &GroupError{
				&strconv.NumError{
					Func: "Atoi",
					Num:  "wrong",
					Err:  strconv.ErrSyntax,
				},
			},
		},
		{
			name: "Incorrect capacity",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"Thursday",
					"14:00",
					"15:30",
					"C-2 313",
					"03-05-20",
					"2",
					"1b",
					"wrong",
				},
			},
			err: &GroupError{
				Err: &strconv.NumError{
					Func: "Atoi",
					Num:  "wrong",
					Err:  strconv.ErrSyntax,
				},
			},
		},
		{
			name: "Successfully creates group",
			args: args{
				[]string{
					"Programming",
					"L",
					"teacher",
					"Thursday",
					"14:00",
					"15:30",
					"C-2 313",
					"03-05-20",
					"2",
					"1b",
					"30",
				},
			},
			want: &Group{
				Type:      Laboratory,
				Teacher:   "teacher",
				Weekday:   time.Thursday,
				StartTime: time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC),
				EndTime:   time.Date(0, 1, 1, 15, 30, 0, 0, time.UTC),
				Place:     "C-2 313",
				StartDate: time.Date(2020, 3, 5, 0, 0, 0, 0, time.UTC),
				Frequency: 2,
				Name:      "1b",
				Capacity:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGroup(tt.args.subjects)
			if !cmp.Equal(err, tt.err, cmp.Comparer(tools.CompareErrors)) {
				t.Errorf("NewGroup() error = %v, err %v", err, tt.err)
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("NewGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Conflicts(t *testing.T) {
	tests := []struct {
		name string
		g    *Group
		want int
	}{
		{
			name: "Returns 0 conflicts",
			g:    &Group{},
		},
		{
			name: "Returns number of conflicts",
			g: &Group{
				Capacity: 1,
				PriorityStudents: []*Student{
					{
						Name: "a",
					},
				},
				Students: []*Student{
					{
						Name: "b",
					},
					{
						Name: "c",
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.Conflicts()
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Group.Conflicts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Collide(t *testing.T) {
	type args struct {
		a *Group
	}
	tests := []struct {
		name string
		args args
		g    *Group
		want bool
	}{
		{
			name: "Returns false because of the different weekdays",
			args: args{
				a: &Group{
					Weekday: time.Monday,
				},
			},
			g: &Group{
				Weekday: time.Tuesday,
			},
		},
		{
			name: "Returns false because of the different start time",
			args: args{
				a: &Group{
					Weekday:   time.Monday,
					StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
				},
			},
			g: &Group{
				Weekday:   time.Monday,
				StartTime: time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Returns true because groups collide",
			args: args{
				a: &Group{
					Weekday:   time.Monday,
					StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
				},
			},
			g: &Group{
				Weekday:   time.Monday,
				StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.Collide(tt.args.a)
			if got != tt.want {
				t.Errorf("Group.Collide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_RemoveStudent(t *testing.T) {
	type args struct {
		st *Student
	}
	tests := []struct {
		name string
		args args
		g    *Group
	}{
		{
			name: "Successfully removes student",
			args: args{
				st: &Student{
					Name: "student",
				},
			},
			g: &Group{
				Students: []*Student{
					{
						Name: "studentA",
					},
					{
						Name: "student",
					},
					{
						Name: "studentB",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.RemoveStudent(tt.args.st)
			for _, s := range tt.g.Students {
				if cmp.Equal(tt.args.st, s) {
					t.Errorf("Group.RemoveStudent() did not remove a student %v", tt.args.st)
				}
			}
		})
	}
}

func TestSortGroup(t *testing.T) {
	tests := []struct {
		name string
		g    *Group
		want *Group
	}{
		{
			name: "Successfully sort students by happieness",
			g: &Group{
				Students: []*Student{
					{
						Name: "a",
						Happieness: map[string]float64{
							"Math":        100.0,
							"Programming": 50.0,
							"Algorithms":  25.0,
						},
					},
					{
						Name: "b",
						Happieness: map[string]float64{
							"Math":        100.0,
							"Programming": 100.0,
							"Algorithms":  100.0,
						},
					},
					{
						Name: "c",
						Happieness: map[string]float64{
							"Math":        100.0,
							"Programming": 50.0,
							"Algorithms":  100.0,
						},
					},
				},
			},
			want: &Group{
				Students: []*Student{
					{
						Name: "b",
						Happieness: map[string]float64{
							"Math":        100.0,
							"Programming": 100.0,
							"Algorithms":  100.0,
						},
					},
					{
						Name: "c",
						Happieness: map[string]float64{
							"Math":        100.0,
							"Programming": 50.0,
							"Algorithms":  100.0,
						},
					},
					{
						Name: "a",
						Happieness: map[string]float64{
							"Math":        100.0,
							"Programming": 50.0,
							"Algorithms":  25.0,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.g)
			if !cmp.Equal(tt.g, tt.want) {
				t.Errorf("sort.Sort(Group) got = %v, want %v", tt.g, tt.want)
			}
		})
	}
}
