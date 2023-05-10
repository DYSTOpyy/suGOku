import pygame
from pygame.locals import *
import numpy as np
pygame.init()
 
fenetre = pygame.display.set_mode((450,500))        # mettre l'écran
continuer = True
nb_cases_cote = 9
TAILLE = 9
TAILLE_CASE = min(fenetre.get_size()) / nb_cases_cote # min renvoie la valeur minimale d'une liste, ici la dimension de la fenêtre

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
 
print(TAILLE_CASE)
 
font = pygame.font.SysFont("",35)
flag = 0
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
            flag = 1
 
    for x in range(TAILLE):  # ligne
        for y in range(TAILLE):  # colonne 
 
            pygame.draw.rect(fenetre, [0]*3, [x*TAILLE_CASE, y*TAILLE_CASE, TAILLE_CASE, TAILLE_CASE], 1) # dessin rectangle 
            lettre = font.render("%d" % grille[0][x][y], True, [0]*3) # la valeur de la case du sudoku
            lettre_rect = lettre.get_rect() 
            lettre_rect.center = [x*TAILLE_CASE + 1/2*TAILLE_CASE, y*TAILLE_CASE + 1/2*TAILLE_CASE] # mise du centre du rect au milieu de la case
            fenetre.blit( lettre , lettre_rect ) # on blit le tout

    pygame.draw.line(fenetre, [0,0,255], [TAILLE_CASE*3-1, 0], [ TAILLE_CASE*3-1, TAILLE_CASE*TAILLE], 3)
    pygame.draw.line(fenetre, [0,0,255], [ TAILLE_CASE*TAILLE, TAILLE_CASE*3-1], [ 0 , TAILLE_CASE*3-1], 3)
    pygame.draw.line(fenetre, [0,0,255], [TAILLE_CASE*6-1, 0], [ TAILLE_CASE*6-1, TAILLE_CASE*TAILLE], 3)
    pygame.draw.line(fenetre, [0,0,255], [ TAILLE_CASE*TAILLE, TAILLE_CASE*6-1], [ 0 , TAILLE_CASE*6-1], 3)

    if flag == 1:
        highlight(selected_x, selected_y)
 
    pygame.display.flip()