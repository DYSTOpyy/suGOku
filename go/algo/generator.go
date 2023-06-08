package algo

import (
	"math"
	"math/rand"

	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)

type Difficulty int

const (
	Easy   Difficulty = 1
	Medium Difficulty = 2
	Hard   Difficulty = 3
	Evil   Difficulty = 4
)

const TAILLE int = utils.TAILLE

// Retourne la fonction de Digging, le nombre de cases données et le minimum par ligne
func EditDifficulty(diff Difficulty) (func(*[TAILLE][TAILLE]bool, int, int) (int, int), int, int) {
	var (
		nb_case  int
		min_line int
		algoDig  func(*[TAILLE][TAILLE]bool, int, int) (int, int)
	)
	switch diff {
	case Easy:
		nb_case = rand.Int()%14 + 36
		min_line = 4
		algoDig = RandomDig
	case Medium:
		nb_case = rand.Int()%3 + 32
		min_line = 3
		algoDig = JumpingDig
	case Hard:
		nb_case = rand.Int()%3 + 28
		min_line = 2
		algoDig = WanderingDig
	case Evil:
		nb_case = rand.Int()%5 + 22
		min_line = 0
		algoDig = TopBottomDig
	}
	return algoDig, nb_case, min_line
}

// Géneration complète d'une grille en fonction d'un paramètre de difficulté
func GeneratorTotal(diff Difficulty) *[TAILLE + 2][TAILLE + 1]int {
	grille := [TAILLE + 2][TAILLE + 1]int{}
	GenGridFull(&grille)
	liste := GenSlice(&grille)
	possibilite := utils.Generer_possibilite(&grille)
	Algo_backtracking(&grille, &possibilite, liste)
	utils.Maj_compteurs(&grille)
	DiggingHoles(&grille, diff)
	return &grille
}

// init_grille génère une grille puis détermine ses compteurs et ses possibilité.
// C'est une fonction à retour multiple, qui renvoie la grille et les possibilités.
//
// La grille en argument doit être sous forme de pointeur.		DOIT DISPARAITRE POUR LAISSER PLACE A UN ARGUMENT "DIFFICULTE"
func Init_grille(difficulteChoisie Difficulty) ([TAILLE + 2][TAILLE + 1]int, [TAILLE][TAILLE][]int, [TAILLE][TAILLE]bool, [TAILLE][TAILLE]bool) {
	// générer la grille
	grille := GeneratorTotal(difficulteChoisie)
	// mettre à jour ses compteurs
	utils.Maj_compteurs(grille)
	// générer la grille de possibilité
	possibilite := utils.Generer_possibilite(grille)
	// generer le masque de la grille
	masque := utils.Generer_masque(grille)
	// generer la grille des cases vérifiées
	verifier := [TAILLE][TAILLE]bool{}
	// renvoyer les deux
	return *grille, possibilite, masque, verifier
}

// Permet de créer une liste d'entier allant de 0 à la taille de la grille au carré, en enlevant les indices des cases déja occupées
func GenSlice(grille *[TAILLE + 2][TAILLE + 1]int) []int {
	size := utils.Size
	liste := []int{}
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			if grille[line][column] == 0 {
				liste = append(liste, size*line+column)
			}
		}
	}
	return liste
}

// Initialise la grille en remplissant les carrés présent sur la diagonale de la grille avec des valeurs aléatoires
func GenGridFull(grille *[TAILLE + 2][TAILLE + 1]int) {
	size := utils.Size
	possibilite := utils.Generer_possibilite(grille)
	liste := []int{}
	root := int(math.Sqrt(float64(size)))
	for i := 0; i < size; i++ {
		for j := 0; j < size; j = j + root {
			liste = append(liste, i%root+int(i/root)*size+j*(size)+j)
		}
	}
	liste = utils.ListRandomize(liste)
	Algo_backtracking(grille, &possibilite, liste)
}

