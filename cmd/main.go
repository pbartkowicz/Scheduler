package main

import (
	"flag"
	"fmt"

	"github.com/pbartkowicz/scheduler/internal/university"
	"github.com/pbartkowicz/scheduler/internal/xlsx"
)

func main() {
	gf := flag.String("groups", "data/groups.xlsx", "Path to file containing groups")
	if err := createSchedule(*gf); err != nil {
		fmt.Printf(err.Error())
	}
}

// TODO - change package name to university
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
