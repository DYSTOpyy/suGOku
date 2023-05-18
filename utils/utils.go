package utils

import (
	"math"
	"fmt"
	"strconv"
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

// print_grille permet d'afficher une grille de sudoku ainsi que ses compteurs.
// Si string_output=true, alors la grille n'est pas affichée mais renvoyée en string.
//
// La grille en argument doit être sous forme de pointeur.
func Print_grille(grille *[TAILLE + 2][TAILLE + 1]int,size int ,string_output bool) string {

	text := ""
	for line := 0; line < size; line++ {

		// séparateur de bloc horizontale
		if line%int(math.Sqrt(float64(size))) == 0 && line != 0 {
			for i := 0; i < (int(math.Sqrt(float64(size)))+1)*int(math.Sqrt(float64(size))); i++ {
				text = text + "-"
			}
			text = text + "\n"
		}

		for column := 0; column < size; column++ {

			// séparateur de bloc vertical
			if column%int(math.Sqrt(float64(size))) == 0 && column != 0 {
				text = text + "|"
			}

			if grille[line][column] == 0 {
				text = text + "."

			} else {

				// conversion en ASCII pour les valeurs alphabétiques
				// note : rune() utilisé à la place de string() car sinon incompatible avec %c
				if grille[line][column] > 9 {
					text = text + string(grille[line][column]-10+65)
				} else {
					text = text + string(grille[line][column]+48)
				}
			}
		}

		// Compteur de ligne
		text = text + " (" + strconv.Itoa(grille[line][size]) + ")\n"

	}

	text = text + "\n"

	// compteur de colonne + bloc
	for line := size; line < size+2; line++ {
		for column := 0; column < size; column++ {

			if column%int(math.Sqrt(float64(size))) == 0 && column != 0 {

				text += "|"
			}
			text += strconv.Itoa(grille[line][column])
		}
		text += "\n"
	}

	text += "Nombre d'indices total :" + strconv.Itoa(grille[size][size])

	if string_output {
		return text
	} else {
		fmt.Println(text)
		return ""
	}
}

// generer_possibilite calcule les valeurs possibles pour chaque case vide
// d'une grille qu'il place dans un slice de dimension TAILLE:TAILLE (donc pas la même dimension que la grille !)
// qui EST RENVOYÉ où chaque case contient un slice représentant les valeurs possibles (vide si la case est occupée).
//
// La grille en argument doit être sous forme de pointeur.
func Generer_possibilite(grille *[TAILLE + 2][TAILLE + 1]int, size int) [TAILLE][TAILLE][]int {

	possibilite := [TAILLE][TAILLE][]int{}

	// pour chaque case
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {

			c := []int{}

			for value := 1; value < size+1; value++ {

				if grille[line][column] == 0 && IsOkayCase(line, column, *grille, size, value) {

					// si la valeur est possible, on l'ajoute au slice des possibilitées
					c = append(c, value)

				}
			}
			possibilite[line][column] = c
		}
	}
	return possibilite
}

// generer_masque crée un slice de boolean qui indique si la case était là au départ ou non.
// si une case vaut vrai, cela signifie que la valeur correspondante dans la grille était déjà là.
// si c'est faux, cela signifie que la case était vide au départ et que son contenu est choisi par le joueur.
//
// La grille en argument doit être sous forme de pointeur.
func Generer_masque (grille *[TAILLE + 2][TAILLE + 1]int,size int) [TAILLE][TAILLE]bool {

	masque := [TAILLE][TAILLE]bool{}

	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			
			if grille[line][column] != 0 {
				masque[line][column] = true
			} else {
				masque[line][column] = false
			}
		}
	}
	return masque
}

