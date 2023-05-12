package main

import (
	"fmt" // I/O
	"math"
)

const TAILLE int = 9

func maj_compteurs(grille [][]int) {
	for line := 0; line < TAILLE; line++ {
		for column := 0; column < TAILLE; column++ {
			if grille[line][column] != 0 {
				grille[line][TAILLE] += 1   // maj des lignes
				grille[TAILLE][column] += 1 // maj des colonnes
				grille[TAILLE][TAILLE] += 1
				grille[TAILLE+1][(line/int(math.Sqrt(float64(TAILLE))))*int(math.Sqrt(float64(TAILLE)))+(column/int(math.Sqrt(float64(TAILLE))))] += 1
			}
		}
	}
}

func main() {

	grille := [][]int{{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0, 0},
		{0, 3, 2, 1, 9, 0, 0, 8, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 4, 0},
		{0, 0, 4, 9, 0, 0, 0, 6, 1, 0},
		{0, 0, 9, 0, 7, 4, 0, 0, 0, 0},
		{0, 9, 6, 7, 0, 1, 5, 0, 0, 0},
		{5, 0, 0, 0, 3, 0, 0, 2, 7, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	// grille = make([][]int, TAILLE+2)
	// for i := 0; i < TAILLE+2; i++ {
	// 	grille[i] = make([]int, TAILLE+1)
	// }
	// for i := 0; i < TAILLE+2; i++ {
	// 	fmt.Println(grille[i])
	// }

	maj_compteurs(grille)
	fmt.Print("heyyy")

}

// github pour faire du sdl en go veandco go sdl 2
// demander pour le "go run UTILS.GO MAIN.GO"
// et l'histoire des packages
