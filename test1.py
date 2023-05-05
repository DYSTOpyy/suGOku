import numpy as np
INPUT1 = ".5.4283..7.....9....3....2.6....1.5.3..6.2..1.8.7....3.3....6....9.....5..1863.9."
SIZE = 3
SIZE2 = SIZE**2
SIZE4 = SIZE2**2

def affichage():
    input1=INPUT1
    string=""
    for i in range(9):
        for j in range(9):
            string+=input1[j]
            if j%3==2 and j != 8:
                string+="|"
        input1=input1[9:]
        string+="\n"
        if i%3==2and i != 8:
                string+=11*"-"+"\n"
    print(string)

affichage()

def toTab():
    tab = np.zeros((SIZE2+2,SIZE2+1),int)
    for i in range(SIZE4):
        char=INPUT1[i]
        if char =='.':
            char = "0"
        tab[i//9,i%9]= char
    return tab

def valLig(tab):   
    for i in range (SIZE2):
        somme=0
        for j in range (SIZE2):
            if tab[i,j]!=0 :
                somme+=1
        tab[i,SIZE2]=somme

def valCol(tab):
    for i in range(SIZE2):
        somme=0
        for j in range (SIZE2):
            if tab[j,i]!=0 :
                somme+=1
        tab[SIZE2,i]=somme

def valSquare(tab):
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
    tab=toTab()
    for i in range(SIZE2):
        valLig(tab)
        valCol(tab)
        valSquare(tab)
    print(tab)


init()