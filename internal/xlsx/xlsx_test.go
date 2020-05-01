package xlsx

import (
	"reflect"
	"testing"

	"github.com/pbartkowicz/scheduler/test/tools"
)

func TestError(t *testing.T) {
	e := &Error{
		Op:   ReadOp,
		File: "file.xlsx",
		Err:  ErrFileNotExists,
	}
	want := "read file.xlsx: file does not exists"
	got := e.Error()
	if got != want {
		t.Errorf("Error() got = %v, want %v", got, want)
	}
}

func TestRead(t *testing.T) {
	type args struct {
		n    string
		skip bool
	}
	tests := []struct {
		name string
		args args
		want [][]string
		err  error
	}{
		{
			name: "Fails on non-existing file",
			args: args{
				n: "no_such_file.xlsx",
			},
			err: &Error{
				Op:   ReadOp,
				File: "no_such_file.xlsx",
				Err:  ErrFileNotExists,
			},
		},
		{
			name: "Successfully reads data with skipping first row",
			args: args{
				n:    "../../test/data/xlsx/skip.xlsx",
				skip: true,
			},
			want: [][]string{
				{"AA", "AA"},
				{"BB", "BB"},
				{"CC", "CC"},
			},
		},
		{
			name: "Successfully reads data without shipping first row",
			args: args{
				n: "../../test/data/xlsx/skip.xlsx",
			},
			want: [][]string{
				{"name", "surname"},
				{"AA", "AA"},
				{"BB", "BB"},
				{"CC", "CC"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.n, tt.args.skip)
			if !tools.CompareErrors(err, tt.err) {
				t.Errorf("Read() error = %v, err %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}
