package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pbartkowicz/scheduler/internal/university"
	"github.com/pbartkowicz/scheduler/internal/xlsx"
)

func main() {
	gf := flag.String("groups", "data/groups.xlsx", "Path to file containing groups")
	sd := flag.String("students", "data/students", "Path to directory containing students")
	sf := flag.String("priority", "data/priority_students.xlsx", "Path to file containing priority students")
	if err := readSchedule(*gf); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	if err := readStudents(*sd); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	if err := readPriorityStudents(*sf); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

func readSchedule(gf string) error {
	g, err := xlsx.Read(gf, true)
	if err != nil {
		return err
	}
	s, err := university.NewSchedule(g)
	if err != nil {
		return err
	}
	for _, i := range s.Subjects {
		fmt.Printf("%+v\n", i)
	}
	return nil
}

func readStudents(sd string) error {
	p, err := filepath.Abs(sd)
	if err != nil {
		return err
	}
	sfs, err := ioutil.ReadDir(p)
	if err != nil {
		return err
	}
	for _, sf := range sfs {
		s, err := xlsx.Read(sd+"/"+sf.Name(), true)
		if err != nil {
			return err
		}
		st, err := university.NewStudent(s, sf.Name())
		if err != nil {
			return err
		}
		fmt.Printf("%+v", st)
	}
	return nil
}

func readPriorityStudents(sf string) error {
	s, err := xlsx.Read(sf, false)
	if err != nil {
		return err
	}
	for _, ss := range s {
		fmt.Printf("%v\n", ss)
	}
	return nil
}
