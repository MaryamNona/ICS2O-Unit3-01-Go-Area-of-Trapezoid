// Created by: Maryam Nona
// Created on: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main() {
	// This function does addition
	var aBase int
	var bBase int
	var height int

	// input
	fmt.Println("This program finds the area of a trapezoid")
	fmt.Println()
	fmt.Println("Enter the a base (mm): ")
	fmt.Scanln(&aBase)
	fmt.Println("Enter the b base (mm): ")
	fmt.Scanln(&bBase)
	fmt.Println("Enter the height (mm): ")
	fmt.Scanln(&height)

	// output
	fmt.Println("The area is: ", ((aBase+bBase)/2)*height, "mm²")
	fmt.Println("Done.")
}