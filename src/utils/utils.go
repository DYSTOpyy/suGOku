package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"golang.org/x/exp/slices"
)

var  Taille int32 = 9

const MAX int32 = 16

func Is_In(valeur int, slice []int) bool {

	for _, v := range slice {
		if valeur == v {
			return true
		}
	}
	return false

}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur valeur, et renvoie un boolean
func CaseBonne(x int32, y int32, grille [MAX + 2][MAX + 1]int, valeur int) bool {
	ligne := LigneBonne(x, y, grille, valeur)
	colonne := ColonneBonne(x, y, grille, valeur)
	boite := BoiteBonne(x, y, grille, valeur)
	return ligne && colonne && boite
}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur valeur sur la ligne, et renvoie un boolean
func LigneBonne(x int32, y int32, grille [MAX + 2][MAX + 1]int, valeur int) bool {
	var i int32 = 0
	for i < Taille && grille[x][i] != valeur {
		i++
	}
	return i == Taille
}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur valeur sur la colonne, et renvoie un boolean
func ColonneBonne(x int32, y int32, grille [MAX + 2][MAX + 1]int, valeur int) bool {
	var i int32 = 0
	for i < Taille && grille[i][y] != valeur {
		i++
	}
	return i == Taille
}

// Verifie si la case (x,y) de la grille en paramètre peut acceuillir la valeur valeur sur la boite, et renvoie un boolean
func BoiteBonne(x int32, y int32, grille [MAX + 2][MAX + 1]int, valeur int) bool {
	racine := int32(math.Sqrt(float64(Taille)))
	var i int32 = 0
	for i < Taille && grille[int32(x/racine)*racine+int32(i/racine)][int32(y/racine)*racine+i%racine] != valeur {
		i++
	}
	return i == Taille
}

// maj_compteurs met à jour les compteurs du nombre d'indices par ligne/colonne/bloc
// d'une grille qui sont situés respectivement sur la colonne MAX,
// la ligne MAX et la ligne MAX+1, ainsi que le nombre total d'indice
// de la grille situé en MAX:MAX.
//
// La grille en argument doit être sous forme de pointeur.
func Maj_compteurs(grille *[MAX + 2][MAX + 1]int) {
	// Reinitialisation à 0
	for index := int32(0); index < Taille; index++ {
		grille[index][Taille] = 0 // maj des lignes
		grille[Taille][index] = 0
		grille[Taille+1][index] = 0
	}
	grille[Taille][Taille] = 0
	// Remplissage
	for ligne := int32(0); ligne < Taille; ligne++ {
		for colonne := int32(0); colonne < Taille; colonne++ {
			if grille[ligne][colonne] != 0 {
				grille[ligne][Taille] += 1                                                                                                        // maj des lignes
				grille[Taille][colonne] += 1                                                                                                      // maj des colonnes
				grille[Taille+1][(ligne / int32(math.Sqrt(float64(Taille)))) * int32(math.Sqrt(float64(Taille))) + (colonne/int32(math.Sqrt(float64(Taille))))] += 1 // maj des bloc
				grille[Taille][Taille] += 1                                                                                                        // maj du nombre total d'indice
			}
		}
	}
}

// print_grille permet d'afficher une grille de sudoku ainsi que ses compteurs.
//
// La grille en argument doit être sous forme de pointeur.
func Afficher_grille(grille *[MAX + 2][MAX + 1]int) {

	text := ""
	for ligne := int32(0); ligne < Taille; ligne++ {

		// séparateur de bloc horizontale
		if ligne%int32(math.Sqrt(float64(Taille))) == 0 && ligne != 0 {
			for i := 0; i < (int(math.Sqrt(float64(Taille)))+1)*int(math.Sqrt(float64(Taille))); i++ {
				text = text + "-"
			}
			text = text + "\n"
		}

		for colonne := int32(0); colonne < Taille; colonne++ {

			// séparateur de bloc vertical
			if colonne%int32(math.Sqrt(float64(Taille))) == 0 && colonne != 0 {
				text = text + "|"
			}

			if grille[ligne][colonne] == 0 {
				text = text + "."

			} else {

				// conversion en ASCII pour les valeurs alphabétiques
				// note : rune() utilisé à la place de string() car sinon incompatible avec %c
				if grille[ligne][colonne] > 9 {
					text = text + string(grille[ligne][colonne]-10+65)
				} else {
					text = text + string(grille[ligne][colonne]+48)
				}
			}
		}

		// Compteur de ligne
		text = text + " (" + strconv.Itoa(grille[ligne][Taille]) + ")\n"

	}

	text = text + "\n"

	// compteur de colonne + bloc
	for ligne := Taille; ligne < Taille+2; ligne++ {
		for colonne := int32(0); colonne < Taille; colonne++ {

			if colonne%int32(math.Sqrt(float64(Taille))) == 0 && colonne != 0 {

				text += "|"
			}
			text += strconv.Itoa(grille[ligne][colonne])
		}
		text += "\n"
	}

	text += "Nombre d'indices total :" + strconv.Itoa(grille[Taille][Taille])

	fmt.Println(text)

}

