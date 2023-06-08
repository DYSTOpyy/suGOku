package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"golang.org/x/exp/slices"
)

var Size = 9

const TAILLE int = 16

func Is_In(value int, slice []int) bool {

	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false

}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur value, et renvoie un boolean
func IsOkayCase(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, value int) bool {
	line := IsOkayLine(x, y, grid, value)
	column := IsOkayColumn(x, y, grid, value)
	box := IsOkayBox(x, y, grid, value)
	return line && column && box
}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur value sur la ligne, et renvoie un boolean
func IsOkayLine(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, value int) bool {
	i := 0
	for i < Size && grid[x][i] != value {
		i++
	}
	return i == Size
}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur value sur la colonne, et renvoie un boolean
func IsOkayColumn(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, value int) bool {
	i := 0
	for i < Size && grid[i][y] != value {
		i++
	}
	return i == Size
}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur value sur la boite, et renvoie un boolean
func IsOkayBox(x int, y int, grid [TAILLE + 2][TAILLE + 1]int, value int) bool {
	root := int(math.Sqrt(float64(Size)))
	i := 0
	for i < Size && grid[int(x/root)*root+int(i/root)][int(y/root)*root+i%root] != value {
		i++
	}
	return i == Size
}

// maj_compteurs met à jour les compteurs du nombre d'indices par ligne/colonne/bloc
// d'une grille qui sont situés respectivement sur la colonne TAILLE,
// la ligne TAILLE et la ligne TAILLE+1, ainsi que le nombre total d'indice
// de la grille situé en TAILLE:TAILLE.
//
// La grille en argument doit être sous forme de pointeur.
func Maj_compteurs(grille *[TAILLE + 2][TAILLE + 1]int) {
	// Reinitialisation à 0
	for index := 0; index < Size; index++ {
		grille[index][Size] = 0 // maj des lignes
		grille[Size][index] = 0
		grille[Size+1][index] = 0
	}
	grille[Size][Size] = 0
	// Remplissage
	for line := 0; line < Size; line++ {
		for column := 0; column < Size; column++ {
			if grille[line][column] != 0 {
				grille[line][Size] += 1                                                                                                        // maj des lignes
				grille[Size][column] += 1                                                                                                      // maj des colonnes
				grille[Size+1][(line/int(math.Sqrt(float64(Size))))*int(math.Sqrt(float64(Size)))+(column/int(math.Sqrt(float64(Size))))] += 1 // maj des bloc
				grille[Size][Size] += 1                                                                                                        // maj du nombre total d'indice
			}
		}
	}
}

// print_grille permet d'afficher une grille de sudoku ainsi que ses compteurs.
//
// La grille en argument doit être sous forme de pointeur.
func Print_grille(grille *[TAILLE + 2][TAILLE + 1]int) {

	text := ""
	for line := 0; line < Size; line++ {

		// séparateur de bloc horizontale
		if line%int(math.Sqrt(float64(Size))) == 0 && line != 0 {
			for i := 0; i < (int(math.Sqrt(float64(Size)))+1)*int(math.Sqrt(float64(Size))); i++ {
				text = text + "-"
			}
			text = text + "\n"
		}

		for column := 0; column < Size; column++ {

			// séparateur de bloc vertical
			if column%int(math.Sqrt(float64(Size))) == 0 && column != 0 {
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
		text = text + " (" + strconv.Itoa(grille[line][Size]) + ")\n"

	}

	text = text + "\n"

	// compteur de colonne + bloc
	for line := Size; line < Size+2; line++ {
		for column := 0; column < Size; column++ {

			if column%int(math.Sqrt(float64(Size))) == 0 && column != 0 {

				text += "|"
			}
			text += strconv.Itoa(grille[line][column])
		}
		text += "\n"
	}

	text += "Nombre d'indices total :" + strconv.Itoa(grille[Size][Size])

	fmt.Println(text)

}

// Permet d'afficher de façon propre les tableaux de booléans
func Print_boolTable(grille *[TAILLE][TAILLE]bool) {

	text := ""
	for line := 0; line < Size; line++ {

		// séparateur de bloc horizontale
		if line%int(math.Sqrt(float64(Size))) == 0 && line != 0 {
			for i := 0; i < (int(math.Sqrt(float64(Size)))+1)*int(math.Sqrt(float64(Size))); i++ {
				text = text + "-"
			}
			text = text + "\n"
		}

		for column := 0; column < Size; column++ {

			// séparateur de bloc vertical
			if column%int(math.Sqrt(float64(Size))) == 0 && column != 0 {
				text = text + "|"
			}

			if grille[line][column] {
				text = text + "1"
			} else {
				text = text + "0"
			}

		}
		text = text + "\n"
	}
	fmt.Println(text)
}

// Permet de renvoyer la nombre de fois où une case est marqué vraie dans un tableau
func Sum_BoolTable(grille *[TAILLE][TAILLE]bool) int {
	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if grille[i][j] {
				count += 1
			}
		}
	}
	return count
}

