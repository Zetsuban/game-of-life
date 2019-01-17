package entities

// Cell structure yes
type Cell struct {
	alive    bool
	nbNearby int
}

// Str : return string
func (c Cell) Str() string {
	if c.Alive() {
		return "*"
	}
	return " "
}

// Rune : return rune for termbox
func (c Cell) Rune() rune {
	if c.Alive() {
		return ' '
	} else {
		return ' '
	}
}

// SetAlive : sets the state of alive
func (c *Cell) SetAlive(alive bool) {
	c.alive = alive
}

// Alive : Return the state of cell
func (c Cell) Alive() bool {
	return c.alive
}

// NextGen : Sets cell state based on nearby
func (c *Cell) NextGen() {
	switch c.alive {
	case true:
		switch c.nbNearby {
		case 0, 1:
			c.alive = false
		case 4, 5, 6, 7, 8:
			c.alive = false
		default:
			break
		}
	case false:
		if c.nbNearby == 3 {
			c.alive = true
		}
	}
}

// Copy the cell
func (c Cell) Copy() Cell {
	var newCell Cell
	newCell.SetAlive(c.Alive())
	return newCell
}

// Step create copy of itself then set alive and dead status
func (c Cell) Step(neighbors []Cell) Cell {
	var liveCount int

	for _, n := range neighbors {
		if n.Alive() {
			liveCount = liveCount + 1
		}
	}

	newCell := c.Copy()

	if c.Alive() {
		if liveCount < 2 || liveCount > 3 {
			newCell.SetAlive(false)
		}
	} else {
		if !c.Alive() && liveCount == 3 {
			newCell.SetAlive(true)
		}
	}

	return newCell
}
