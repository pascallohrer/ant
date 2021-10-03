# ant

This project implements [Langton's Ant][https://en.wikipedia.org/wiki/Langton%27s_ant].

# Usage

`LangtonsAnt(initialCells)` creates a field with the default Langton's Ant settings. The parameter is a `map[Coordinate]Color`. All cells not defined in the map are considered white. For an empty initial field, pass nil. The first step of the automaton will then turn the center cell white and not move the ant, which always starts at `Coordinate{0,0}` facing north. This is a robustness fallback feature, the best way to use the function is passing a populated map with at least the starting cell set to a defined color.

Call `field.Next(results)` to advance the automaton one step. The results (changed cells) will be sent to the channel of type `chan modifier` passed as a parameter. See the `ant_test.go` file for a simple dummy channel. This channel can be used to update a representation of the current field. For a snapshot of the entire field at any point, call `field.GetCells()`. For a running simulation, it is recommended to use the updates from the channel as these include changes cells only, while calling `GetCells` after every step would keep retrieving the entire data set, which is inefficient.

# Custom rules

Ants can be passed arbitrary movement strategies of type `func(currentCell Color, currentForward Direction) (dx, dy int, newForward Direction)`. They are called with the current ant cell's color and the ant's current "forward" facing direction and they must return the number of cells moved in the x and y direction as well as the ant's "forward" facing direction after the move. A fallback for "undefined" (zero) parameter values is highly recommended. Fields can be passed arbitrary color transitions of type `map[Color]Color`. This map will be used to determine a cell's color after a step was evaluated with the ant on it.

`CustomAnt(initialCells, antMovementStrategy, colorTransitions)` is used to setup a field with these custom properties.

Multiple colors and other directions (such as hexagonal grids) can be used even when they're not defined as constants, those are just for convenience. It is recommended to define appropriate constants of types `Color` or `Direction` in the scope of the strategy and transition definitions, respectively, instead of using numeric values.
