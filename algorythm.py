import numpy as np
from math import *
from tqdm import tqdm

TAILLE: int = 9

def afficher (grille, plan:int) :   
    """ afficher comme une grille un des plan de la matrice 
    (0:sudoku, 123...:possibilités) """
    
    print("Affichage du plan n°",plan)
    # pour chaque ligne 
    for i in range (TAILLE) :

        if (i % int(sqrt(TAILLE)) == 0 
            and i != 0) :
            print("-----------")

        # pour chaque colonne
        for j in range (TAILLE) :

            if (j % int(sqrt(TAILLE)) == 0 
                and j != 0) :
                print("|", end="")

            if grille[plan][i][j] == 0 :
                print(".", end="")

            else :
                print(grille[plan][i][j], end="")

        # compteur de ligne
        print(" (",grille[plan][i][TAILLE],")")

    print()

    # compteur de colonne + bloc
    for i in range (TAILLE, TAILLE+2) :
        for j in range (TAILLE) :

            if (j % int(sqrt(TAILLE)) == 0 
                and j != 0) :
                print("|", end="")

            print(grille[plan][i][j], end="")

        print()

    print("Nombre d'indices total :", grille[0][TAILLE][TAILLE])
    

def maj_compteurs (grille) :
    """ met a jour les compteurs de nb de cases par ligne/colonne/carré """

    for i in range (TAILLE) :
        for j in range (TAILLE) :

            if (grille[0][i][j] != 0) :
                grille[0][i][TAILLE] += 1                       # maj des lignes 
                grille[0][TAILLE][j] += 1                       # maj des colonnes 
                grille[0][TAILLE][TAILLE] += 1
                grille[0][TAILLE+1][(i//int(sqrt(TAILLE)))*int(sqrt(TAILLE))+(j//int(sqrt(TAILLE)))] += 1       # maj des carrés

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

    for i in range (TAILLE) :
        for j in range (TAILLE) :

            if not is_okay_case(grille, grille[0][i][j], i, j) :

                return False
            
    return True


def remplir_possibilite (grille) :
    """ rempli toute la matrice 3D de sorte a avoir des 1 sur le plan X si X 
    est possible d'être mis à cette case """

    for i in range (TAILLE) :
        for j in range (TAILLE) :

            if (grille[0][i][j] == 0) :

                for valeur in range (1,TAILLE+1) :

                    if (is_okay_case(grille, valeur, i, j)) :

                        grille[valeur][i][j] = 1

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

        for i in range (1,TAILLE+1) :

            if (is_okay_case(grille, i, ligne, col)) :

                grille[0][ligne][col] = i
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
        for i in range (1,TAILLE+1) :

            bar.update(1)

            if (is_okay_case(grille, i, ligne, col)) :

                grille[0][ligne][col] = i
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
    for i in range (TAILLE) :
        for j in range (TAILLE) :

            if grille[0][i][j] == 0 :
                output = output + "."

            else :
                output = output + str(grille[0][i][j])

    return output
        
grille = np.zeros((TAILLE+1,TAILLE+2,TAILLE+1), dtype=int)
grille[0] = [[0, 0, 0, 0, 0, 0, 0, 1, 0, 0,],
             [0, 0, 0, 0, 4, 0, 0, 0, 0, 0,],
             [0, 3, 2, 1, 9, 0, 0, 8, 0, 0,],
             [0, 0, 5, 3, 0, 0, 9, 0, 4, 0,],
             [0, 0, 4, 9, 0, 0, 0, 6, 1, 0,],
             [0, 0, 9, 0, 7, 4, 0, 0, 0, 0,],
             [0, 9, 6, 7, 0, 1, 5, 0, 0, 0,],
             [5, 0, 0, 0, 3, 0, 0, 2, 7, 0,],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,]]

maj_compteurs(grille)
remplir_possibilite(grille) 
afficher(grille, 0)

backtracking(grille, True)
afficher(grille,0)
 
#line_to_grille(".5..1.2....85...1.....3...8.....2....7...6.3.1...7.9..7.....5..4......6.3..8...4.")
#grille_to_line(grille)