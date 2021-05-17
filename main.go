package main

import (
	"fmt"
)

func main() {
	population := map[string]int{"TamilNadu": 6000000, "MadhyaPradesh": 7000000,
		"karnataka": 5000000}
	pop, ok := population["TamilNadu"]
	fmt.Println("population_map_details:", population)
	fmt.Println("Map access with key:", population["TamilNadu"])
	fmt.Println("delete the value from map:", delete(population, "MadhyaPradesh"))
	fmt.Println("after deleted the map looks like:", population)
	fmt.Println("If we try to access the key of deleted value:", population["MadhyaPradesh"])
	fmt.Println("By Using OK to know the existence of key:")
	fmt.Println(pop, ok)
}
