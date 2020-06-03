// Package xlsx provides standard operations on files with .xlsx format.
package xlsx

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var (
	// ErrPathNotExists is returned when path to a file was not found.
	ErrPathNotExists = errors.New("path to a file does not exists")
	// ErrFileNotExists is returned when file was not found.
	ErrFileNotExists = errors.New("file does not exists")
	// ErrSheetNotExists is returned when sheet was not found.
	ErrSheetNotExists = errors.New("sheet does not exists in a file")
	// ErrRows is returned when rows could not be read.
	ErrRows = errors.New("could not read rows")
)

// Operation is a type of action which can be performed on a .xlsx file.
// It is used in error messages.
type Operation string

const (
	// ReadOp is a read operation.
	ReadOp Operation = "read"
	// WriteOp is a write operation.
	WriteOp Operation = "write"
)

// Error represents an error struct returned by this package.
type Error struct {
	Op   Operation
	File string
	Err  error
}

func (e *Error) Error() string {
	return string(e.Op) + " " + e.File + ": " + e.Err.Error()
}

// Read retrieves data from the first sheet of file n.
// If skip is set to true, it skips the first line.
// It returns an error if an absolute path of file was not found
func Read(n string, skip bool) ([][]string, error) {
	p, err := filepath.Abs(n)
	if err != nil {
		return nil, &Error{Op: ReadOp, File: n, Err: ErrPathNotExists}
	}
	f, err := excelize.OpenFile(p)
	if err != nil {
		return nil, &Error{Op: ReadOp, File: n, Err: ErrFileNotExists}
	}
	s := f.GetSheetName(1)
	if s == "" {
		return nil, &Error{Op: ReadOp, File: n, Err: ErrSheetNotExists}
	}
	rows, err := f.Rows(s)
	if err != nil {
		return nil, &Error{Op: ReadOp, File: n, Err: ErrRows}
	}
	// Skip heading
	if skip {
		rows.Next()
	}
	data := make([][]string, 0)
	for rows.Next() {
		data = append(data, rows.Columns())
	}
	return data, nil
}

// Write creates file with a given name in a given path and saves passed data in it.
func Write(n, p string, dd [][]string) error {
	rp, err := filepath.Abs(p)
	if err != nil {
		return &Error{Op: WriteOp, File: p, Err: ErrPathNotExists}
	}
	f := excelize.NewFile()
	s := "Sheet1"
	f.NewSheet(s)
	for i, d := range dd {
		f.SetSheetRow(s, fmt.Sprintf("A%v", i+1), &d)
	}
	return f.SaveAs(fmt.Sprintf("%s/%s.xlsx", rp, n))
}
