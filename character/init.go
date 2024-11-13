package character

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init(floorWidth, floorHeight int) {
	c.animationStep = 1

	c.X = floorWidth / 2
	c.Y = floorHeight / 2
}
