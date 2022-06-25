package main

import (
	"fmt"
)

type WorldCellContent int
type SnakeMovementDirection int

const (
	WorldCellEmpty WorldCellContent = iota
	WorldCellSnakeMovingUp
	WorldCellSnakeMovingRight
	WorldCellSnakeMovingDown
	WorldCellSnakeMovingLeft
	WorldCellWall
	WorldCellFood
)

var SnakeWorldCellContents = []WorldCellContent{
	WorldCellSnakeMovingUp,
	WorldCellSnakeMovingRight,
	WorldCellSnakeMovingDown,
	WorldCellSnakeMovingLeft,
}

const (
	SnakeMoveUp SnakeMovementDirection = iota
	SnakeMoveRight
	SnakeMoveDown
	SnakeMoveLeft
)

var world [][]WorldCellContent

type Position struct {
	x int
	y int
}

var currentSnakeDirection = SnakeMoveRight
var currentSnakeHead = Position{y: 0, x: 0}
var currentSnakeTail = Position{y: 0, x: 0}
var isSnakeAlive = true
var currentScore = 0
var tickIntervalId = 0

func loadStage(stageIndex int) {
	stage := stages[stageIndex]
	world = stage.initialWorld
	currentSnakeHead = stage.initialSnakeHead
	currentSnakeTail = stage.initialSnakeTail
	worldHtml := ""
	for i := range world {
		worldHtml += "\t<div class=\"row\">\n"
		for j := range world[i] {
			worldHtml += fmt.Sprintf("\t\t<div class=\"cell\" id=\"cell%d%d\"></div>\n", i, j)
		}
		worldHtml += "</div>"
		div, err := getElementById("world")
		if err != nil {
			fmt.Printf("Failed to get world div: %s\n", err)
		} else {
			div.Set("innerHTML", worldHtml)
		}
	}
}

func setSnakeHeadDirection(direction SnakeMovementDirection) bool {
	currentDirection := world[currentSnakeHead.y][currentSnakeHead.x]
	if direction == SnakeMoveUp && currentDirection != WorldCellSnakeMovingDown {
		world[currentSnakeHead.y][currentSnakeHead.x] = WorldCellSnakeMovingUp
		return true
	} else if direction == SnakeMoveRight && currentDirection != WorldCellSnakeMovingLeft {
		world[currentSnakeHead.y][currentSnakeHead.x] = WorldCellSnakeMovingRight
		return true
	} else if direction == SnakeMoveDown && currentDirection != WorldCellSnakeMovingUp {
		world[currentSnakeHead.y][currentSnakeHead.x] = WorldCellSnakeMovingDown
		return true
	} else if direction == SnakeMoveLeft && currentDirection != WorldCellSnakeMovingRight {
		world[currentSnakeHead.y][currentSnakeHead.x] = WorldCellSnakeMovingLeft
		return true
	}
	return false
}

func moveSnake() {
	previousHeadDirection := world[currentSnakeHead.y][currentSnakeHead.x]
	switch previousHeadDirection {
	case WorldCellSnakeMovingRight:
		currentSnakeHead.x += 1
	case WorldCellSnakeMovingLeft:
		currentSnakeHead.x -= 1
	case WorldCellSnakeMovingUp:
		currentSnakeHead.y -= 1
	case WorldCellSnakeMovingDown:
		currentSnakeHead.y += 1
	}
	if currentSnakeHead.x < 0 || currentSnakeHead.x >= len(world[0]) ||
		currentSnakeHead.y < 0 || currentSnakeHead.y >= len(world) ||
		(world[currentSnakeHead.y][currentSnakeHead.x] != WorldCellEmpty &&
			world[currentSnakeHead.y][currentSnakeHead.x] != WorldCellFood) {
		isSnakeAlive = false
		return
	}
	if world[currentSnakeHead.y][currentSnakeHead.x] != WorldCellFood {
		previousTailDirection := world[currentSnakeTail.y][currentSnakeTail.x]
		world[currentSnakeTail.y][currentSnakeTail.x] = WorldCellEmpty
		switch previousTailDirection {
		case WorldCellSnakeMovingRight:
			currentSnakeTail.x += 1
		case WorldCellSnakeMovingLeft:
			currentSnakeTail.x -= 1
		case WorldCellSnakeMovingUp:
			currentSnakeTail.y -= 1
		case WorldCellSnakeMovingDown:
			currentSnakeTail.y += 1
		}
	} else {
		currentScore++
	}
	world[currentSnakeHead.y][currentSnakeHead.x] = previousHeadDirection
}

