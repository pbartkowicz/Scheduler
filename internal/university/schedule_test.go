package university

import (
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pbartkowicz/scheduler/test/tools"
)

func TestNewSchedule(t *testing.T) {
	type args struct {
		groups [][]string
	}
	tests := []struct {
		name string
		args args
		want *Schedule
		err  error
	}{
		{
			name: "Fails on GroupError",
			args: args{
				groups: [][]string{
					{
						"subject",
						"wrong",
					},
				},
			},
			err: &GroupError{
				Err: ErrWrongClassType,
			},
		},
		{
			name: "Successfully creates schedule",
			args: args{
				groups: [][]string{
					{
						"Programming",
						"W",
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
					{
						"Programming",
						"C",
						"teacher",
						"Monday",
						"14:00",
						"15:30",
						"C-2 313",
						"03-05-20",
						"2",
						"1b",
						"30",
					},
				},
			},
			want: &Schedule{
				Subjects: []*Subject{
					{
						Name: "Programming",
						Lectures: []*Group{
							{
								Type:      Lecture,
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
						Groups: []*Group{
							{
								Type:      Class,
								Teacher:   "teacher",
								Weekday:   time.Monday,
								StartTime: time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC),
								EndTime:   time.Date(0, 1, 1, 15, 30, 0, 0, time.UTC),
								Place:     "C-2 313",
								StartDate: time.Date(2020, 3, 5, 0, 0, 0, 0, time.UTC),
								Frequency: 2,
								Name:      "1b",
								Capacity:  30,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSchedule(tt.args.groups)
			if !cmp.Equal(err, tt.err, cmp.Comparer(tools.CompareErrors)) {
				t.Errorf("NewSchedule() error = %v, err %v", err, tt.err)
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("NewSchedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule_GetSubject(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		s    *Schedule
		want *Subject
	}{
		{
			name: "Returns nil because subject does not exist",
			args: args{
				n: "subject",
			},
			s: &Schedule{},
		},
		{
			name: "Successfully returns subject",
			args: args{
				n: "subject",
			},
			s: &Schedule{
				Subjects: []*Subject{
					{
						Name: "subject",
					},
					{
						Name: "subject2",
					},
				},
			},
			want: &Subject{
				Name: "subject",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GetSubject(tt.args.n)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Schedule.GetSubject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortSchedule(t *testing.T) {
	tests := []struct {
		name string
		sch  *Schedule
		want *Schedule
	}{
		{
			name: "Successfully sort subjects by number of conflicts",
			sch: &Schedule{
				Subjects: []*Subject{
					{
						Name: "Subject1",
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
					{
						Name: "Subject3",
						Groups: []*Group{
							{
								Name:     "1",
								Capacity: 1,
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
							},
							{
								Name:     "3",
								Capacity: 1,
							},
							{
								Name:     "4",
								Capacity: 1,
							},
						},
					},
					{
						Name: "Subject2",
						Groups: []*Group{
							{
								Name:     "1",
								Capacity: 1,
								Students: []*Student{
									{
										Name: "a",
									},
								},
							},
						},
					},
				},
			},
			want: &Schedule{
				Subjects: []*Subject{
					{
						Name: "Subject2",
						Groups: []*Group{
							{
								Name:     "1",
								Capacity: 1,
								Students: []*Student{
									{
										Name: "a",
									},
								},
							},
						},
					},
					{
						Name: "Subject1",
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
					{
						Name: "Subject3",
						Groups: []*Group{
							{
								Name:     "1",
								Capacity: 1,
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
							},
							{
								Name:     "3",
								Capacity: 1,
							},
							{
								Name:     "4",
								Capacity: 1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.sch)
			if !cmp.Equal(tt.sch, tt.want) {
				t.Errorf("sort.Sort(Schedule) got = %v, want %v", tt.sch, tt.want)
			}
		})
	}
}
