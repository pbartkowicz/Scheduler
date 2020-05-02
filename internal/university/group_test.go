package university

import (
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
