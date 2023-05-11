# -*- coding: utf-8 -*-

import numpy as np
import subprocess
SIZE = 3
SIZE2 = SIZE**2
SIZE4 = SIZE2**2

def affichage(tab : np.ndarray):
    string = ""
    for i in range(9):
        for j in range(9):
            value = str(tab[i][j])
            if value == "0":
                value = " "
            string+=value
            if j%3==2 and j != 8:
                string+="|"
        string+="\n"
        if i%3==2 and i != 8:
                string+=11*"-"+"\n"
    print(string)



def valLig(tab : np.ndarray):   
    for i in range (SIZE2):
        somme=0
        for j in range (SIZE2):
            if tab[i,j]!=0 :
                somme+=1
        tab[i,SIZE2]=somme

def valCol(tab : np.ndarray):
    for i in range(SIZE2):
        somme=0
        for j in range (SIZE2):
            if tab[j,i]!=0 :
                somme+=1
        tab[SIZE2,i]=somme

def valSquare(tab : np.ndarray):
    for s in range(SIZE2):
        somme=0
        lig=s%SIZE*SIZE
        col=s//SIZE*SIZE
        for i in range (lig,lig+SIZE):
            for j in range (col,col+SIZE):
                if tab[i,j]!=0 :
                    somme+=1
        tab[SIZE2+1,s]=somme


# def totalGrid(tab):
#     lig = 0
#     col = 0
#     somme = 0
#     for _ in range(SIZE4):
#         somme += tab[SIZE2,col]*tab[lig,SIZE2]*tab[SIZE2+1,col]
#         col+=1
#         if col == 9:
#             col = 0
#             lig +=1
#     return somme

def init():
    tab = importSudoku()
    for i in range(SIZE2):
        valLig(tab) # type: ignore
        valCol(tab) # type: ignore
        valSquare(tab) # type: ignore
    affichage(tab)  # type: ignore

def importSudoku() :
    try : 
        data = open("output.txt", "r")
    except FileNotFoundError :
        return FileNotFoundError
    input=data.read().split('\n')[0]
    tab = np.zeros((SIZE2+2,SIZE2+1),int)
    for i in range(SIZE4):
        char=input[i]
        if char =='.':
            char = "0"
        tab[i//9,i%9]= char
    return tab


init()
subprocess.run(["go","run","learn.go"])