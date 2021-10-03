package main

func LangtonsAnt(initialCells map[Coordinate]Color) *Field {
	antStrategy := func(currentCell Color, currentForward Direction) (dx int, dy int, newForward Direction) {
		switch true {
		case currentCell == COLOR_WHITE && currentForward == DIR_NORTH:
			return 1, 0, DIR_EAST
		case currentCell == COLOR_WHITE && currentForward == DIR_EAST:
			return 0, -1, DIR_SOUTH
		case currentCell == COLOR_WHITE && currentForward == DIR_SOUTH:
			return -1, 0, DIR_WEST
		case currentCell == COLOR_WHITE && currentForward == DIR_WEST:
			return 0, 1, DIR_NORTH
		case currentCell == COLOR_BLACK && currentForward == DIR_NORTH:
			return -1, 0, DIR_WEST
		case currentCell == COLOR_BLACK && currentForward == DIR_EAST:
			return 0, 1, DIR_NORTH
		case currentCell == COLOR_BLACK && currentForward == DIR_SOUTH:
			return 1, 0, DIR_EAST
		case currentCell == COLOR_BLACK && currentForward == DIR_WEST:
			return 0, -1, DIR_SOUTH
		default:
			return 0, 0, currentForward
		}
	}
	colorTransitions := map[Color]Color{
		COLOR_BLACK:     COLOR_WHITE,
		COLOR_WHITE:     COLOR_BLACK,
		COLOR_UNDEFINED: COLOR_WHITE,
	}
	return CustomAnt(initialCells, antStrategy, colorTransitions)
}

func CustomAnt(
	initialCells map[Coordinate]Color,
	antMovementStrategy func(currentCell Color, currentForward Direction) (dx int, dy int, newForward Direction),
	colorTransitions map[Color]Color,
) *Field {
	if initialCells == nil {
		initialCells = make(map[Coordinate]Color)
	}
	return &Field{
		cells: initialCells,
		ant: &ant{
			x:        0,
			y:        0,
			forward:  DIR_NORTH,
			strategy: antMovementStrategy,
		},
		transitions: colorTransitions,
	}
}
