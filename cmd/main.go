package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pbartkowicz/scheduler/internal/schedule"
	"github.com/pbartkowicz/scheduler/internal/university"
	"github.com/pbartkowicz/scheduler/internal/xlsx"
)

func main() {
	gf := flag.String("groups", "data/groups.xlsx", "Path to file containing groups")
	sd := flag.String("students", "data/students", "Path to directory containing students")
	psf := flag.String("priority", "data/priority_students.xlsx", "Path to file containing priority students")

	sch, err := readSchedule(*gf)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	students, err := readStudents(*sd)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	if err := readPriorityStudents(*psf, students); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	schedule.Enroll(sch, students)
}

func readSchedule(gf string) (*university.Schedule, error) {
	g, err := xlsx.Read(gf)
	if err != nil {
		return nil, err
	}
	s, err := university.NewSchedule(g)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func readStudents(sd string) ([]*university.Student, error) {
	p, err := filepath.Abs(sd)
	if err != nil {
		return nil, err
	}
	sfs, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}
	var students []*university.Student
	for _, sf := range sfs {
		pref, err := xlsx.Read(sd + "/" + sf.Name())
		if err != nil {
			return nil, err
		}
		st, err := university.NewStudent(pref, sf.Name())
		if err != nil {
			return nil, err
		}
		students = append(students, st)
	}
	return students, nil
}

func readPriorityStudents(psf string, students []*university.Student) error {
	ps, err := xlsx.Read(psf)
	if err != nil {
		return err
	}
	for _, p := range ps {
		var found bool
		for _, st := range students {
			if p[0] == st.Name {
				st.Priority = true
				found = true
				continue
			}
		}
		if !found {
			return fmt.Errorf("missing %s student", p[0])
		}
	}
	return nil
}
