package camera

// Init met en place une caméra.
func (c *Camera) Init(characterPosX, characterPosY int) {
	c.X = characterPosX
	c.Y = characterPosY
}
