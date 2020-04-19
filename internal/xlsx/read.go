// Package xlsx ...
package xlsx

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pbartkowicz/scheduler/internal/schedule"
)

// ReadOptions ...
type ReadOptions string

const (
	Group            ReadOptions = "Group"
	PriorityStudents ReadOptions = "Priority Students"
	Student          ReadOptions = "Student"
)

// Read ...
func Read(n string, op ReadOptions) error {
	p, err := filepath.Abs(n)
	if err != nil {
		return err
	}
	f, err := excelize.OpenFile(p)
	if err != nil {
		return err
	}
	s := f.GetSheetName(1)
	if s == "" {
		return errors.New("invalid")
	}
	rows, err := f.Rows(s)
	if err != nil {
		return errors.New("invalid")
	}
	/*if !rows.Next() {
		return errors.New("invalid")
	}
	names := rows.Columns()*/
	/*for _, ff := range names {
		fmt.Printf("%v\n", ff)
	}*/
	/*for rows.Next() {
		row := rows.Columns()
		for _, col := range row {
			fmt.Printf("%T %v\n", col, col)
		}
	}*/
	switch op {
	case Group:
		createGroups(rows)
	}
	return nil
}

func createGroups(rows *excelize.Rows) error {
	// TODO - firstly create subjects
	if !rows.Next() {
		return errors.New("invalid")
	}
	//names := rows.Columns()
	for rows.Next() {
		row := rows.Columns()
		g := schedule.NewGroup(row)
		fmt.Printf("%+v", g)
	}
	return nil
}

/*func newGroup(names []string, values []string) *schedule.Group {
	return
}*/
