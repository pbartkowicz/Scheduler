package xlsx

import (
	"fmt"
	"os"
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
			name: "Successfully reads data",
			args: args{
				n:    "../../test/data/xlsx/read.xlsx",
				skip: true,
			},
			want: [][]string{
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

func TestWrite(t *testing.T) {
	p := "./tmp"
	type args struct {
		n  string
		p  string
		s  string
		dd [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
		err  error
	}{
		{
			name: "Successfully creates a file",
			args: args{
				n: "new-file",
				p: p,
				s: "Sheet1",
				dd: [][]string{
					{
						"aaa", "aaa", "aaa",
					},
					{
						"bbb", "bbb", "bbb",
					},
					{
						"ccc", "ccc", "ccc",
					},
				},
			},
			want: [][]string{
				{
					"aaa", "aaa", "aaa",
				},
				{
					"bbb", "bbb", "bbb",
				},
				{
					"ccc", "ccc", "ccc",
				},
			},
		},
	}
	// Create tmp directory for test
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.Mkdir(p, os.ModePerm)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Write(tt.args.n, tt.args.p, tt.args.s, tt.args.dd)
			if !tools.CompareErrors(err, tt.err) {
				t.Errorf("Write() error = %v, err %v", err, tt.err)
			}
			got, _ := Read(fmt.Sprintf("%s/%s.xlsx", tt.args.p, tt.args.n), false)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
	os.RemoveAll(p)
}
