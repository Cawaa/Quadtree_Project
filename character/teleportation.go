package character

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

/// EXTENSION 4 : téléportation ///

// teleport permet à un personnage de se téléporter entre deux portails.
// La fonction vérifie si la touche 'T' est pressée pour définir les positions des portails.
// Une fois les portails définis, la téléportation est autorisée si le personnage n'est pas déjà sur un portail.
// Si le personnage se trouve sur un portail et que la téléportation est autorisée, il est téléporté à l'autre portail.
// La fonction retourne un booléen indiquant si la téléportation a eu lieu.
func (c *Character) teleport() (teleported bool) {

	// Variable indiquant si on s'est téléporté //
	teleported=false

	// Défini les positions des portails //
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		if c.portail_actif==0 {
			c.portail_depart_x = c.X
			c.portail_depart_y = c.Y
			c.portail_actif = 1
			fmt.Println("Vous avez un téléporteur en X :",c.portail_depart_x,", Y :",c.portail_depart_y)
		} else if c.portail_actif==1 && (c.portail_depart_x!=c.X || c.portail_depart_y!=c.Y) {
			c.portail_arrivee_x = c.X
			c.portail_arrivee_y = c.Y
			c.portail_actif++
			fmt.Println("Vous avez votre second téléporteur en X :",c.portail_arrivee_x,", Y :",c.portail_arrivee_y)
		} else if c.portail_actif==2 && (c.portail_arrivee_x!=c.X || c.portail_arrivee_y!=c.Y) {
			c.portail_depart_x=c.portail_arrivee_x
			c.portail_depart_y=c.portail_arrivee_y
			c.portail_arrivee_x=c.X
			c.portail_arrivee_y=c.Y
			c.tp_autorise=false
			fmt.Println("Vos téléporteurs sont maintenant en X :",c.portail_depart_x,", Y :",c.portail_depart_y, "et en X :",c.portail_arrivee_x,", Y :",c.portail_arrivee_y)
		}
	}

	// Si tous les portails sont déjà posés et qu'on ne vient pas d'en poser un juste avant, on autorise la téléportation //
	if c.portail_actif<2 {
		c.tp_autorise=false
	} else if (c.portail_arrivee_x!=c.X || c.portail_arrivee_y!=c.Y) && (c.portail_depart_x!=c.X || c.portail_depart_y!=c.Y) {
		c.tp_autorise=true
	}

	// Gère la téléportation d'un portail à l'autre //
	if (c.X == c.portail_arrivee_x) && (c.Y == c.portail_arrivee_y) && c.tp_autorise {
		c.X = c.portail_depart_x
		c.Y = c.portail_depart_y
		c.tp_autorise=false
		teleported=true
	} else if (c.X == c.portail_depart_x) && (c.Y == c.portail_depart_y) && c.tp_autorise {
		c.X = c.portail_arrivee_x
		c.Y = c.portail_arrivee_y
		c.tp_autorise=false
		teleported=true
	}

	return teleported
}