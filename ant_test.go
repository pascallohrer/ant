package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

type TestCases map[int]map[Coordinate]Color

func TestLangtonsAnt(t *testing.T) {
	field := LangtonsAnt(nil)
	dummyChan := make(chan modifier)
	go func() modifier {
		dummy, ok := modifier{}, true
		for ok {
			dummy, ok = <-dummyChan
		}
		return dummy
	}()
	field.Next(dummyChan) // initialize field to contain a white cell at the ant
	sample := TestCases{
		0: map[Coordinate]Color{
			{0, 0}: COLOR_WHITE,
		},
		1: map[Coordinate]Color{
			{0, 0}: COLOR_BLACK,
			{1, 0}: COLOR_WHITE,
		},
		4: map[Coordinate]Color{
			{0, 0}:  COLOR_BLACK,
			{1, 0}:  COLOR_BLACK,
			{1, -1}: COLOR_BLACK,
			{0, -1}: COLOR_BLACK,
		},
		8: map[Coordinate]Color{
			{0, 0}:  COLOR_WHITE,
			{1, 0}:  COLOR_BLACK,
			{1, -1}: COLOR_BLACK,
			{0, -1}: COLOR_BLACK,
			{-1, 0}: COLOR_BLACK,
			{-1, 1}: COLOR_BLACK,
			{0, 1}:  COLOR_BLACK,
		},
		12: map[Coordinate]Color{
			{0, 0}:   COLOR_BLACK,
			{1, 0}:   COLOR_BLACK,
			{1, -1}:  COLOR_BLACK,
			{0, -1}:  COLOR_BLACK,
			{-1, 0}:  COLOR_WHITE,
			{-1, 1}:  COLOR_BLACK,
			{0, 1}:   COLOR_BLACK,
			{-2, 0}:  COLOR_WHITE,
			{-2, -1}: COLOR_BLACK,
			{-1, -1}: COLOR_BLACK,
		},
	}
	for i := range [12]bool{} {
		expected, exists := sample[i]
		if exists {
			assert.Equal(t, fmt.Sprint(expected), fmt.Sprint(field.GetCells()))
		}
		field.Next(dummyChan)
	}
}
