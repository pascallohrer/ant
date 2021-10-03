package main

type Color uint8
type Coordinate struct{ x, y int }

const (
	COLOR_UNDEFINED = Color(0)
	COLOR_WHITE     = Color(1)
	COLOR_BLACK     = Color(2)
)

type Field struct {
	cells       map[Coordinate]Color
	ant         *ant
	transitions map[Color]Color
}

type modifier struct {
	x, y     int
	newColor Color
}

func (f *Field) Next(results chan<- modifier) {
	currentColor := f.getAntCellColor()
	results <- f.setAntCellColor(f.transitions[currentColor])
	f.ant.move(currentColor)
	if f.getAntCellColor() == COLOR_UNDEFINED {
		results <- f.setAntCellColor(COLOR_WHITE)
	}
}

func (f *Field) getAntCellColor() Color {
	return f.cells[Coordinate{f.ant.x, f.ant.y}]
}

func (f *Field) setAntCellColor(newColor Color) modifier {
	f.cells[Coordinate{f.ant.x, f.ant.y}] = newColor
	return modifier{
		f.ant.x,
		f.ant.y,
		newColor,
	}
}

func (f *Field) GetCells() map[Coordinate]Color {
	return f.cells
}
