import pygame
from pygame.locals import *
import numpy as np
from math import *

TAILLE = 9
TAILLE_CASE = 50

taille_grille = TAILLE * TAILLE_CASE

pygame.init()
fenetre = pygame.display.set_mode((TAILLE_CASE*TAILLE, TAILLE_CASE*TAILLE))        # mettre l'écran
continuer = True

def highlight (x,y) :
    
    x = x // TAILLE_CASE
    y = y // TAILLE_CASE

    pygame.draw.rect(fenetre, [255,0,0], [x*TAILLE_CASE, y*TAILLE_CASE, TAILLE_CASE, TAILLE_CASE], 3) # dessin rectangle

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

masque = np.copy(grille)
 
font = pygame.font.SysFont("",35)
selection_case = 0
selected_x, selected_y = 0, 0

def change_case (valeur, x, y) :

    x = x // TAILLE_CASE
    y = y // TAILLE_CASE

    if masque[0][x][y] == 0 :
        grille[0][x][y] = valeur

while continuer:

    fenetre.fill((255,255,255))

    events = pygame.event.get()
    for event in events:

        if event.type == QUIT:
            continuer = False
            # mettre la sauvegarde ici !!

        #if event.type == MOUSEBUTTONDOWN:  # on click
            # récuperer les coords et selon elle, 
        if event.type == MOUSEBUTTONDOWN:
            selected_x, selected_y = pygame.mouse.get_pos()
            if (0 <= selected_x and selected_x <= taille_grille) and (0 <= selected_y and selected_y <= taille_grille) : 
                selection_case = 1
            else :
                selection_case = 0
        if event.type == pygame.KEYUP and selection_case:
                if event.key == pygame.K_1 or event.key == pygame.K_KP1:
                    change_case(1, selected_x, selected_y)
                elif event.key == pygame.K_2 or event.key == pygame.K_KP2:
                    change_case(2, selected_x, selected_y)
                elif event.key == pygame.K_3 or event.key == pygame.K_KP3:
                    change_case(3, selected_x, selected_y)
                elif event.key == pygame.K_4 or event.key == pygame.K_KP4:
                    change_case(4, selected_x, selected_y)
                elif event.key == pygame.K_5 or event.key == pygame.K_KP5:
                    change_case(5, selected_x, selected_y)
                elif event.key == pygame.K_6 or event.key == pygame.K_KP6:
                    change_case(6, selected_x, selected_y)
                elif event.key == pygame.K_7 or event.key == pygame.K_KP7:
                    change_case(7, selected_x, selected_y)
                elif event.key == pygame.K_8 or event.key == pygame.K_KP8:
                    change_case(8, selected_x, selected_y)
                elif event.key == pygame.K_9 or event.key == pygame.K_KP9:
                    change_case(9, selected_x, selected_y)

 
    for x in range(TAILLE):  # ligne
        for y in range(TAILLE):  # colonne 
 
            pygame.draw.rect(fenetre, [0]*3, [x*TAILLE_CASE, y*TAILLE_CASE, TAILLE_CASE, TAILLE_CASE], 1) # dessin rectangle 
            lettre = font.render("%d" % grille[0][x][y], True, [0]*3) # la valeur de la case du sudoku
            lettre_rect = lettre.get_rect() 
            lettre_rect.center = [x*TAILLE_CASE + 1/2*TAILLE_CASE, y*TAILLE_CASE + 1/2*TAILLE_CASE] # mise du centre du rect au milieu de la case
            fenetre.blit( lettre , lettre_rect ) # on blit le tout

    # Dessin des 4 lignes séparant les blocs

    for nb in range (1, int(sqrt(TAILLE))) :
        pygame.draw.line(fenetre, [0,0,255], [TAILLE_CASE*int(nb*sqrt(TAILLE))-1, 0], [ TAILLE_CASE*int(nb*sqrt(TAILLE))-1, TAILLE_CASE*TAILLE], 3)
        pygame.draw.line(fenetre, [0,0,255], [ TAILLE_CASE*TAILLE, TAILLE_CASE*int(nb*sqrt(TAILLE))-1], [ 0 , TAILLE_CASE*int(nb*sqrt(TAILLE))-1], 3)


    if selection_case:
        highlight(selected_x, selected_y)
 
    pygame.display.flip()