// Algorythme Top-Down qui permet de vider une grille en fonction de la difficulté choisie
func DiggingHoles(grille *[TAILLE + 2][TAILLE + 1]int, diff Difficulty) {
	AlgoDig, nb_case, min_line := EditDifficulty(diff)
	digTable := FillDiggable()
	nbRemovedCase := 1
	linToDig, colToDig := rand.Int()%utils.Size, rand.Int()%utils.Size
	Dig(grille, digTable, linToDig, colToDig)
	for nbRemovedCase < utils.Size*utils.Size-nb_case && ContinueDigging(digTable) {
		newLinToDig, newColToDig := AlgoDig(digTable, linToDig, colToDig)
		if grille[newLinToDig][utils.Size]-1 < min_line || grille[utils.Size][newColToDig]-1 < min_line {
			digTable[newLinToDig][newColToDig] = false
			continue
		}
		Copy := *grille
		if IsSolutionUnique(Copy, newLinToDig, newColToDig) {
			linToDig, colToDig = newLinToDig, newColToDig
			Dig(grille, digTable, linToDig, colToDig)
			nbRemovedCase += 1

		} else {
			digTable[newLinToDig][newColToDig] = false
		}
	}
}

// Permet d'initaliser la DigTable avec true dans toutes ses cases
func FillDiggable() *[TAILLE][TAILLE]bool {
	digTable := &[TAILLE][TAILLE]bool{}
	for i := 0; i < utils.Size; i++ {
		for j := 0; j < utils.Size; j++ {
			digTable[i][j] = true
		}
	}
	return digTable
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par hasard
func RandomDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew := rand.Int() % utils.Size
	colNew := rand.Int() % utils.Size
	for !digTable[linNew][colNew] {
		linNew = rand.Int() % utils.Size
		colNew = rand.Int() % utils.Size
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Jumping One
func JumpingDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew, colNew := linPrev, colPrev
	ToLeft := linNew % 2
	for !digTable[linNew][colNew] {
		if ToLeft == 1 {
			colNew -= 2
			if colNew < 0 {
				colNew += 1 + 2*((colNew+3)%2)
				linNew += 1
				if linNew >= utils.Size {
					linNew = 0
					colNew = 0 + (colNew % 2)
				}
			}
		} else {
			colNew += 2
			if colNew >= utils.Size {
				colNew -= 1 + 2*((colNew+1)%2)
				linNew += 1
				if linNew >= utils.Size {
					linNew = 0
					colNew = 0 + (colNew % 2)
				}
			}
		}
		ToLeft = linNew % 2
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Wandering along S
func WanderingDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew, colNew := linPrev, colPrev
	ToLeft := linNew % 2
	for !digTable[linNew][colNew] {
		if ToLeft == 1 {
			colNew -= 1
			if colNew < 0 {
				colNew = 0
				linNew += 1
			}
		} else {
			colNew += 1
			if colNew >= utils.Size {
				colNew = utils.Size - 1
				linNew += 1
			}
		}
		if linNew >= utils.Size {
			linNew = 0
			colNew = 0
		}
		ToLeft = linNew % 2
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Left to Right then Top to Bottom
func TopBottomDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew, colNew := linPrev, colPrev
	for !digTable[linNew][colNew] {
		colNew += 1
		if colNew >= utils.Size {
			colNew -= utils.Size
			linNew += 1
			if linNew >= utils.Size {
				linNew = 0
			}
		}
	}
	return linNew, colNew
}

// Verifie l'unicité d'une solution et renvoie true si la solution est unique, false sinon
func IsSolutionUnique(grille [TAILLE + 2][TAILLE + 1]int, linToDig int, colToDig int) bool {
	listValue := []int{}
	for i := 1; i < utils.Size+1; i++ {
		if i != grille[linToDig][colToDig] {
			listValue = append(listValue, i)
		}
	}
	for _, value := range listValue {
		if utils.IsOkayCase(linToDig, colToDig, grille, value) {
			grille[linToDig][colToDig] = value
			possibilite := utils.Generer_possibilite(&grille)
			slice := GenSlice(&grille)
			if Algo_backtracking(&grille, &possibilite, slice) {
				return false
			}
		}
	}
	return true
}

// Permet de miner la case indiqué, en mettant à jour la grille et la DigTable
func Dig(grille *[TAILLE + 2][TAILLE + 1]int, digTable *[TAILLE][TAILLE]bool, linToDig int, colToDig int) {
	digTable[linToDig][colToDig] = false
	grille[linToDig][colToDig] = 0
	utils.Maj_compteurs(grille)
}

// Fonction qui renvoie un boolean indiquant si il reste encore des cases à miner dans notre grille
func ContinueDigging(digTable *[TAILLE][TAILLE]bool) bool {
	for i := 0; i < utils.Size; i++ {
		for j := 0; j < utils.Size; j++ {
			if digTable[i][j] {
				return true
			}
		}
	}
	return false
}
