package main

import (
	"strings"
)

type Stage struct {
	initialWorld     [][]WorldCellContent
	initialSnakeHead Position
	initialSnakeTail Position
	name             string
	speed            int
	boundary         bool
}

func getWorldCellContentArrayFromString(strWorld string) [][]WorldCellContent {
	lines := strings.Split(strWorld, "\n")
	ret := make([][]WorldCellContent, len(lines))
	for i := range lines {
		line := strings.Split(lines[i], "")
		ret[i] = make([]WorldCellContent, len(line))
		for j := range line {
			ret[i][j] = getCellContentByChar(line[j])
		}
	}
	return ret
}

var stages = []Stage{
	{
		initialSnakeHead: Position{y: 1, x: 4},
		initialSnakeTail: Position{y: 1, x: 2},
		initialWorld: getWorldCellContentArrayFromString(
			"          \n" +
				"  >>>     \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          ",
		),
		name:     "The begining",
		speed:    500,
		boundary: false,
	},
	{
		initialSnakeHead: Position{y: 1, x: 4},
		initialSnakeTail: Position{y: 1, x: 2},
		initialWorld: getWorldCellContentArrayFromString(
			"          \n" +
				"  >>>     \n" +
				"          \n" +
				"          \n" +
				"XXXXX     \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          ",
		),
		name:     "A wall",
		speed:    500,
		boundary: true,
	},
}
