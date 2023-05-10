import numpy as np
from math import *
from tqdm import tqdm

TAILLE: int = 9

def afficher_grille (grille:np.ndarray, plan:int) -> None:   
    """ afficher comme une grille un des plan de la matrice 
    (0:sudoku, 123...:possibilités).
    
    Arguments :
    grille  : numpy.ndarray  -- la grille de jeu à afficher
    plan    : int            -- le plan de la grille à afficher
    """
    
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
    
def maj_compteurs (grille:np.ndarray) -> None:
    """ met a jour les compteurs de nb de cases par ligne/colonne/carré
    
    Arguments :
    grille  : numpy.ndarray  -- la grille de jeu avec les compteurs
    """

    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if (grille[0][line][column] != 0) :
                grille[0][line][TAILLE] += 1                       # maj des lignes 
                grille[0][TAILLE][column] += 1                       # maj des colonnes 
                grille[0][TAILLE][TAILLE] += 1
                grille[0][TAILLE+1][(line//int(sqrt(TAILLE)))*int(sqrt(TAILLE))+(column//int(sqrt(TAILLE)))] += 1       # maj des carrés

def is_okay_case (grille:np.ndarray, valeur:int, ligne:int, col:int) -> bool :
    """ vérifie si une valeur est possible à un endroit ligne/col donné 
    dans la grille de jeu 
    
    Arguments :
    grille  : numpy.ndarray -- la grille de jeu contenant la case à vérifier
    valeur  : int           -- la valeur à tester
    ligne   : int           -- la ligne sur laquelle est la case
    col     : int           -- la colonne sur laquelle est la case
    """

    for i in range (TAILLE) :

        if (grille[0][ligne][i] == valeur 
            or grille[0][i][col] == valeur 
            or grille[0][(ligne // int(sqrt(TAILLE))) * int(sqrt(TAILLE)) + i // int(sqrt(TAILLE))][(col // int(sqrt(TAILLE))) * int(sqrt(TAILLE)) + i % int(sqrt(TAILLE))] == valeur) :
            return False 
        
    return True

def is_okay_grille (grille:np.ndarray) -> bool :
    """ prend n'importe quelle grille COMPLETE et renvoi si elle est valide 
    ou non selon les contraintes, sans distinction de case (déjà là ou 
    rentrée manuellement) 
    
    Arguments :
    grille  : numpy.ndarray -- la grille de jeu à vérifier
    """

    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if not is_okay_case(grille, grille[0][line][column], line, column) :

                return False
            
    return True


def remplir_possibilite (grille:np.ndarray) -> None :
    """ rempli toute la matrice 3D de sorte a avoir des 1 sur le plan X si X 
    est possible d'être mis à cette case 
    
    Arguments :
    grille  : numpy.ndarray -- la grille de jeu 3D à remplir
    """

    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if (grille[0][line][column] == 0) :

                for valeur in range (1,TAILLE+1) :

                    if (is_okay_case(grille, valeur, line, column)) :

                        grille[valeur][line][column] = 1

def algo_backtracking (grille:np.ndarray, ligne:int, col:int) -> bool :         
    """ resout avec le backtracking mais sans afficher si solution unique 
    ou pas 
    
    Arguments :
    grille  : numpy.ndarray -- la grille de jeu à résoudre
    ligne   : int           -- ligne de la case testée actuellement (pour la récursivité)
    col     : int           -- colonne de la case testée actuellement (pour la récursivité)
    """

    # condition d'arrêt : si on a parcouru toute la grille cad arrivé en case (TAILLE, TAILLE)
    if (ligne >= TAILLE 
        or col >= TAILLE) :
        return True
    
    # ne vérifier que les cases qui ne sont pas pré-remplies
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
    """ cherche le nombre de solutions possible à un sudoku, TRES LONG
    RENVOIE UNIQUEMENT LE NOMBRE DE SOLUTION & MODIFIE LA MATRICE D'ENTREE !! 
    (elle deviendra la dernière solution trouvée si y en a plusieurs) 
    
    Arguments :
    grille  : numpy.ndarray -- la grille de jeu à résoudre
    ligne   : int           -- ligne de la case testée actuellement (pour la récursivité)
    col     : int           -- colonne de la case testée actuellement (pour la récursivité)
    bar                     -- pour la barre de chargement dans le terminal (A ENLEVER)
    """

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

def backtracking (grille: np.ndarray, printNbSolution: bool = False) :    
    """ lance l'algo de résolution avec ou sans le nb de solutions 
    
    Arguments :
    grille          : numpy.ndarray -- la grille de jeu à résoudre
    printNbSolution : bool          -- choix du mode (solutions multiples ou non)
    """

    if printNbSolution :
        n = 3265920             # 9 factorielle * 9 
        bar = tqdm(total=n)     # bar de progression
        n = algo_backtracking_multiples(grille, 0, 0, bar)
        bar.close()
        print(n)

    else :
        algo_backtracking(grille, 0, 0)
    
def line_to_grille (input : str) -> np.ndarray :         
    """ transformer une ligne de texte sudoku en grille 

    Arguments :
    input   : str   -- la ligne de texte formatée à convertir
    """

    grille = np.zeros((TAILLE+1, TAILLE+2, TAILLE+1), dtype=int)
    for i in range (len(input)) :

        if (input[i] == ".") :
            grille[0][i//TAILLE][i%TAILLE] = 0

        else :
            grille[0][i//TAILLE][i%TAILLE] = int(input[i])

    return grille

def grille_to_line (grille:np.ndarray) -> str:        
    """ transformer une grille en ligne de texte sudoku 
    
    Arguments :
    grille  : numpy.ndarray -- la grille de jeu à convertir
    """

    output = ""
    for line in range (TAILLE) :
        for column in range (TAILLE) :

            if grille[0][line][column] == 0 :
                output = output + "."

            else :
                output = output + str(grille[0][line][column])

    return output
        
grille = np.zeros((TAILLE+1,TAILLE+2,TAILLE+1), dtype=int)
print(type(grille))
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
#remplir_possibilite(grille) 
afficher_grille(grille, 0)

backtracking(grille, False)
#afficher_grille(grille,0)
 
#line_to_grille(".5..1.2....85...1.....3...8.....2....7...6.3.1...7.9..7.....5..4......6.3..8...4.")
#grille_to_line(grille)

#A = (np.random.rand(640, 480) * 255 * 255 * 255).astype(np.uint32)