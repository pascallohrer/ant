package main

type Direction uint8

const (
	DIR_UNDEFINED = Direction(0)
	DIR_NORTH     = Direction(1)
	DIR_EAST      = Direction(2)
	DIR_SOUTH     = Direction(3)
	DIR_WEST      = Direction(4)
)

type ant struct {
	x, y     int
	forward  Direction
	strategy func(currentCell Color, currentForward Direction) (dx, dy int, newForward Direction)
}

func (a *ant) move(currentCell Color) {
	motionX, motionY, rotation := a.strategy(currentCell, a.forward)
	a.x += motionX
	a.y += motionY
	a.forward = rotation
}
