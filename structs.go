package main

import (
	"fmt"
)

type struc_eg struct {
	number    int
	name      string
	Ocupation []string
}

func main() {

	dhanush := struc_eg{
		1, "Dhanush", []string{"Lyricist", "Producer", "actor", "Director", "Singer", "Screenwriter"},
	}
	fmt.Println(dhanush)

}
