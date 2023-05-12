package main

import (
	"fmt"
	"math"
	"time"
)

const TAILLE int = 9

func maj_compteurs(grille *[TAILLE + 2][TAILLE + 1]int) {
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

func remplir_possibilite(grille [TAILLE + 2][TAILLE + 1]int) [TAILLE][TAILLE][]int {

	possibilite := [TAILLE][TAILLE][]int{}

	for line := 0; line < TAILLE; line++ {

		for column := 0; column < TAILLE; column++ {

			c := []int{}

			for value := 1; value < TAILLE+1; value++ {

				if IsOkayCase(line, column, grille, TAILLE, value) {

					c = append(c, value)

				}

			}

			possibilite[line][column] = c

		}

	}

	return possibilite
}

func print_grille(grille [TAILLE + 2][TAILLE + 1]int) {

	for line := 0; line < TAILLE; line++ {

		if line%int(math.Sqrt(float64(TAILLE))) == 0 && line != 0 {
			fmt.Println("-----------")
		}
		for column := 0; column < TAILLE; column++ {

			if column%int(math.Sqrt(float64(TAILLE))) == 0 && column != 0 {
				fmt.Print("|")
			}

			if grille[line][column] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(grille[line][column])
			}
		}

		fmt.Println(" (", grille[line][TAILLE], ")")

	}

}

func algo_backtracking(grille *[TAILLE + 2][TAILLE + 1]int, ligne int, col int, possibilite *[TAILLE][TAILLE][]int) bool {

	if ligne >= TAILLE || col >= TAILLE {

		return true
	} else if grille[ligne][col] != 0 {

		return algo_backtracking(grille, ligne+(col+1)/TAILLE, (col+1)%TAILLE, possibilite)

	} else {

		for _, value := range possibilite[ligne][col] {

			if IsOkayCase(ligne, col, *grille, TAILLE, value) {

				grille[ligne][col] = value

				if algo_backtracking(grille, ligne+(col+1)/TAILLE, (col+1)%TAILLE, possibilite) {

					return true

				}

			}

		}

		grille[ligne][col] = 0
		return false
	}
}

func algo_backtracking_experimental(grille *[TAILLE + 2][TAILLE + 1]int, ligne int, col int, possibilite [TAILLE][TAILLE][]int) bool {
	// la difference c'est que je redefinie les possibilités à chaque valeur testée
	// tandis que dans l'autre, je le fait qu'au début afin d'avoir moins de valeur à tester plutot que toutes les valeurs de 1 à 9
	// pour y arriver, j'exploite le fait que lorsqu'un slice est passé en argument, la fonction en crée en fait une copie locale plutot que de travailler sur l'original
	if ligne >= TAILLE || col >= TAILLE {

		return true
	} else if grille[ligne][col] != 0 {

		return algo_backtracking_experimental(grille, ligne+(col+1)/TAILLE, (col+1)%TAILLE, possibilite)

	} else {

		for _, value := range possibilite[ligne][col] {

			if IsOkayCase(ligne, col, *grille, TAILLE, value) {

				grille[ligne][col] = value

				if algo_backtracking_experimental(grille, ligne+(col+1)/TAILLE, (col+1)%TAILLE, remplir_possibilite(*grille)) {

					return true

				}

			}

		}

		grille[ligne][col] = 0
		return false
	}
}

func main() {

	// p := fmt.Println

	grille := [TAILLE + 2][TAILLE + 1]int{{0, 0, 0, 0, 4, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 0, 0},
		{0, 3, 2, 1, 9, 0, 0, 8, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 4, 0},
		{0, 0, 4, 9, 0, 0, 0, 6, 1, 0},
		{0, 2, 9, 0, 7, 4, 0, 0, 0, 0},
		{0, 9, 6, 7, 0, 1, 5, 0, 0, 0},
		{5, 0, 0, 0, 3, 0, 0, 2, 7, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	maj_compteurs(&grille)
	print_grille(grille)

	possibilite := remplir_possibilite(grille)

	before := time.Now()
	// algo_backtracking(&grille, 0, 0, &possibilite)
	algo_backtracking_experimental(&grille, 0, 0, possibilite) // PREND PLUS DE TEMPS (logique, mais peut être moins que l'autre pour du 16x16 ?)
	after := time.Now()
	diff := after.Sub(before)
	fmt.Println(diff.Milliseconds())
	print_grille(grille)

}

// github pour faire du sdl en go veandco go sdl 2
// demander pour le "go run UTILS.GO MAIN.GO"
// et l'histoire des packages
