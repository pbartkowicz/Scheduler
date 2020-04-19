package main

import (
	"fmt"

	"github.com/pbartkowicz/Scheduler/internal/xlsx"
)

func main() {
	// p, err := open("../data/students/groups.xlsx")
	// if err != nil {
	// 	fmt.Printf("An error occurred: %s", err.Error())
	// 	return
	// }
	// if err := convert.Convert(p); err != nil {
	// 	fmt.Printf("An error occurred: %s", err.Error())
	// }
	if err := xlsx.Read("data/groups.xlsx", xlsx.Group); err != nil {
		fmt.Printf("An error occurred: %s", err.Error())
	}
}

// func open(f string) (string, error) {
// 	return filepath.Abs(f)
// }
