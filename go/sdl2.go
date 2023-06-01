package main

import (
	"fmt"
	"math"

	// pour voir le temps de résolution
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	TAILLE = 9 // Taille de la grille de sudoku
	Game   = iota
	Menu
	Difficulty
)

type mouseState struct {
	leftButton  bool
	rightButton bool
	x, y        int32
}

func mouseUpdate() mouseState {
	var souris mouseState
	x, y, mouseButtonState := sdl.GetMouseState()
	souris.x = x
	souris.y = y
	leftButton := mouseButtonState & sdl.ButtonLMask()
	rightButton := mouseButtonState & sdl.ButtonRMask()

	souris.leftButton = !(leftButton == 0)
	souris.rightButton = !(rightButton == 0)

	return souris

}

// structure de couleur : rouge, vert, bleu, opacité
type color struct {
	r, g, b, a byte
}

// variable ici pour que TOUT LE MONDE puisse les lire
var (
	largeur                int32 = 600
	hauteur                int32 = 800
	renderer               *sdl.Renderer
	window                 *sdl.Window
	tex                    *sdl.Texture
	masque                 [TAILLE][TAILLE]bool
	verifier               [TAILLE][TAILLE]bool
	titre                  string    = "Sudoku!"
	black                  sdl.Color = sdl.Color{0, 0, 0, 255}
	white                  sdl.Color = sdl.Color{255, 255, 255, 255}
	taille_case            int32
	mouse                  mouseState
	lastMouse              mouseState
	running                bool = true
	screen                 int
	fontMedium             *ttf.Font
	fontSurface            *sdl.Surface
	selectionCase          bool = false
	selected_x, selected_y int32
)

