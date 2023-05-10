import pygame
from pygame.locals import *
import numpy as np
pygame.init()
 
fenetre = pygame.display.set_mode((450,500))        # mettre l'écran
continuer = True
nb_cases_cote = 9
TAILLE = 9

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
 
taille_case = min(fenetre.get_size()) / nb_cases_cote # min renvoie la valeur minimale d'une liste, ici la dimension de la fenêtre
print(taille_case)
 
font = pygame.font.SysFont("",20)
 
while continuer:
    fenetre.fill((255,255,255))
 
    events = pygame.event.get()
    for event in events:
        if event.type == QUIT:
            continuer = False
            # mettre la sauvegarde ici !!

        #if event.type == MOUSEBUTTONDOWN:  # on click
            # récuperer les coords et selon elle, 
 
    for x in range(TAILLE):  # ligne
        for y in range(TAILLE):  # colonne 
 
            pygame.draw.rect(fenetre, [0]*3, [x*taille_case, y*taille_case, taille_case, taille_case], 1) # dessin rectangle 
            pygame.draw.line(fenetre, [0,0,255], [taille_case*3-1, 0], [ taille_case*3-1, taille_case*TAILLE], 3)
            lettre = font.render("%d" % grille[0][x][y], True, [0]*3) # la valeur de la case du sudoku
            lettre_rect = lettre.get_rect() 
            lettre_rect.center = [x*taille_case + 1/2*taille_case, y*taille_case + 1/2*taille_case] # mise du centre du rect au milieu de la case
            fenetre.blit( lettre , lettre_rect ) # on blit le tout
 
    pygame.display.flip()