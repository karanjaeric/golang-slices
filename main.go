package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Slices are what you need when working sequence of data")

	var universePlanets = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	fmt.Printf("Here are the unoverse planet %v\n", universePlanets)

	fmt.Printf("The Blue Planet is %s\n", universePlanets[2])
	fmt.Printf("The Largest Planet is %s\n", universePlanets[4])
	fmt.Printf("The Smallest Planet is %s\n", universePlanets[0])

	var x []int
	fmt.Printf("X value is %v", x)
	fmt.Println(x != nil)

	var y = []int{1, 2, 3, 4, 5}
	var z = []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Equal(y, z))

	fmt.Printf("The size of x is %d\n", len(x))
	fmt.Printf("The size of y is %d\n", len(y))
	x = append(x, 10)
	fmt.Printf("The new size of x is %d\n", len(x))
	fmt.Printf("The length of universePlanet is %d\n", len(universePlanets))
	fmt.Printf("The capacity of universePlanet is %d\n", cap(universePlanets))
	universePlanets = append(universePlanets, "pluto")
	fmt.Printf("The new capacity after adding pluto to universePlanet is %d\n", cap(universePlanets))
}
