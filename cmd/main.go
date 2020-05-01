package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pbartkowicz/scheduler/internal/university"
	"github.com/pbartkowicz/scheduler/internal/xlsx"
)

func main() {
	gf := flag.String("groups", "data/groups.xlsx", "Path to file containing groups")
	sf := flag.String("students", "data/students", "Path to directory containing students")
	if err := createSchedule(*gf); err != nil {
		fmt.Printf(err.Error())
		return
	}
	if err := readStudents(*sf); err != nil {
		fmt.Printf(err.Error())
		return
	}
}

// TODO - move it as new schedule
func createSchedule(gf string) error {
	s := &university.Schedule{}
	g, err := xlsx.Read(gf, true)
	if err != nil {
		return err
	}
	for _, gg := range g {
		ng, err := university.NewGroup(gg)
		if err != nil {
			return err
		}
		sub := s.GetSubject(gg[0])
		if sub != nil {
			if ng.Type == university.Lecture {
				sub.Lectures = append(sub.Lectures, ng)
			} else {
				sub.Groups = append(sub.Groups, ng)
			}
		} else {
			sub = &university.Subject{
				Name: gg[0],
			}
			if ng.Type == university.Lecture {
				sub.Lectures = append(sub.Lectures, ng)
			} else {
				sub.Groups = append(sub.Groups, ng)
			}
			s.Subjects = append(s.Subjects, sub)
		}
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
