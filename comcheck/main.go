package main

import (
	"fmt"
	"os"
)

func main() {
	strs := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	arguments := os.Args[]

	for i := range arguments {
		for _, s := range strs {
			if arguments[i] == s {
				count++
			}
		}
	}
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
