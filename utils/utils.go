package utils

import (
	"math"
)



func IsOkayCase(x int, y int, grid [TAILLE+2][TAILLE+1]int,size int, value int) bool {
	line := IsOkayLine(x,y,grid,size,value)	
	column := IsOkayColumn(x,y,grid,size,value)	
	box := IsOkayBox(x,y,grid,size,value)	
	return line && column && box
}

func IsOkayLine(x int, y int, grid [TAILLE+2][TAILLE+1]int, size int, value int) bool {
	i:=0
	for i < size && grid[x][i] != value {
		i++
	}
	return i == size
}

func IsOkayColumn(x int, y int, grid [TAILLE+2][TAILLE+1]int, size int, value int) bool {
	i:=0
	for i < size && grid[i][y] != value {
		i++
	}
	return i == size
}

func IsOkayBox(x int, y int, grid [TAILLE+2][TAILLE+1]int, size int, value int) bool {
	root := int(math.Sqrt(float64(size)))
	i:=0
	for i < size && grid[int(x/root)*root+int(i/root)][int(y/root)*root+i%root] != value {
		i++
	}
	return i == size
}

func Is_In (value int, slice []int) bool {

	for _, v := range slice {
		if (value == v) {
			return true
		}
	}
	return false

}

// maj_compteurs met à jour les compteurs du nombre d'indices par ligne/colonne/bloc
// d'une grille qui sont situés respectivement sur la colonne TAILLE,
// la ligne TAILLE et la ligne TAILLE+1, ainsi que le nombre total d'indice
// de la grille situé en TAILLE:TAILLE.
//
// La grille en argument doit être sous forme de pointeur.
func maj_compteurs(grille *[TAILLE + 2][TAILLE + 1]int) {

	for line := 0; line < TAILLE; line++ {
		for column := 0; column < TAILLE; column++ {
			if grille[line][column] != 0 {

				grille[line][TAILLE] += 1                                                                                                              // maj des lignes
				grille[TAILLE][column] += 1                                                                                                            // maj des colonnes
				grille[TAILLE+1][(line/int(math.Sqrt(float64(TAILLE))))*int(math.Sqrt(float64(TAILLE)))+(column/int(math.Sqrt(float64(TAILLE))))] += 1 // maj des bloc
				grille[TAILLE][TAILLE] += 1                                                                                                            // maj du nombre total d'indice
			}
		}
	}
}