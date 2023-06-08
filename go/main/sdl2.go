package main

import (
	"math"
	"strconv"
	"time"

	"git.saussesylva.in/DYSTO_pyy/Sudoku/algo"
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	TAILLE = utils.TAILLE // Taille MAX de la MATRICE de sudoku
	Game   = iota
	Menu
	Loading
)

// variable ici pour que TOUT LE MONDE puisse les lire
var (
	windowVertical   int32  = 600
	windowHorizontal int32  = 800
	titre            string = "Sudoku!"
	size             int    = utils.Size // ordre du sudoku en jeu (9x9, 16x16)

	renderer    *sdl.Renderer
	window      *sdl.Window
	font        *ttf.Font
	fontSurface *sdl.Surface

	hooverButtonGame [4]bool // état des boutons en jeu (souris dessus ou non)
	hooverButtonMenu [5]bool // état des boutons en menue

	masque      [TAILLE][TAILLE]bool        // garde en mémoire les cases ajoutées par le joueur ou présentes depuis le début
	verifier    [TAILLE][TAILLE]bool        // indique si des cases sont incorrectes ou non
	grille      [TAILLE + 2][TAILLE + 1]int // grille de jeu 16x16 avec 2 lignes de compteurs de cases et 1 colonne de compteur de cases
	possibilite [TAILLE][TAILLE][]int       // répertorie pour chaque case les valeurs possibles

	taille_case int32 // taille d'un carré de sudoku

	running       bool = true
	selectionCase bool = false
	resoudre      bool = false

	screen                 int // écran actuel
	difficulte             int = 1
	selected_x, selected_y int32
	startTime              time.Time
	messageGame            string = ""
	messageMenu            string = ""
	err                    error
)

// Initialise les composants nécessaire pour SDL
func start() {

	// init
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	// init ttf
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	// init windows (la fenêtre de jeu)
	window, err = sdl.CreateWindow(titre, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, windowHorizontal, windowVertical, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	// init renderer  (pour dessiner dessus)
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		panic(err)
	}

}

// Change la valeur d'une case si elle n'est pas un indice
func changeCase(valeur int, grille *[TAILLE + 2][TAILLE + 1]int, masque [TAILLE][TAILLE]bool) {
	if !masque[selected_y][selected_x] {
		grille[selected_y][selected_x] = valeur
	}
}

// Libération de la mémoire à la fin du programme
func destroy() {

	window.Destroy()
	ttf.Quit()
	sdl.Quit()
}

