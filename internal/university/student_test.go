package university

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pbartkowicz/scheduler/test/tools"
)

func TestNewStudent(t *testing.T) {
	type args struct {
		v [][]string
		n string
	}
	tests := []struct {
		name string
		args args
		want *Student
		err error
	}{
		{
			name: "Incorrect priority",
			args: args{
				v: [][]string{
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
			},
		},
		{
			name: "Successfully creates student",
			args: args{
				v: [][]string{
					{"subject1", "g1", "1"},
					{"subject1", "g2", "1"},
					{"subject2", "g1", "1"},
					{"subject2", "g2", "2"},
				},
				n: "student.xlsx",
			},
			want: &Student{
				Name: "student",
				Happieness: 100,
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 1,
					{"subject1", "g2"}: 1,
					{"subject2", "g1"}: 1,
					{"subject2", "g2"}: 2,
				},
				FinalGroups: make(map[string]*Group),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStudent(tt.args.v, tt.args.n)
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
	tests := []struct{
		name string
		s *Student
		err error
	}{
		{
			name: "Fails on first number lesser than one",
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: -1,
				},
			},
			err: &StudentError{
				Err: ErrWrongPriority,
			},
		},
		{
			name: "Fails on number greater than one",
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 10,
				},
			},
			err: &StudentError{
				Err: ErrWrongPriority,
			},
		},
		{
			name: "Fails difference bigger than one",
			s: &Student{
				Preferences: map[SubjectGroup]int{
					{"subject1", "g1"}: 1,
					{"subject1", "g2"}: 3,
				},
			},
			err: &StudentError{
				Err: ErrWrongSubPriority,
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