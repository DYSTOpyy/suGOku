import numpy as np
from math import *

taille = 9

def afficher (grille, plan) :   # afficher comme une grille un des plan de la matrice (0:sudoku, 123...:possibilités)
    
    print("Affichage du plan n°",plan)
    # pour chaque ligne 
    for i in range (taille) :

        if (i%int(sqrt(taille)) == 0 and i != 0) :
            print("-----------")

        # pour chaque colonne
        for j in range (taille) :

            if (j%int(sqrt(taille)) == 0 and j != 0) :
                print("|", end="")

            if (grille[plan][i][j] == 0) :
                print(".", end="")
            else :
                print(grille[plan][i][j], end="")

        print(" (",grille[plan][i][taille],")", end="")
        print()

    print()

    for i in range (taille, taille+2) :

        for j in range (taille) :

            if (j%int(sqrt(taille)) == 0 and j != 0) :
                print("|", end="")

            print(grille[plan][i][j], end="")

        print()
    print()

def majCompteurs (grille) :         # met a jour les compteurs de nb de cases par ligne/colonne/carré
    for i in range (taille) :
        for j in range (taille) :

            if (grille[0][i][j] != 0) :
                grille[0][i][taille] += 1                       # maj des lignes 
                grille[0][taille][j] += 1                       # maj des colonnes 
                grille[0][taille+1][(i//int(sqrt(taille)))*int(sqrt(taille))+(j//int(sqrt(taille)))] += 1       # maj des carrés

def isOkay (grille, valeur, ligne, col) :           # vérifie si une valeur est possible à un endroit ligne/col donné dans la grille de jeu

    for i in range (taille) :
        if (grille[0][ligne][i] == valeur or grille[0][i][col] == valeur or grille[0][(ligne//int(sqrt(taille)))*int(sqrt(taille))+i//int(sqrt(taille))][(col//int(sqrt(taille)))*int(sqrt(taille))+i%int(sqrt(taille))] == valeur) :
            return False 
        
    return True

def remplirPossibilite (grille) :               # rempli toute la matrice 3D de sorte a avoir des 1 sur le plan X si X est possible d'être mis à cette case

    for i in range (taille) :

        for j in range (taille) :

            if (grille[0][i][j] == 0) :

                for valeur in range (1,taille+1) :

                    if (isOkay(grille, valeur, i, j)) :

                        grille[valeur][i][j] = 1

def algoBacktracking (grille, ligne, col) :         # resout avec le backtracking mais sans afficher si solution unique ou pas

    if (ligne >= taille or col >= taille) :

        return True
    
    elif (grille[0][ligne][col] != 0) :

        return algoBacktracking(grille, ligne+(col+1)//taille, (col+1)%taille)
    
    else :

        for i in range (1,taille+1) :

            if (isOkay(grille, i, ligne, col)) :

                grille[0][ligne][col] = i
                if (algoBacktracking(grille, ligne+(col+1)//taille, (col+1)%taille)) :

                    return True

        grille[0][ligne][col] = 0
        return False 
    
def algoBacktracking2 (grille, ligne, col) :        # AFFICHE LE NB DE SOLUTIONS

    if (ligne >= taille or col >= taille) :

        return 1
    
    elif (grille[0][ligne][col] != 0) :

        return algoBacktracking2(grille, ligne+(col+1)//taille, (col+1)%taille)
    
    else :

        n = 0

        for i in range (1,taille+1) :

            if (isOkay(grille, i, ligne, col)) :

                grille[0][ligne][col] = i
                n = n + algoBacktracking2(grille, ligne+(col+1)//taille, (col+1)%taille)

        grille[0][ligne][col] = 0
        return n 


def backtracking (grille) :

    n = algoBacktracking2(grille, 0, 0)
    print(n)

def printLineToGrille (input) :         # transformer une ligne de texte sudoku en grille       NOTE : renvoyer la grille au lieu de print

    grille = np.zeros((taille+1,taille+2,taille+1), dtype=int)

    for i in range (len(input)) :

        if (input[i] == ".") :
            grille[0][i//taille][i%taille] = 0
        else :
            grille[0][i//taille][i%taille] = int(input[i])

    afficher(grille, 0)

def printGrilleToLine (grille) :        # transformer une grille en ligne de texte sudoku       NOTE : renvoyer au lieu de print
    output = ""

    for i in range (taille) :
        for j in range (taille) :
            if grille[0][i][j] == 0 :
                output = output + "."
            else :
                output = output + str(grille[0][i][j])

    print(output)

        
grille = np.zeros((taille+1,taille+2,taille+1), dtype=int)

grille[0] = [[0, 0, 0, 0, 4, 0, 0, 1, 0, 0,],
 [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
 [0, 3, 2, 1, 9, 0, 0, 8, 0, 0,],
 [0, 0, 5, 3, 0, 0, 9, 0, 4, 0,],
 [0, 0, 4, 9, 0, 0, 0, 6, 1, 0,],
 [0, 0, 9, 0, 7, 4, 0, 0, 0, 0,],
 [0, 9, 6, 7, 0, 1, 5, 0, 0, 0,],
 [5, 0, 0, 0, 3, 0, 0, 2, 7, 0,],
 [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
 [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
 [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,]]

majCompteurs(grille)
afficher(grille,0)
#remplirPossibilite(grille) 
#afficher(grille, 1)

#backtracking(grille)
#afficher(grille,0)
 
#printLineToGrille(".5..1.2....85...1.....3...8.....2....7...6.3.1...7.9..7.....5..4......6.3..8...4.")
printGrilleToLine(grille)