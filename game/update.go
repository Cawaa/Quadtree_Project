package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"os"
	"fmt"
	"bufio"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)
	
	
	if configuration.Global.Zoom{
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			configuration.Global.ScreenHeight -= configuration.Global.TileSize * configuration.Global.NumTileX / 2
			configuration.Global.ScreenWidth -= configuration.Global.TileSize * configuration.Global.NumTileY / 2
			configuration.Global.ScreenCenterTileX -= configuration.Global.TileSize / 6
			configuration.Global.ScreenCenterTileY -= configuration.Global.TileSize / 6
		}
	


		if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
			configuration.Global.ScreenHeight += configuration.Global.TileSize * configuration.Global.NumTileX / 2
			configuration.Global.ScreenWidth += configuration.Global.TileSize * configuration.Global.NumTileY / 2
			configuration.Global.ScreenCenterTileX += configuration.Global.TileSize / 6
			configuration.Global.ScreenCenterTileY += configuration.Global.TileSize / 6
		}
	}
	if configuration.Global.InfiniteGenExtension {
		widhtFile, heightFile := getWidhtHeightOfFile(configuration.Global.FloorFile)
		if g.character.X == 0 {
			g.floor.GenerateNewChunk(2)
			if configuration.Global.RandomGenExtension {
				g.character.X += configuration.Global.RandomWidht
				g.camera.X += configuration.Global.RandomWidht
			} else {
				g.character.X += widhtFile
				g.camera.X += widhtFile
			}
		}
		if g.character.Y == 0 {
			g.floor.GenerateNewChunk(3)
			if configuration.Global.RandomGenExtension {
				g.character.Y += configuration.Global.RandomHeight
				g.camera.Y += configuration.Global.RandomHeight
			} else {
				g.character.Y += heightFile
				g.camera.Y += heightFile
			}

		}
		if g.character.X == g.floor.GetWidthQuad()-1 {
			g.floor.GenerateNewChunk(0)
		}

		if g.character.Y == g.floor.GetHeightQuad()-1 {
			g.floor.GenerateNewChunk(1)
		}

	}

	return nil
}
func getWidhtHeightOfFile(fileName string) (width, height int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ligneCount int
	var colonneCount int
	for scanner.Scan() {
		ligneCount++
		ligne := scanner.Text()
		if ligneCount == 1 {
			colonneCount = len(ligne)
		}
	}
	return colonneCount, ligneCount
}
