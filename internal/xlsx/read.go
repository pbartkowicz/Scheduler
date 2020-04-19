// Package xlsx ...
package xlsx

import (
	"errors"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Read ...
func Read(n string) ([][]string, error) {
	data := make([][]string, 0)
	p, err := filepath.Abs(n)
	if err != nil {
		return data, err
	}
	f, err := excelize.OpenFile(p)
	if err != nil {
		return data, err
	}
	s := f.GetSheetName(1)
	if s == "" {
		return data, errors.New("invalid")
	}
	rows, err := f.Rows(s)
	if err != nil {
		return data, errors.New("invalid")
	}
	// Skip first line
	rows.Next()
	for rows.Next() {
		data = append(data, rows.Columns())
	}
	return data, nil
}