func printWorldString() {
	content := fmt.Sprintf("Score: %d\n|", currentScore)
	for range world[0] {
		content += "-"
	}
	content += "|\n"
	for i := range world {
		content += "|"
		for j := range world[i] {
			content += getCharByCellContent(world[i][j])
		}
		content += "|\n"
	}
	content += "|"
	for range world[0] {
		content += "-"
	}
	content += "|\n"
	fmt.Println(content)
}

func getCellId(y, x int) string {
	return fmt.Sprintf("cell%d%d", y, x)
}

func updateHtmlWorld() {
	for y := range world {
		for x := range world[y] {
			cell, err := getElementById(getCellId(y, x))
			if err != nil {
				fmt.Printf("Failed to get worldDiv: %s\n", err)
				break
			}
			cell.Set("className", getCellClasses(y, x))
		}
	}
	bMessage, err := getElementById("message")
	if err != nil {
		fmt.Printf("Failed to get bMessage: %s\n", err)
		return
	}
	if isSnakeAlive {
		bMessage.Set("innerHTML", fmt.Sprintf("Score: %d", currentScore))
	} else {
		bMessage.Set("innerHTML", fmt.Sprintf("GAME OVER! Score: %d", currentScore))
	}
}

func isValidContentType(haystack []WorldCellContent, needle WorldCellContent) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func getAround(value int, max int) []int {
	var ret []int
	if value > 1 {
		ret = append(ret, value-1)
	}
	ret = append(ret, value)
	if value < max-1 {
		ret = append(ret, value+1)
	}
	return ret
}

func getXs(x int) []int {
	return getAround(x, len(world[0]))
}

func getYs(y int) []int {
	return getAround(y, len(world))
}

func getCellClasses(y, x int) string {
	classes := "cell"
	cell := world[y][x]
	isNotOnTop := y > 0
	isNotOnBottom := y < len(world)-1
	isNotOnLeft := x > 0
	isNotOnRight := x < len(world[0])-1
	if isValidContentType(SnakeWorldCellContents, cell) {
		classes += " snake"
		if isNotOnTop && isValidContentType(SnakeWorldCellContents, world[y-1][x]) {
			classes += " top"
		}
		if isNotOnBottom && isValidContentType(SnakeWorldCellContents, world[y+1][x]) {
			classes += " bottom"
		}
		if isNotOnLeft && isValidContentType(SnakeWorldCellContents, world[y][x-1]) {
			classes += " left"
		}
		if isNotOnRight && isValidContentType(SnakeWorldCellContents, world[y][x+1]) {
			classes += " right"
		}
	} else if cell == WorldCellWall {
		classes += " wall"
		if !isNotOnTop || (isNotOnTop && world[y-1][x] == WorldCellWall) {
			classes += " top"
		}
		if !isNotOnBottom || (isNotOnBottom && world[y+1][x] == WorldCellWall) {
			classes += " bottom"
		}
		if !isNotOnLeft || (isNotOnLeft && world[y][x-1] == WorldCellWall) {
			classes += " left"
		}
		if !isNotOnRight || (isNotOnRight && world[y][x+1] == WorldCellWall) {
			classes += " right"
		}
	} else if cell == WorldCellFood {
		classes += " food"
	}
	return classes
}

func getCharByCellContent(content WorldCellContent) string {
	switch content {
	case WorldCellSnakeMovingUp:
		return "^"
	case WorldCellSnakeMovingDown:
		return "v"
	case WorldCellSnakeMovingRight:
		return ">"
	case WorldCellSnakeMovingLeft:
		return "<"
	case WorldCellWall:
		return "X"
	case WorldCellFood:
		return "0"
	default:
		return " "
	}
}

func getCellContentByChar(char string) WorldCellContent {
	switch char {
	case "^":
		return WorldCellSnakeMovingUp
	case "v":
		return WorldCellSnakeMovingDown
	case ">":
		return WorldCellSnakeMovingRight
	case "<":
		return WorldCellSnakeMovingLeft
	case "X":
		return WorldCellWall
	case "0":
		return WorldCellFood
	default:
		return WorldCellEmpty
	}
}

func printWorld() {
	printWorldString()
	updateHtmlWorld()
}

func tick() bool {
	fmt.Println("ticking...")
	moveSnake()
	printWorld()
	if !isSnakeAlive {
		clearInterval(tickIntervalId)
		tickIntervalId = 0
	}
	return isSnakeAlive
}

func startGame() {
	if tickIntervalId == 0 {
		fmt.Println("starting game...")
		isSnakeAlive = true
		currentScore = 0
		loadStage(0)
		printWorld()
		intervalId, err := setInterval("tickGame", 500)
		if err != nil {
			fmt.Printf("Failed to set interval! %s\n", err)
		}
		tickIntervalId = intervalId
	}
}