// Permet d'afficher de façon propre les tableaux de booléans
func Afficher_boolTable(grille *[MAX][MAX]bool) {

	text := ""
	for ligne := int32(0); ligne < Taille; ligne++ {

		// séparateur de bloc horizontale
		if ligne % int32(math.Sqrt(float64(Taille))) == 0 && ligne != 0 {
			for i := 0; i < (int(math.Sqrt(float64(Taille))) + 1) * int(math.Sqrt(float64(Taille))); i++ {
				text = text + "-"
			}
			text = text + "\n"
		}

		for colonne := int32(0); colonne < Taille; colonne++ {

			// séparateur de bloc vertical
			if colonne % int32(math.Sqrt(float64(Taille))) == 0 && colonne != 0 {
				text = text + "|"
			}

			if grille[ligne][colonne] {
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
func Somme_BoolTable(grille *[MAX][MAX]bool) int32 {
	var count int32 = 0
	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			if grille[i][j] {
				count += 1
			}
		}
	}
	return count
}

// generer_possibilite calcule les valeurs possibles pour chaque case vide
// d'une grille qu'il place dans un slice de dimension MAX:MAX (donc pas la même dimension que la grille !)
// qui EST RENVOYÉ où chaque case contient un slice représentant les valeurs possibles (vide si la case est occupée).
//
// La grille en argument doit être sous forme de pointeur.
func Generer_possibilite(grille *[MAX + 2][MAX + 1]int) [MAX][MAX][]int {

	possibilite := [MAX][MAX][]int{}

	// pour chaque case
	for ligne := int32(0); ligne < Taille; ligne++ {
		for colonne := int32(0); colonne < Taille; colonne++ {

			c := []int{}

			for valeur := 1; valeur < int(Taille+1); valeur++ {

				if grille[ligne][colonne] == 0 && CaseBonne(ligne, colonne, *grille, valeur) {

					// si la valeur est possible, on l'ajoute au slice des possibilitées
					c = append(c, valeur)

				}
			}
			possibilite[ligne][colonne] = c
		}
	}
	return possibilite
}

// generer_masque crée un slice de boolean qui indique si la case était là au départ ou non.
// si une case vaut vrai, cela signifie que la valeur correspondante dans la grille était déjà là.
// si c'est faux, cela signifie que la case était vide au départ et que son contenu est choisi par le joueur.
//
// La grille en argument doit être sous forme de pointeur.
func Generer_masque(grille *[MAX + 2][MAX + 1]int) [MAX][MAX]bool {

	masque := [MAX][MAX]bool{}

	for ligne := int32(0); ligne < Taille; ligne++ {
		for colonne := int32(0); colonne < Taille; colonne++ {

			if grille[ligne][colonne] != 0 {
				masque[ligne][colonne] = true
			} else {
				masque[ligne][colonne] = false
			}
		}
	}
	return masque
}

// Permet d'ordonner une slice de façon aléatoire
func ListRandomize(liste []int32) []int32 {
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
func IndexToLinCol(index int32) (int32, int32) {
	ligne, colonne := int32(index/Taille), index%Taille
	return ligne, colonne
}

// Renvoie une grille avec les erreurs présente sur le plateau, sous la forme d'un tableau de boolean qui prends pour valeur true si la case contient une erreur, false sinon
func TrouverErreurs(grille *[MAX + 2][MAX + 1]int, mask *[MAX][MAX]bool) [MAX][MAX]bool {
	errors := [MAX][MAX]bool{}
	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			if !mask[i][j] && grille[i][j] != 0 {
				valeur := grille[i][j]
				grille[i][j] = 0
				errors[i][j] = !CaseBonne(i, j, *grille, valeur)
				grille[i][j] = valeur
			}
		}
	}
	return errors
}

func RestartGrille(grille *[MAX + 2][MAX + 1]int, mask *[MAX][MAX]bool) {
	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			if !mask[i][j] {
				grille[i][j] = 0
			}
		}
	}
}

// détermine si un tableau 2D booleen est vide (chaque case est false) -> renvoi true
// sinon (si y a au moins 1 case true) -> renvoi false
func EmptyBoolArray(array *[MAX][MAX]bool) bool {
	for _, ligne := range array {
		for _, valeur := range ligne {
			if valeur {
				return false
			}
		}
	}
	return true
}

// détermine si un tableau 2D d'entiers est rempli (aucune case ne vaut 0) -> renvoi true
// sinon (si y a au moins une case avec 0) -> renvoi false
func FullIntArray(grille *[MAX + 2][MAX + 1]int) bool {

	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			if grille[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
