import numpy as np
from math import *
from tqdm import tqdm

TAILLE: int = 16

def afficher_grille (grille, plan:int) :   
    """ afficher comme une grille un des plan de la matrice 
    (0:sudoku, 123...:possibilités) """
    
    print("Affichage du plan n°",plan)
    # pour chaque ligne 
    for line in range (TAILLE) :

        if (line % int(sqrt(TAILLE)) == 0 
            and line != 0) :
            print("-----------")

        # pour chaque colonne
        for column in range (TAILLE) :

            if (column % int(sqrt(TAILLE)) == 0 
                and column != 0) :
                print("|", end="")

            if grille[plan][line][column] == 0 :
                print(".", end="")

            else :
                print(grille[plan][line][column], end="")

        # compteur de ligne
        print(" (",grille[plan][line][TAILLE],")")

    print()

    # compteur de colonne + bloc
    for line in range (TAILLE, TAILLE+2) :
        for column in range (TAILLE) :

            if (column % int(sqrt(TAILLE)) == 0 
                and column != 0) :
                print("|", end="")

            print(grille[plan][line][column], end="")

        print()

    print("Nombre d'indices total :", grille[0][TAILLE][TAILLE])
    
def maj_compteurs (grille) :
    """ met a jour les compteurs de nb de cases par ligne/colonne/carré """

    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if (grille[0][line][column] != 0) :
                grille[0][line][TAILLE] += 1                       # maj des lignes 
                grille[0][TAILLE][column] += 1                       # maj des colonnes 
                grille[0][TAILLE][TAILLE] += 1
                grille[0][TAILLE+1][(line//int(sqrt(TAILLE)))*int(sqrt(TAILLE))+(column//int(sqrt(TAILLE)))] += 1       # maj des carrés

def is_okay_case (grille, valeur:int, ligne:int, col:int) :
    """ vérifie si une valeur est possible à un endroit ligne/col donné 
    dans la grille de jeu """

    for i in range (TAILLE) :

        if (grille[0][ligne][i] == valeur 
            or grille[0][i][col] == valeur 
            or grille[0][(ligne // int(sqrt(TAILLE))) * int(sqrt(TAILLE)) + i // int(sqrt(TAILLE))][(col // int(sqrt(TAILLE))) * int(sqrt(TAILLE)) + i % int(sqrt(TAILLE))] == valeur) :
            return False 
        
    return True

def is_okay_grille (grille) :
    """ prend n'importe quelle grille COMPLETE et renvoi si elle est valide 
    ou non selon les contraintes, sans distinction de case (déjà là ou 
    rentrée manuellement) """

    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if not is_okay_case(grille, grille[0][line][column], line, column) :

                return False
            
    return True


def remplir_possibilite (grille) :
    """ rempli toute la matrice 3D de sorte a avoir des 1 sur le plan X si X 
    est possible d'être mis à cette case """

    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if (grille[0][line][column] == 0) :

                for valeur in range (1,TAILLE+1) :

                    if (is_okay_case(grille, valeur, line, column)) :

                        grille[valeur][line][column] = 1

def algo_backtracking (grille, ligne:int, col:int) :         
    """ resout avec le backtracking mais sans afficher si solution unique 
    ou pas """

    # condition d'arrêt : si on a parcouru toute la grille cad arrivé en case (TAILLE, TAILLE)
    if (ligne >= TAILLE 
        or col >= TAILLE) :
        return True
    
    # si la case n'est pas vide on passe à la suivante
    elif (grille[0][ligne][col] != 0) :

        return algo_backtracking(grille, ligne+(col+1)//TAILLE, (col+1)%TAILLE)
    
    else :

        for valeur in range (1,TAILLE+1) :

            if (is_okay_case(grille, valeur, ligne, col)) :

                grille[0][ligne][col] = valeur
                if (algo_backtracking(grille, ligne+(col+1)//TAILLE, (col+1)%TAILLE)) :

                    return True

        grille[0][ligne][col] = 0
        return False 
    
def algo_backtracking_multiples (grille, ligne:int, col:int, bar) -> int :        
    """ RENVOIE UNIQUEMENT LE NOMBRE DE SOLUTION & MODIFIE LA MATRICE D'ENTREE !! 
    (elle deviendra la dernière solution trouvée si y en a plusieurs) """

    # condition d'arrêt : si on a parcouru toute la grille cad arrivé en case (TAILLE, TAILLE)
    if (ligne >= TAILLE 
        or col >= TAILLE) :
        return 1        # on renvoi +1 car on a trouvé une nouvelle solution
    
    elif (grille[0][ligne][col] != 0) :

        return algo_backtracking_multiples(grille, ligne+(col+1)//TAILLE, (col+1)%TAILLE, bar)
        
    else :

        n = 0
        for valeur in range (1,TAILLE+1) :

            bar.update(1)

            if (is_okay_case(grille, valeur, ligne, col)) :

                grille[0][ligne][col] = valeur
                n = n + algo_backtracking_multiples(grille, ligne+(col+1)//TAILLE, (col+1)%TAILLE, bar)    # pas de if ici car on test toutes les valeurs du for même si on a trouvé une première solution

        grille[0][ligne][col] = 0
        return n 

def backtracking (grille, printNbSolution: bool = False) :    
    """ lance l'algo de résolution avec ou sans le nb de solutions """

    if printNbSolution :
        n = 3265920             # 9 factorielle * 9 
        bar = tqdm(total=n)     # bar de progression
        n = algo_backtracking_multiples(grille, 0, 0, bar)
        bar.close()
        print(n)

    else :
        algo_backtracking(grille, 0, 0)
    
def line_to_grille (input : str) :         
    """ transformer une ligne de texte sudoku en grille """

    grille = np.zeros((TAILLE+1, TAILLE+2, TAILLE+1), dtype=int)
    for i in range (len(input)) :

        if (input[i] == ".") :
            grille[0][i//TAILLE][i%TAILLE] = 0

        else :
            grille[0][i//TAILLE][i%TAILLE] = int(input[i])

    return grille

def grille_to_line (grille) -> str:        
    """ transformer une grille en ligne de texte sudoku """

    output = ""
    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if grille[0][line][column] == 0 :
                output = output + "."

            else :
                output = output + str(grille[0][line][column])

    return output
        
grille = np.zeros((TAILLE+1,TAILLE+2,TAILLE+1), dtype=int)
"""grille[0] = [[0, 0, 0, 0, 0, 0, 0, 1, 0, 0,],
             [0, 0, 0, 0, 4, 0, 0, 0, 0, 0,],
             [0, 3, 2, 1, 9, 0, 0, 8, 0, 0,],
             [0, 0, 5, 3, 0, 0, 9, 0, 4, 0,],
             [0, 0, 4, 9, 0, 0, 0, 6, 1, 0,],
             [0, 0, 9, 0, 7, 4, 0, 0, 0, 0,],
             [0, 9, 6, 7, 0, 1, 5, 0, 0, 0,],
             [5, 0, 0, 0, 3, 0, 0, 2, 7, 0,],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,]] """

grille[0] = [[0, 15, 0, 1, 0, 2, 10, 14, 12, 0, 0, 0, 0, 0, 0, 0, 0],
[0, 6, 3, 16, 12, 0, 8, 4, 14, 15, 1, 0, 2, 0, 0, 0, 0],
[14, 0, 9, 7, 11, 3, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
[4, 13, 2, 12, 0, 0, 0, 0, 6, 0, 0, 0, 0, 15, 0, 0, 0],
[0, 0, 0, 0, 14, 1, 11, 7, 3, 5, 10, 0, 0, 8, 0, 12, 0],
[3, 16, 0, 0, 2, 4, 0, 0, 0, 14, 7, 13, 0, 0, 5, 15, 0],
[11, 0, 5, 0, 0, 0, 0, 0, 0, 9, 4, 0, 0, 6, 0, 0, 0],
[0, 0, 0, 0, 13, 0, 16, 5, 15, 0, 0, 12, 0, 0, 0, 0, 0],
[0, 0, 0, 0, 9, 0, 1, 12, 0, 8, 3, 10, 11, 0, 15, 0, 0],
[2, 12, 0, 11, 0, 0, 14, 3, 5, 4, 0, 0, 0, 0, 9, 0, 0],
[6, 3, 0, 4, 0, 0, 13, 0, 0, 11, 9, 1, 0, 12, 16, 2, 0],
[0, 0, 10, 9, 0, 0, 0, 0, 0, 0, 12, 0, 8, 0, 6, 7, 0],
[12, 8, 0, 0, 16, 0, 0, 10, 0, 13, 0, 0, 0, 5, 0, 0, 0],
[5, 0, 0, 0, 3, 0, 4, 6, 0, 1, 15, 0, 0, 0, 0, 0, 0],
[0, 9, 1, 6, 0, 14, 0, 11, 0, 0, 2, 0, 0, 0, 10, 8, 0],
[0, 14, 0, 0, 0, 13, 9, 0, 4, 12, 11, 8, 0, 0, 2, 0, 0],
[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]]

maj_compteurs(grille)
#remplir_possibilite(grille) 
afficher_grille(grille, 0)

backtracking(grille, False)
#afficher_grille(grille,0)
 
#line_to_grille(".5..1.2....85...1.....3...8.....2....7...6.3.1...7.9..7.....5..4......6.3..8...4.")
#grille_to_line(grille)

#A = (np.random.rand(640, 480) * 255 * 255 * 255).astype(np.uint32)

