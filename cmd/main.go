package main

import (
	"fmt"

	"github.com/pbartkowicz/scheduler/internal/xlsx"
)

func main() {
	g, err := xlsx.Read("data/groups.xlsx")
	if err != nil {
		fmt.Printf("An error occurred: %s", err.Error())
		return
	}
	for _, gg := range g {
		for _, ggg := range gg {
			fmt.Printf("%v", ggg)
		}
		fmt.Printf("\n")
	}
}
