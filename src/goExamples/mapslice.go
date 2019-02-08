package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"texas" : []string{"austin", "dallas", "houston", "sa"},
		"sfo": []string{"sfo", "la", "sd"},
	}
	
	for _, v := range m {
		for index, val := range v {
			fmt.Println(index, val)
		}
	}
	
	
}
