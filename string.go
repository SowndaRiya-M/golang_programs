package main

import "fmt"

func main() {
	for _, i := range "HelloTeam" {
		fmt.Println(string(i))
		if i == 'T' {
			break
		}
	}
}