func start() {

	// init
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	// windows
	window, err := sdl.CreateWindow(titre, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, hauteur, largeur, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	// renderer  : pour dessiner dessus
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	// // texture
	// tex, err = renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, hauteur, largeur)
	// if err != nil {
	// 	panic(err)
	// }

}

func changeCase(valeur int, grille *[TAILLE + 2][TAILLE + 1]int, masque [TAILLE][TAILLE]bool) {
	if !masque[selected_x][selected_y] {
		grille[selected_x][selected_y] = valeur
	}
}

func destroy() {
	window.Destroy()
	// renderer.Destroy()		// fait tout ralentir ??
	tex.Destroy()

	ttf.Quit()
	sdl.Quit()
}

func interface_jeu(grille *[TAILLE + 2][TAILLE + 1]int) {

	img, err := img.Load("assets/button.png")
	if err != nil {
		panic(err)
	}

	fontMedium, err := ttf.OpenFont("assets/ARIAL.TTF", int(taille_case/2))
	if err != nil {
		panic(err)
	}

	fontSurface, err := fontMedium.RenderUTF8Blended("text", sdl.Color{0, 255, 0, 255})
	if err != nil {
		panic(err)
	}

	fontTexture, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		panic(err)
	}

	_, _, w, h, err := fontTexture.Query()
	if err != nil {
		panic(err)
	}

	//? libérer l'espace mémoire ???

	var rect sdl.Rect

	bouton, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(err)
	}
	coordBouton := [4]int32{(TAILLE + 2) * taille_case, 3 * taille_case, int32(float64(img.W) * 0.4), int32(float64(img.H) * 0.4)} // coin haut gauche x, y, largeur, hauteur

	rect = sdl.Rect{coordBouton[0], coordBouton[1], coordBouton[2], coordBouton[3]}
	renderer.Copy(bouton, nil, &rect)

	renderer.SetDrawColor(black.R, black.G, black.B, black.A)
	for X := 0; X < TAILLE; X++ {
		for Y := 0; Y < TAILLE; Y++ {

			if verifier[X][Y] {
				fmt.Print("hefrgh")
				renderer.SetDrawColor(255, 0, 0, 255)
				rect = sdl.Rect{int32(X) * taille_case, int32(Y) * taille_case, taille_case, taille_case}
				renderer.FillRect(&rect)
				renderer.SetDrawColor(black.R, black.G, black.B, black.A)
			}

			rect = sdl.Rect{int32(X) * taille_case, int32(Y) * taille_case, taille_case, taille_case}
			renderer.DrawRect(&rect)

		}
	}

	// dessiner les lignes verticales/horizontales
	renderer.SetDrawColor(0, 0, 255, 255)
	for i := 1; i <= 2; i++ {
		renderer.FillRect(&sdl.Rect{taille_case*int32(i)*int32(math.Sqrt(TAILLE)) - 1, 0, 3, taille_case * int32(TAILLE)})
		renderer.FillRect(&sdl.Rect{0, taille_case*int32(i)*int32(math.Sqrt(TAILLE)) - 1, taille_case * int32(TAILLE), 3})
	}

	if selectionCase {
		renderer.SetDrawColor(255, 0, 0, 255)
		rect = sdl.Rect{selected_x * taille_case, selected_y * taille_case, taille_case, taille_case}
		renderer.FillRect(&rect)
		renderer.SetDrawColor(white.R, white.G, white.B, white.A)
		rect = sdl.Rect{selected_x*taille_case + 2, selected_y*taille_case + 2, taille_case - 4, taille_case - 4}
		renderer.FillRect(&rect)
	}

	renderer.SetDrawColor(black.R, black.G, black.B, black.A)
	for X := 0; X < TAILLE; X++ {
		for Y := 0; Y < TAILLE; Y++ {

			if grille[X][Y] != 0 {

				if grille[X][Y] > 9 {
					fontSurface, err = fontMedium.RenderUTF8Blended(string(grille[X][Y]-10+65), black)
				} else {
					fontSurface, err = fontMedium.RenderUTF8Blended(string(grille[X][Y]+48), black)
				}
				if err != nil {
					panic(err)
				}

				fontTexture, err = renderer.CreateTextureFromSurface(fontSurface)
				if err != nil {
					panic(err)
				}

				_, _, w, h, err = fontTexture.Query()
				if err != nil {
					panic(err)
				}

				renderer.Copy(fontTexture, nil, &sdl.Rect{int32(X)*taille_case + (taille_case-w)/2, int32(Y)*taille_case + (taille_case-h)/2, w, h})

			}

		}
	}

	renderer.Present()

	// evenements
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch event.(type) { // case/switch sur le type de l'event (ce sont tjr des pointeurs)

		case *sdl.QuitEvent:
			running = false

		case *sdl.KeyboardEvent:

			keys := sdl.GetKeyboardState()

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
			if keys[sdl.SCANCODE_BACKSPACE] == 1 || keys[sdl.SCANCODE_DELETE] == 1 {
				changeCase(0, grille, masque)
			}

			if TAILLE == 16 {
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

			if keys[sdl.SCANCODE_LEFT] == 1 {
				selected_x -= 1
				if selected_x < 0 {
					selected_x = 0
				}
			}
			if keys[sdl.SCANCODE_RIGHT] == 1 {
				selected_x += 1
				if selected_x >= TAILLE {
					selected_x = TAILLE - 1
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
				if selected_y >= TAILLE {
					selected_y = TAILLE - 1
				}
			}

		}
	}

	// evenements souris
	if !mouse.leftButton && lastMouse.leftButton {
		if (0 <= mouse.x && mouse.x <= TAILLE*taille_case) && (0 <= mouse.y && mouse.y <= TAILLE*taille_case) {
			selectionCase = true
			selected_x, selected_y = mouse.x/taille_case, mouse.y/taille_case
		} else {
			selectionCase = false
		}

		if (coordBouton[0]+coordBouton[2] >= mouse.x && mouse.x >= coordBouton[0]) && (coordBouton[1]+coordBouton[3] >= mouse.y && mouse.y >= coordBouton[1]) {
			fmt.Print("fghjkl")
		}
	} else if !mouse.rightButton && lastMouse.rightButton {
		fmt.Println("droite")
	}

	// fontSurface.Free()
	// fontTexture.Destroy()
	// fontMedium.Close()

}

func main() {

	grille, _, _, verifier := init_grille(Grille_sudoku_exemple(TAILLE)) // REMPLACER _ PAR masque QUAND ON L'UTILISE

	verifier[2][2] = true

	fmt.Println("GRILLE DE DEPART : ")
	print_grille(&grille, false)
	fmt.Println()

	Grille9 := [9 + 2][9 + 1]int{{0, 0, 0, 0, 4, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 0, 0},
		{0, 3, 2, 1, 9, 0, 0, 8, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 4, 0},
		{0, 0, 4, 9, 0, 0, 0, 6, 1, 0},
		{0, 2, 9, 0, 7, 4, 0, 0, 0, 0},
		{0, 9, 6, 7, 0, 1, 5, 0, 0, 0},
		{5, 0, 0, 0, 3, 0, 0, 2, 7, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	start()

	if TAILLE == 16 {
		taille_case = 30
	} else if TAILLE == 9 {
		taille_case = 50
	}

	screen = Game

	mouse = mouseUpdate()
	lastMouse = mouse

	for running {

		// before := time.Now()

		mouse = mouseUpdate()

		renderer.SetDrawColor(white.R, white.G, white.B, white.A)
		renderer.Clear()

		switch screen {
		case Game:
			interface_jeu(&Grille9)
		}

		lastMouse = mouse

		// after := time.Now()
		// temps := after.Sub(before)
		// fmt.Println(temps.Milliseconds(), "MILLISECONDES.")

		sdl.Delay(10)

	}

	destroy()

}

//https://github.com/veandco/go-sdl2-examples/blob/master/examples/drawing-text/drawing-text.go
