package main

import (
	"fmt"
)

// given an image represented in an NxN matrix, write a method to rotate the image 90 degrees
// can you do this in place?

// i dunno what's going on tonight, but i just can't do this one

func main() {

	const size = 5
	matrix := [size][size]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25}}

	fmt.Println("Testing function rotateMatrix")
	fmt.Print("First call ")
	matrix = rotateMatrix(matrix)
	fmt.Println(matrix)
	fmt.Print("Second call ")
	matrix = rotateMatrix(matrix)
	fmt.Println(matrix)
	fmt.Print("Two more calls ")
	matrix = rotateMatrix(matrix)
	matrix = rotateMatrix(matrix)
	fmt.Println(matrix)
}

func rotateMatrix(arr [5][5]int) [5][5]int {
	const size = 5
	var matrix [size][size]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[j][i] = arr[i][j]
		}
	}

	return matrix
}
