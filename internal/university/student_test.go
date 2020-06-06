package university

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pbartkowicz/scheduler/test/tools"
)

func TestNewStudent(t *testing.T) {
	type args struct {
		pref [][]string
		n    string
	}
	tests := []struct {
		name string
		args args
		want *Student
		err  error
	}{
		{
			name: "Incorrect priority",
			args: args{
				pref: [][]string{
					{"subject1", "g1", "wrong"},
				},
				n: "student.xlsx",
			},
			err: &StudentError{
				Err: &strconv.NumError{
					Func: "Atoi",
					Num:  "wrong",
					Err:  strconv.ErrSyntax,
				},
				Name: "student",
			},
		},
		{
			name: "Successfully creates student",
			args: args{
				pref: [][]string{
					{"subject1", "g1", "1"},
					{"subject1", "g2", "1"},
					{"subject2", "g1", "1"},
					{"subject2", "g2", "2"},
				},
				n: "student.xlsx",
			},
			want: &Student{
				Name: "student",
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 1,
					{"subject1", "g2"}: 1,
					{"subject2", "g1"}: 1,
					{"subject2", "g2"}: 2,
				},
				FinalGroups: make(map[string]*Group),
				Happieness:  make(map[string]float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStudent(tt.args.pref, tt.args.n)
			if !cmp.Equal(err, tt.err, cmp.Comparer(tools.CompareErrors)) {
				t.Errorf("NewStudent() error = %v, err %v", err, tt.err)
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("NewStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_validate(t *testing.T) {
	tests := []struct {
		name string
		s    *Student
		err  error
	}{
		{
			name: "Fails on first number lesser than one",
			s: &Student{
				Name: "student",
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: -1,
				},
			},
			err: &StudentError{
				Err:  ErrWrongPriority,
				Name: "student",
			},
		},
		{
			name: "Fails on number greater than one",
			s: &Student{
				Name: "student",
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 10,
				},
			},
			err: &StudentError{
				Err:  ErrWrongPriority,
				Name: "student",
			},
		},
		{
			name: "Fails difference bigger than one",
			s: &Student{
				Name: "student",
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 1,
					{"subject1", "g2"}: 3,
				},
			},
			err: &StudentError{
				Err:  ErrWrongSubPriority,
				Name: "student",
			},
		},
		{
			name: "Successfully validates groups",
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 1,
					{"subject1", "g2"}: 1,
					{"subject1", "g3"}: 2,
				},
			},
		},
		{
			name: "Successfully validates group",
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.validate()
			if !cmp.Equal(err, tt.err, cmp.Comparer(tools.CompareErrors)) {
				t.Errorf("NewStudent() error = %v, err %v", err, tt.err)
			}
		})
	}
}

func TestStudent_GetPreferredGroup(t *testing.T) {
	type args struct {
		subject string
		groups  []string
	}
	tests := []struct {
		name    string
		args    args
		s       *Student
		wantRes string
	}{
		{
			name: "Successfully returns the most wanted group name",
			args: args{
				subject: "Math",
				groups:  []string{"1", "2", "3"},
			},
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{
						Subject: "Math",
						Group:   "1",
					}: 2,
					{
						Subject: "Math",
						Group:   "2",
					}: 1,
					{
						Subject: "Math",
						Group:   "3",
					}: 3,
				},
			},
			wantRes: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.s.GetPreferredGroup(tt.args.subject, tt.args.groups); gotRes != tt.wantRes {
				t.Errorf("Student.GetPreferredGroup() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStudent_Likes(t *testing.T) {
	type args struct {
		sub string
		gn  string
	}
	tests := []struct {
		name string
		args args
		s    *Student
		want bool
	}{
		{
			name: "Student likes this group",
			args: args{
				sub: "Math",
				gn:  "1",
			},
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{
						Subject: "Math",
						Group:   "1",
					}: 1,
				},
			},
			want: true,
		},
		{
			name: "Student doesn't like this group",
			args: args{
				sub: "Math",
				gn:  "1",
			},
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{
						Subject: "Math",
						Group:   "1",
					}: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Likes(tt.args.sub, tt.args.gn); got != tt.want {
				t.Errorf("Student.Likes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_CanMove(t *testing.T) {
	type args struct {
		sub string
		g   *Group
	}
	tests := []struct {
		name string
		args args
		s    *Student
		want bool
	}{
		{
			name: "Student can be moved to other group",
			args: args{
				sub: "Math",
				g: &Group{
					Weekday:   time.Monday,
					StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
				},
			},
			s: &Student{
				FinalGroups: map[string]*Group{
					"Math2.0": {
						Weekday:   time.Friday,
						StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
					},
					"Programming": {
						Weekday:   time.Monday,
						StartTime: time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC),
					},
				},
			},
			want: true,
		},
		{
			name: "Student can't be moved to other groups",
			args: args{
				sub: "Math",
				g: &Group{
					Weekday:   time.Monday,
					StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
				},
			},
			s: &Student{
				FinalGroups: map[string]*Group{
					"Math2.0": {
						Weekday:   time.Friday,
						StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
					},
					"Programming": {
						Weekday:   time.Monday,
						StartTime: time.Date(0, 0, 0, 12, 30, 0, 0, time.UTC),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CanMove(tt.args.sub, tt.args.g); got != tt.want {
				t.Errorf("Student.CanMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_GetHappieness(t *testing.T) {
	tests := []struct {
		name    string
		s       *Student
		wantRes float64
	}{
		{
			name: "Successfully counts student's happieness",
			s: &Student{
				Happieness: map[string]float64{
					"Math":        100.0,
					"Programming": 50.0,
					"Algorithms":  25.0,
				},
			},
			wantRes: (100.0 + 50.0 + 25.0) / 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.s.GetHappieness(); gotRes != tt.wantRes {
				t.Errorf("Student.GetHappieness() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStudent_CalculateHappieness(t *testing.T) {
	type args struct {
		sn string
	}
	tests := []struct {
		name string
		args args
		s    *Student
		want map[string]float64
	}{
		{
			name: "Successfully calculates student's happieness",
			args: args{
				sn: "Math",
			},
			s: &Student{
				Happieness: make(map[string]float64),
				Preferences: map[SubjectGroup]int{
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
			want: map[string]float64{
				"Math": (1.0 / float64(3)) * 100.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.CalculateHappieness(tt.args.sn)
			if !cmp.Equal(tt.s.Happieness, tt.want) {
				t.Errorf("Student.CalculateHappieness() = %v, want %v", tt.s.Happieness, tt.want)
			}
		})
	}
}

func TestStudent_Save(t *testing.T) {
	tests := []struct {
		name string
		s    *Student
		want [][]string
	}{
		{
			name: "Successfully saves groups for student",
			s: &Student{
				FinalGroups: map[string]*Group{
					"Math":        {Name: "1"},
					"Programming": {Name: "2a"},
					"Algorithms":  nil,
				},
			},
			want: [][]string{
				{
					"Math", "1",
				},
				{
					"Programming", "2a",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.s.Save(); !reflect.DeepEqual(gotRes, tt.want) {
				t.Errorf("Student.Save() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}
