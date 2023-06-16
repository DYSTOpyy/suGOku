package utils

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// Renvoie le chemin absolu du dossier racine du projet
func GetPackagePath() string {
	Path, _ := filepath.Abs("")
	for filepath.Base(Path) != "src" {
		Path, _ = filepath.Split(Path)
		Path = filepath.Clean(Path)
	}
	return Path
}

// Enregistre une grille dans un fichier save.txt, présent dans le dossier files, sous la forme de 2 blocs de chiffres :
// le premier pour la grille
// et le second pour le masque
// Ce programme peut renvoyer une erreur de type PathError ou une erreur d'écriture
func SaveFile(grille *[MAX + 2][MAX + 1]int, mask [MAX][MAX]bool) error {

	file, err := os.Create("files/save.txt")
	if err != nil {
		return err
	}
	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			char := strconv.FormatInt(int64(grille[i][j]), 17)
			if char == "0" {
				char = "."
			}
			_, err2 := file.WriteString(char)
			if err2 != nil {
				file.Close()
				return err2
			}
		}
	}
	_, err2 := file.WriteString(" ")
	if err2 != nil {
		file.Close()
		return err2
	}
	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			char := "0"
			if mask[i][j] {
				char = "1"
			}
			_, err2 := file.WriteString(char)
			if err2 != nil {
				file.Close()
				return err2
			}
		}
	}
	return nil
}

// Permet d'importer un sudoku totalemant neuf, ou de charger une sauvegarde
// Le programme renvoie la grille, le masque associé ainsi qu'une erreur
// Peut renvoyer une erreur d'ouverture, de taille ou de conversion
func ImportFile() ([MAX + 2][MAX + 1]int, [MAX][MAX]bool, error) {
	grille := [MAX + 2][MAX + 1]int{}
	mask := [MAX][MAX]bool{}
	buffer, err := os.ReadFile("files/save.txt")
	if err != nil {
		return grille, mask, err
	}
	var nbChar int32 = int32(len(buffer))
	var root int32
	var isSave bool
	if root = int32(math.Sqrt(float64(nbChar))); root*root == nbChar {
		isSave = false
		Taille = root
	} else if root = int32(math.Sqrt(float64((nbChar - 1) / 2))); root*root*2+1 == nbChar {
		Taille = root
		isSave = true
	} else {
		return grille, mask, errors.New("TailleError")
	}

	for i := int32(0); i < Taille; i++ {
		for j := int32(0); j < Taille; j++ {
			char := string(buffer[i*Taille+j])
			if char == "." {
				char = "0"
			}
			Case, err := strconv.ParseInt(char, 17, 0)
			if err != nil {
				return grille, mask, err
			}
			grille[i][j] = int(Case)
		}
	}

	if isSave {
		for i := int32(0); i < Taille; i++ {
			for j := int32(0); j < Taille; j++ {
				char := string(buffer[(Taille*Taille+1)+i*Taille+j])
				if char == "0" {
					mask[i][j] = false
				} else {
					mask[i][j] = true
				}
			}
		}
	} else {
		mask = Generer_masque(&grille)
	}
	Maj_compteurs(&grille)
	return grille, mask, nil
}