// Écran de jeu du sudoku
func interface_jeu(grille *[TAILLE + 2][TAILLE + 1]int) {

	// init la transparence
	if err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		panic(err)
	}

	// init de l'image bouton Vérifier
	imgVerifier, err := img.Load("assets/Button 9.png")
	if hooverButtonGame[0] {
		imgVerifier, err = img.Load("assets/Button 9 Hoover.png")
	}
	defer imgVerifier.Free()
	if err != nil {
		panic(err)
	}
	coordBoutonVerifier := [4]int32{int32(size+1) * taille_case, 3 * taille_case, int32(float64(imgVerifier.W)), int32(float64(imgVerifier.H))} // coin haut gauche x, y, windowVertical, windowHorizontal

	// init de l'image bouton Menu
	imgMenu, err := img.Load("assets/Button Menu.png")
	if hooverButtonGame[1] {
		imgMenu, err = img.Load("assets/Button Menu Hoover.png")
	}
	defer imgMenu.Free()
	if err != nil {
		panic(err)
	}
	coordBoutonMenu := [4]int32{int32(size+1) * taille_case, 3*taille_case + 50 + coordBoutonVerifier[3], int32(float64(imgMenu.W)), int32(float64(imgMenu.H))} // coin haut gauche x, y, windowVertical, windowHorizontal

	// init de l'image bouton Recommencer
	imgRecommencer, err := img.Load("assets/Button Recommencer.png")
	if hooverButtonGame[2] {
		imgRecommencer, err = img.Load("assets/Button Recommencer Hoover.png")
	}
	defer imgRecommencer.Free()
	if err != nil {
		panic(err)
	}
	coordBoutonRecommencer := [4]int32{int32(size+1) * taille_case, 3*taille_case + 100 + coordBoutonVerifier[3] + coordBoutonMenu[3], int32(float64(imgRecommencer.W)), int32(float64(imgRecommencer.H))} // coin haut gauche x, y, windowVertical, windowHorizontal

	// init de l'image bouton Resoudre
	imgResoudre, err := img.Load("assets/Button Resoudre.png")
	if hooverButtonGame[3] {
		imgResoudre, err = img.Load("assets/Button Resoudre Hoover.png")
	}
	defer imgResoudre.Free()
	if err != nil {
		panic(err)
	}
	coordBoutonResoudre := [4]int32{int32(size+1) * taille_case, 3*taille_case + 150 + coordBoutonVerifier[3] + coordBoutonMenu[3] + coordBoutonRecommencer[3], int32(float64(imgResoudre.W)), int32(float64(imgResoudre.H))} // coin haut gauche x, y, windowVertical, windowHorizontal

	// init de la police des textes
	font, err = ttf.OpenFont("assets/Acme-Regular.ttf", int(taille_case/2))
	if err != nil {
		panic(err)
	}
	defer font.Close()

	// init de la surface des textes
	fontSurface, err = font.RenderUTF8Blended("text", sdl.Color{R: 0, G: 255, B: 0, A: 255})
	if err != nil {
		panic(err)
	}

	fontTexture, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		panic(err)
	}
	defer fontTexture.Destroy()
	fontSurface.Free()

	// dimensions du texte
	_, _, w, h, err := fontTexture.Query()
	if err != nil {
		panic(err)
	}

	var rect sdl.Rect

	// affichage du bouton Vérifier
	bouton, err := renderer.CreateTextureFromSurface(imgVerifier)
	if err != nil {
		panic(err)
	}

	rect = sdl.Rect{X: coordBoutonVerifier[0], Y: coordBoutonVerifier[1], W: coordBoutonVerifier[2], H: coordBoutonVerifier[3]}
	renderer.Copy(bouton, nil, &rect)
	bouton.Destroy()

	// affichage du bouton Menu
	bouton, err = renderer.CreateTextureFromSurface(imgMenu)
	if err != nil {
		panic(err)
	}
	rect = sdl.Rect{X: coordBoutonMenu[0], Y: coordBoutonMenu[1], W: coordBoutonMenu[2], H: coordBoutonMenu[3]}
	renderer.Copy(bouton, nil, &rect)
	bouton.Destroy()

	// affichage du bouton Recommencer
	bouton, err = renderer.CreateTextureFromSurface(imgRecommencer)
	if err != nil {
		panic(err)
	}
	rect = sdl.Rect{X: coordBoutonRecommencer[0], Y: coordBoutonRecommencer[1], W: coordBoutonRecommencer[2], H: coordBoutonRecommencer[3]}
	renderer.Copy(bouton, nil, &rect)
	bouton.Destroy()

	// affichage du bouton Resoudre
	bouton, err = renderer.CreateTextureFromSurface(imgResoudre)
	if err != nil {
		panic(err)
	}
	rect = sdl.Rect{X: coordBoutonResoudre[0], Y: coordBoutonResoudre[1], W: coordBoutonResoudre[2], H: coordBoutonResoudre[3]}
	renderer.Copy(bouton, nil, &rect)
	bouton.Destroy()

	// dessin des rectangles des cases
	renderer.SetDrawColor(0, 0, 0, 255)
	for X := 0; X < size; X++ {
		for Y := 0; Y < size; Y++ {

			// dessin des cases incorrectes après avoir cliqué sur Vérifier
			if verifier[Y][X] {
				renderer.SetDrawColor(255, 0, 0, 120)
				rect = sdl.Rect{X: int32(X) * taille_case, Y: int32(Y) * taille_case, W: taille_case, H: taille_case}
				renderer.FillRect(&rect)
				renderer.SetDrawColor(0, 0, 0, 255)
			}

			if masque[Y][X] {
				renderer.SetDrawColor(0, 0, 0, 50)
				rect = sdl.Rect{X: int32(X) * taille_case, Y: int32(Y) * taille_case, W: taille_case, H: taille_case}
				renderer.FillRect(&rect)
				renderer.SetDrawColor(0, 0, 0, 255)
			}

			rect = sdl.Rect{X: int32(X) * taille_case, Y: int32(Y) * taille_case, W: taille_case, H: taille_case}
			renderer.DrawRect(&rect)
		}
	}

	// dessiner les lignes verticales/horizontales
	renderer.SetDrawColor(0, 0, 255, 255)
	for i := 1; i <= 2; i++ {
		renderer.FillRect(&sdl.Rect{X: taille_case*int32(i)*int32(math.Sqrt(float64(size))) - 1, Y: 0, W: 3, H: taille_case * int32(size)})
		renderer.FillRect(&sdl.Rect{X: 0, Y: taille_case*int32(i)*int32(math.Sqrt(float64(size))) - 1, W: taille_case * int32(size), H: 3})
	}

	// encadrement d'une case sélectionnée
	if selectionCase {
		renderer.SetDrawColor(255, 0, 0, 255)
		rect = sdl.Rect{X: selected_x * taille_case, Y: selected_y * taille_case, W: taille_case, H: taille_case}
		renderer.FillRect(&rect)
		renderer.SetDrawColor(255, 255, 255, 255)
		rect = sdl.Rect{X: selected_x*taille_case + 2, Y: selected_y*taille_case + 2, W: taille_case - 4, H: taille_case - 4}
		renderer.FillRect(&rect)
	}

	// écriture des chiffres dans le sudoku
	renderer.SetDrawColor(0, 0, 0, 255)
	for X := 0; X < size; X++ {
		for Y := 0; Y < size; Y++ {

			if grille[X][Y] != 0 {

				if grille[X][Y] > 9 {
					fontSurface, err = font.RenderUTF8Blended(string(rune(grille[X][Y]-10+65)), sdl.Color{0, 0, 0, 255})
				} else {
					fontSurface, err = font.RenderUTF8Blended(string(rune(grille[X][Y]+48)), sdl.Color{0, 0, 0, 255})
				}
				if err != nil {
					panic(err)
				}

				fontTexture, err = renderer.CreateTextureFromSurface(fontSurface)
				if err != nil {
					panic(err)
				}
				fontSurface.Free()

				_, _, w, h, err = fontTexture.Query()
				if err != nil {
					panic(err)
				}

				renderer.Copy(fontTexture, nil, &sdl.Rect{X: int32(Y)*taille_case + (taille_case-w)/2, Y: int32(X)*taille_case + (taille_case-h)/2, W: w, H: h})
				fontTexture.Destroy()
			}

		}
	}

	// déclaration du timer
	after := time.Now()
	temps := strconv.Itoa(int(after.Sub(startTime).Hours())) + ":" + strconv.Itoa(int(after.Sub(startTime).Minutes())%60) + ":" + strconv.Itoa(int(after.Sub(startTime).Seconds())%60)
	fontSurface, err = font.RenderUTF8Blended(temps, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		panic(err)
	}
	fontTexture, err = renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		panic(err)
	}
	fontSurface.Free()
	_, _, w, h, err = fontTexture.Query()
	if err != nil {
		panic(err)
	}

	// affichage du timer si le jeu est pas résolu
	if !resoudre {
		renderer.Copy(fontTexture, nil, &sdl.Rect{X: int32(size+1) * taille_case, Y: taille_case, W: w, H: h})
	}
	fontTexture.Destroy()

	if messageGame != "" {
		fontSurface, err = font.RenderUTF8Blended(messageGame, sdl.Color{0, 0, 0, 255})
		if err != nil {
			panic(err)
		}

		fontTexture, err = renderer.CreateTextureFromSurface(fontSurface)
		if err != nil {
			panic(err)
		}
		fontSurface.Free()

		_, _, w, h, err = fontTexture.Query()
		if err != nil {
			panic(err)
		}

		renderer.Copy(fontTexture, nil, &sdl.Rect{X: 20, Y: int32(size+2) * taille_case, W: w, H: h})
		fontTexture.Destroy()
	}
	renderer.Present()

	// evenements
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch event.(type) {

		// fermeture de la fenêtre
		case *sdl.QuitEvent:
			running = false
			if !resoudre {
				err := utils.SaveFile(grille, masque)
				if err != nil {
					panic(err)
				}
			}

		// mouvement de la souris
		case *sdl.MouseMotionEvent:
			x, y, _ := sdl.GetMouseState()

			// affichage des hoover si la souris est sur les boutons
			hooverButtonGame[0] = (coordBoutonVerifier[0]+coordBoutonVerifier[2] >= x && x >= coordBoutonVerifier[0]) && (coordBoutonVerifier[1]+coordBoutonVerifier[3] >= y && y >= coordBoutonVerifier[1])
			hooverButtonGame[1] = (coordBoutonMenu[0]+coordBoutonMenu[2] >= x && x >= coordBoutonMenu[0]) && (coordBoutonMenu[1]+coordBoutonMenu[3] >= y && y >= coordBoutonMenu[1])
			hooverButtonGame[2] = (coordBoutonRecommencer[0]+coordBoutonRecommencer[2] >= x && x >= coordBoutonRecommencer[0]) && (coordBoutonRecommencer[1]+coordBoutonRecommencer[3] >= y && y >= coordBoutonRecommencer[1])
			hooverButtonGame[3] = (coordBoutonResoudre[0]+coordBoutonResoudre[2] >= x && x >= coordBoutonResoudre[0]) && (coordBoutonResoudre[1]+coordBoutonResoudre[3] >= y && y >= coordBoutonResoudre[1])

		// clic de la souris
		case *sdl.MouseButtonEvent:
			mouseEvent := event.(*sdl.MouseButtonEvent)

			// clic gauche
			if mouseEvent.State == sdl.PRESSED && mouseEvent.Button == sdl.BUTTON_LEFT {
				x, y, _ := sdl.GetMouseState()

				// clic sur une case : sélection de la case
				if (0 <= x && x <= int32(size)*taille_case) && (0 <= y && y <= int32(size)*taille_case) {
					selectionCase = true
					selected_x, selected_y = x/taille_case, y/taille_case
				} else {
					selectionCase = false
				}

				// clic sur vérifier : vérification des erreurs
				if (coordBoutonVerifier[0]+coordBoutonVerifier[2] >= x && x >= coordBoutonVerifier[0]) && (coordBoutonVerifier[1]+coordBoutonVerifier[3] >= y && y >= coordBoutonVerifier[1]) {
					verifier = utils.FindErrors(grille, &masque)
					if utils.EmptyBoolArray(&verifier) && utils.FullIntArray(grille) {
						messageGame = "BRAVO ! Vous avez gagné !"
						resoudre = true
					} else if utils.EmptyBoolArray(&verifier) {
						messageGame = "Aucune erreur pour l'instant, continuez !"
					} else {
						messageGame = "Une ou plusieurs incohérences trouvées."
					}

				}

				// clic sur Menu : retour au menu en sauvegardant la partie
				if (coordBoutonMenu[0]+coordBoutonMenu[2] >= x && x >= coordBoutonMenu[0]) && (coordBoutonMenu[1]+coordBoutonMenu[3] >= y && y >= coordBoutonMenu[1]) {
					if !resoudre {
						err := utils.SaveFile(grille, masque)
						if err != nil {
							panic(err)
						}
					}
					screen = Menu

				}

				// clic sur recommencer : vide toutes les entrées utilisateur
				if (coordBoutonRecommencer[0]+coordBoutonRecommencer[2] >= x && x >= coordBoutonRecommencer[0]) && (coordBoutonRecommencer[1]+coordBoutonRecommencer[3] >= y && y >= coordBoutonRecommencer[1]) && !resoudre {
					utils.RestartGrille(grille, &masque)
					verifier = utils.FindErrors(grille, &masque)
				}

				// clic sur resoudre : resoud la grille et fini la partie
				if (coordBoutonResoudre[0]+coordBoutonResoudre[2] >= x && x >= coordBoutonResoudre[0]) && (coordBoutonResoudre[1]+coordBoutonResoudre[3] >= y && y >= coordBoutonResoudre[1]) {
					if algo.Algo_backtracking(grille, &possibilite, algo.GenSlice(grille)) {
						resoudre = true
						messageGame = "Résolution de la grille terminée."
					} else {
						messageGame = "Pas de solution possible."
					}

				}

			}

		// boutons clavier
		case *sdl.KeyboardEvent:

			keys := sdl.GetKeyboardState()
			if !resoudre {
				// Insertion des chiffres de l'utilisateur
				if keys[sdl.SCANCODE_KP_1] == 1 || keys[sdl.SCANCODE_1] == 1 {
					changeCase(1, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_2] == 1 || keys[sdl.SCANCODE_2] == 1 {
					changeCase(2, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_3] == 1 || keys[sdl.SCANCODE_3] == 1 {
					changeCase(3, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_4] == 1 || keys[sdl.SCANCODE_4] == 1 {
					changeCase(4, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_5] == 1 || keys[sdl.SCANCODE_5] == 1 {
					changeCase(5, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_6] == 1 || keys[sdl.SCANCODE_6] == 1 {
					changeCase(6, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_7] == 1 || keys[sdl.SCANCODE_7] == 1 {
					changeCase(7, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_8] == 1 || keys[sdl.SCANCODE_8] == 1 {
					changeCase(8, grille, masque)
				}
				if keys[sdl.SCANCODE_KP_9] == 1 || keys[sdl.SCANCODE_9] == 1 {
					changeCase(9, grille, masque)
				}

				// supprimer une case
				if keys[sdl.SCANCODE_BACKSPACE] == 1 || keys[sdl.SCANCODE_DELETE] == 1 {
					changeCase(0, grille, masque)
				}

				// si c'est un sudoku 16*16 il faut qu'il puisse insérer des lettres
				if size == 16 {
					if keys[sdl.SCANCODE_A] == 1 || keys[sdl.SCANCODE_Q] == 1 {
						changeCase(10, grille, masque)
					}
					if keys[sdl.SCANCODE_B] == 1 {
						changeCase(11, grille, masque)
					}
					if keys[sdl.SCANCODE_C] == 1 {
						changeCase(12, grille, masque)
					}
					if keys[sdl.SCANCODE_D] == 1 {
						changeCase(13, grille, masque)
					}
					if keys[sdl.SCANCODE_E] == 1 {
						changeCase(14, grille, masque)
					}
					if keys[sdl.SCANCODE_F] == 1 {
						changeCase(15, grille, masque)
					}
					if keys[sdl.SCANCODE_G] == 1 {
						changeCase(16, grille, masque)
					}

				}

			}

			// Bouger la case avec les flèches du clavier
			if keys[sdl.SCANCODE_LEFT] == 1 {
				selected_x -= 1
				if selected_x < 0 {
					selected_x = 0
				}
			}
			if keys[sdl.SCANCODE_RIGHT] == 1 {
				selected_x += 1
				if selected_x >= int32(size) {
					selected_x = int32(size) - 1
				}
			}
			if keys[sdl.SCANCODE_UP] == 1 {
				selected_y -= 1
				if selected_y < 0 {
					selected_y = 0
				}
			}
			if keys[sdl.SCANCODE_DOWN] == 1 {
				selected_y += 1
				if selected_y >= int32(size) {
					selected_y = int32(size) - 1
				}
			}

		}
	}

	renderer.Clear()

}

func menu() {

	messageGame = ""

	// init de la police des textes
	font, err := ttf.OpenFont("assets/Acme-Regular.ttf", int(17))
	if err != nil {
		panic(err)
	}
	defer font.Close()

	// la transparence
	if err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		panic(err)
	}

	// init de l'image bouton Facile
	facileImg, err := img.Load("assets/Facile.png")
	if hooverButtonMenu[0] {
		facileImg, err = img.Load("assets/Facile Hoover.png")
	}
	defer facileImg.Free()
	if err != nil {
		panic(err)
	}

	// init de l'image bouton Moyen
	moyenImg, err := img.Load("assets/Moyen.png")
	if hooverButtonMenu[1] {
		moyenImg, err = img.Load("assets/Moyen Hoover.png")
	}
	defer moyenImg.Free()
	if err != nil {
		panic(err)
	}

	// init de l'image bouton Difficile
	difficileImg, err := img.Load("assets/Difficile.png")
	if hooverButtonMenu[2] {
		difficileImg, err = img.Load("assets/Difficile Hoover.png")
	}
	if err != nil {
		panic(err)
	}
	defer difficileImg.Free()

	// init de l'image bouton Diabolique
	diaboliqueImg, err := img.Load("assets/Diabolique.png")
	if hooverButtonMenu[3] {
		diaboliqueImg, err = img.Load("assets/Diabolique Hoover.png")
	}
	if err != nil {
		panic(err)
	}
	defer diaboliqueImg.Free()

	// init de l'image bouton Continuer
	continuerImg, err := img.Load("assets/Continuer.png")
	if hooverButtonMenu[4] {
		continuerImg, err = img.Load("assets/Continuer Hoover.png")
	}
	if err != nil {
		panic(err)
	}
	defer continuerImg.Free()

	// affichage du bouton Facile
	bouton, err := renderer.CreateTextureFromSurface(facileImg)
	if err != nil {
		panic(err)
	}
	renderer.Copy(bouton, nil, &sdl.Rect{X: 100, Y: 100, W: int32(float64(facileImg.W)), H: int32(float64(facileImg.H))})
	bouton.Destroy()

	// affichage du bouton Moyen
	bouton, err = renderer.CreateTextureFromSurface(moyenImg)
	if err != nil {
		panic(err)
	}
	renderer.Copy(bouton, nil, &sdl.Rect{X: 100, Y: 200, W: int32(float64(moyenImg.W)), H: int32(float64(moyenImg.H))})
	bouton.Destroy()

	// affichage du bouton Difficile
	bouton, err = renderer.CreateTextureFromSurface(difficileImg)
	if err != nil {
		panic(err)
	}
	renderer.Copy(bouton, nil, &sdl.Rect{X: 100, Y: 300, W: int32(float64(difficileImg.W)), H: int32(float64(difficileImg.H))})
	bouton.Destroy()

	// affichage du bouton Diabolique
	bouton, err = renderer.CreateTextureFromSurface(diaboliqueImg)
	if err != nil {
		panic(err)
	}
	renderer.Copy(bouton, nil, &sdl.Rect{X: 100, Y: 400, W: int32(float64(diaboliqueImg.W)), H: int32(float64(diaboliqueImg.H))})
	bouton.Destroy()

	// affichage du bouton Continuer
	bouton, err = renderer.CreateTextureFromSurface(continuerImg)
	if err != nil {
		panic(err)
	}
	renderer.Copy(bouton, nil, &sdl.Rect{X: 100, Y: 500, W: int32(float64(continuerImg.W)), H: int32(float64(continuerImg.H))})
	bouton.Destroy()

	if messageMenu != "" {
		fontSurface, err = font.RenderUTF8Blended(messageMenu, sdl.Color{0, 0, 0, 255})
		if err != nil {
			panic(err)
		}

		fontTexture, err := renderer.CreateTextureFromSurface(fontSurface)
		if err != nil {
			panic(err)
		}

		fontSurface.Free()

		_, _, w, h, err := fontTexture.Query()
		if err != nil {
			panic(err)
		}

		renderer.Copy(fontTexture, nil, &sdl.Rect{X: continuerImg.W + 125, Y: 525, W: w, H: h})
		fontTexture.Destroy()
	}

	renderer.Present()

	// évènements
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch event.(type) {

		// fermeture de la fenêtre
		case *sdl.QuitEvent:
			running = false

		// mouvement de la souris
		case *sdl.MouseMotionEvent:
			x, y, _ := sdl.GetMouseState()

			// affichage des hoover si la souris est sur les boutons
			hooverButtonMenu[0] = (100+facileImg.W >= x && x >= 100) && (100+facileImg.H >= y && y >= 100)
			hooverButtonMenu[1] = (100+moyenImg.W >= x && x >= 100) && (200+moyenImg.H >= y && y >= 200)
			hooverButtonMenu[2] = (100+difficileImg.W >= x && x >= 100) && (300+difficileImg.H >= y && y >= 300)
			hooverButtonMenu[3] = (100+diaboliqueImg.W >= x && x >= 100) && (400+diaboliqueImg.H >= y && y >= 400)
			hooverButtonMenu[4] = (100+continuerImg.W >= x && x >= 100) && (500+continuerImg.H >= y && y >= 500)

		// clic de la souris
		case *sdl.MouseButtonEvent:
			mouseEvent := event.(*sdl.MouseButtonEvent)

			// clic gauche
			if mouseEvent.State == sdl.PRESSED && mouseEvent.Button == sdl.BUTTON_LEFT {
				x, y, _ := sdl.GetMouseState()

				// clic sur facile : génération d'une grille facile et entrée en jeu
				if (100+facileImg.W >= x && x >= 100) && (100+facileImg.H >= y && y >= 100) {

					difficulte = 1
					screen = Loading
				}

				// clic sur moyen : génération d'une grille moyen et entrée en jeu
				if (100+moyenImg.W >= x && x >= 100) && (200+moyenImg.H >= y && y >= 200) {
					difficulte = 2
					screen = Loading

				}

				// clic sur difficile : génération d'une grille difficile et entrée en jeu
				if (100+difficileImg.W >= x && x >= 100) && (300+difficileImg.H >= y && y >= 300) {
					difficulte = 3
					screen = Loading

				}

				// clic sur diabolique : génération d'une grille diabolique et entrée en jeu
				if (100+diaboliqueImg.W >= x && x >= 100) && (400+diaboliqueImg.H >= y && y >= 400) {
					screen = Loading
					difficulte = 4
				}

				// clic sur continuer : récupération de la sauvegarde (ou simplement du txt) et entrée en jeu
				// erreur si pas de txt
				if (100+continuerImg.W >= x && x >= 100) && (500+continuerImg.H >= y && y >= 500) {

					grille, masque, err = utils.ImportFile()

					if err != nil {

						messageMenu = "Erreur avec la sauvegarde/le txt. Vérifiez le fichier file/save.txt"
					} else {
						possibilite = utils.Generer_possibilite(&grille)
						verifier = utils.FindErrors(&grille, &masque)
						startTime = time.Now()
						screen = Game
					}

				}
			}
		}
	}

	renderer.Clear()

}

func loading() {

	resoudre = false
	renderer.SetDrawColor(0, 0, 0, 255)

	// init de la police des textes
	font, err := ttf.OpenFont("assets/Acme-Regular.ttf", int(taille_case/2))
	if err != nil {
		panic(err)
	}
	defer font.Close()

	// init de la surface des textes
	switch difficulte {
	case 1:
		fontSurface, err = font.RenderUTF8Blended("Chargement de la grille Facile...", sdl.Color{R: 0, G: 0, B: 0, A: 255})
	case 2:
		fontSurface, err = font.RenderUTF8Blended("Chargement de la grille Moyenne...", sdl.Color{R: 0, G: 0, B: 0, A: 255})
	case 3:
		fontSurface, err = font.RenderUTF8Blended("Chargement de la grille Difficile...", sdl.Color{R: 0, G: 0, B: 0, A: 255})
	case 4:
		fontSurface, err = font.RenderUTF8Blended("Chargement de la grille Diabolique...", sdl.Color{R: 0, G: 0, B: 0, A: 255})
	}
	if err != nil {
		panic(err)
	}

	fontTexture, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		panic(err)
	}
	fontSurface.Free()

	// dimensions du texte
	_, _, w, h, err := fontTexture.Query()
	if err != nil {
		panic(err)
	}

	renderer.Copy(fontTexture, nil, &sdl.Rect{X: windowVertical/2 - h/2, Y: windowHorizontal/2 - w/2, W: w, H: h})
	fontTexture.Destroy()
	renderer.Present()

	switch difficulte {
	case 1:
		grille, possibilite, masque, verifier = algo.Init_grille(1)
	case 2:
		grille, possibilite, masque, verifier = algo.Init_grille(2)
	case 3:
		grille, possibilite, masque, verifier = algo.Init_grille(3)
	case 4:
		grille, possibilite, masque, verifier = algo.Init_grille(4)

	}

	startTime = time.Now()
	screen = Game

}

func main() {

	start()

	// définition de la taille des cases selon le format grille
	if size == 16 {
		taille_case = 30
	} else if size == 9 {
		taille_case = 50
	}

	screen = Menu

	// boucle de jeu
	for running {

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		switch screen {
		// ecran de jeu
		case Game:
			interface_jeu(&grille)
		// ecran de menu
		case Menu:
			menu()
		case Loading:
			loading()
		}

		// delai pour éviter que trop de frames soit générées et que il y ait du lag
		sdl.Delay(7)

	}

	destroy()

}