// generer_possibilite calcule les valeurs possibles pour chaque case vide
// d'une grille qu'il place dans un slice de dimension TAILLE:TAILLE (donc pas la même dimension que la grille !)
// qui EST RENVOYÉ où chaque case contient un slice représentant les valeurs possibles (vide si la case est occupée).
//
// La grille en argument doit être sous forme de pointeur.
func Generer_possibilite(grille *[TAILLE + 2][TAILLE + 1]int) [TAILLE][TAILLE][]int {

	possibilite := [TAILLE][TAILLE][]int{}

	// pour chaque case
	for line := 0; line < Size; line++ {
		for column := 0; column < Size; column++ {

			c := []int{}

			for value := 1; value < Size+1; value++ {

				if grille[line][column] == 0 && IsOkayCase(line, column, *grille, value) {

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
func Generer_masque(grille *[TAILLE + 2][TAILLE + 1]int) [TAILLE][TAILLE]bool {

	masque := [TAILLE][TAILLE]bool{}

	for line := 0; line < Size; line++ {
		for column := 0; column < Size; column++ {

			if grille[line][column] != 0 {
				masque[line][column] = true
			} else {
				masque[line][column] = false
			}
		}
	}
	return masque
}

// Permet d'ordonner une slice de façon aléatoire
func ListRandomize(liste []int) []int {
	length := len(liste)
	nb_change := rand.Int()%int(length/2) + int(length/2)
	for i := 0; i < nb_change; i++ {
		index := rand.Int() % length
		number := liste[index]
		liste = slices.Insert(liste, rand.Int()%(length-1), number)
	}
	return liste
}

// Permet de convertir l'indice d'une case en son numéro de ligne et celui de colonne
func IndexToLinCol(index int) (int, int) {
	line, column := int(index/Size), index%Size
	return line, column
}

// Renvoie une grille avec les erreurs présente sur le plateau, sous la forme d'un tableau de boolean qui prends pour valeur true si la case contient une erreur, false sinon
func FindErrors(grid *[TAILLE + 2][TAILLE + 1]int, mask *[TAILLE][TAILLE]bool) [TAILLE][TAILLE]bool {
	errors := [TAILLE][TAILLE]bool{}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if !mask[i][j] && grid[i][j] != 0 {
				value := grid[i][j]
				grid[i][j] = 0
				errors[i][j] = !IsOkayCase(i, j, *grid, value)
				grid[i][j] = value
			}
		}
	}
	return errors
}

func RestartGrille(grid *[TAILLE + 2][TAILLE + 1]int, mask *[TAILLE][TAILLE]bool) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if !mask[i][j] {
				grid[i][j] = 0
			}
		}
	}
}

// détermine si un tableau 2D booleen est vide (chaque case est false) -> renvoi true
// sinon (si y a au moins 1 case true) -> renvoi false
func EmptyBoolArray(array *[TAILLE][TAILLE]bool) bool {
	for _, line := range array {
		for _, value := range line {
			if value {
				return false
			}
		}
	}
	return true
}

// détermine si un tableau 2D d'entiers est rempli (aucune case ne vaut 0) -> renvoi true
// sinon (si y a au moins une case avec 0) -> renvoi false
func FullIntArray(grid *[TAILLE + 2][TAILLE + 1]int) bool {

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
