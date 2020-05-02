package tools

import (
	"errors"
	"testing"
)

type tError struct {
	Err error
}

func (e *tError) Error() string {
	return e.Err.Error()
}

func TestCompareErrors(t *testing.T) {
	type args struct {
		a, e error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Returns false because actual error is nil",
			args: args{
				a: errors.New("unexpected"),
			},
		},
		{
			name: "Returns false because expected error is nil",
			args: args{
				e: errors.New("unexpected"),
			},
		},
		{
			name: "Returns false because errors are different types",
			args: args{
				a: errors.New("expected"),
				e: &tError{Err: errors.New("expected")},
			},
		},
		{
			name: "Returns true because both errors are nil",
			want: true,
		},
		{
			name: "Returns true because error messages are the same",
			args: args{
				a: errors.New("expected"),
				e: errors.New("expected"),
			},
			want: true,
		},
		{
			name: "Returns false because error messages are different",
			args: args{
				a: errors.New("expected"),
				e: errors.New("unexpected"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareErrors(tt.args.a, tt.args.e)
			if got != tt.want {
				t.Errorf("CompareErrors() got = %v, want %v", got, tt.want)
			}
		})
	}
}
