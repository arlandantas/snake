package main

import (
	"strings"
)

type Stage struct {
	initialWorld     [][]WorldCellContent
	initialSnakeHead Position
	initialSnakeTail Position
	name             string
	initialSpeed     int
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
		initialSpeed: 500,
		name:         "The begining",
		boundary:     false,
	},
	{
		initialSnakeHead: Position{y: 1, x: 4},
		initialSnakeTail: Position{y: 1, x: 2},
		initialWorld: getWorldCellContentArrayFromString(
			"          \n" +
				"  >>>     \n" +
				"          \n" +
				"          \n" +
				"   XXXX   \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          \n" +
				"          ",
		),
		initialSpeed: 500,
		name:         "A wall",
		boundary:     true,
	},
	{
		initialSnakeHead: Position{y: 10, x: 10},
		initialSnakeTail: Position{y: 10, x: 8},
		initialWorld: getWorldCellContentArrayFromString(
			"                    \n" +
				"                    \n" +
				"  XXXX        XXXX  \n" +
				"  X              X  \n" +
				"  X              X  \n" +
				"  X              X  \n" +
				"                    \n" +
				"                    \n" +
				"      XXXXXXXX      \n" +
				"                    \n" +
				"        >>>         \n" +
				"                    \n" +
				"      XXXXXXXX      \n" +
				"                    \n" +
				"  X              X  \n" +
				"  X              X  \n" +
				"  X              X  \n" +
				"  XXXX        XXXX  \n" +
				"                    \n" +
				"                    ",
		),
		initialSpeed: 500,
		name:         "Many walls",
		boundary:     false,
	},
}
