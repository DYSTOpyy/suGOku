package main

import (
	"math"
)

func Is_In (value int, slice []int) bool {

	for _, v := range slice {
		if (value == v) {
			return true
		}
	}
	return false

}

func IsOkayCase(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, size int, value int) bool {
	line := IsOkayLine(x, y, grid, size, value)
	column := IsOkayColumn(x, y, grid, size, value)
	box := IsOkayBox(x, y, grid, size, value)
	return line && column && box
}

func IsOkayLine(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, size int, value int) bool {
	i := 0
	for i < size && grid[x][i] != value {
		i++
	}
	return i == size
}

func IsOkayColumn(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, size int, value int) bool {
	i := 0
	for i < size && grid[i][y] != value {
		i++
	}
	return i == size
}

func IsOkayBox(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, size int, value int) bool {
	root := int(math.Sqrt(float64(size)))
	i := 0
	for i < size && grid[int(x/root)*root+int(i/root)][int(y/root)*root+i%root] != value {
		i++
	}
	return i == size
}