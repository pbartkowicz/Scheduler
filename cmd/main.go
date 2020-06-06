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
	psf := flag.String("priority", "data/priority_students.xlsx", "Path to file containing priority students")
	rf := flag.String("result", "data/result", "Path to the directory where the results will be saved")

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

	sch.Enroll(students)

	if err := saveStudents(students, *rf); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	if err := saveSubjects(sch, *rf); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

func readSchedule(gf string) (*university.Schedule, error) {
	g, err := xlsx.Read(gf, true)
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
		pref, err := xlsx.Read(sd+"/"+sf.Name(), true)
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
	ps, err := xlsx.Read(psf, true)
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

func saveStudents(students []*university.Student, p string) error {
	for _, st := range students {
		if err := xlsx.Write(st.Name, p, st.Name, st.Save()); err != nil {
			return err
		}
	}
	return nil
}

func saveSubjects(schedule *university.Schedule, p string) error {
	for _, sub := range schedule.Subjects {
		saveGroups(sub.Name, p, sub.Groups)
		saveGroups(sub.Name, p, sub.Lectures)
	}
	return nil
}

func saveGroups(sn, p string, grs []*university.Group) error {
	for _, g := range grs {
		if err := xlsx.Write(sn, p, g.Name, g.Save()); err != nil {
			return err
		}
	}
	return nil
}
