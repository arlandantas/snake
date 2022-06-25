package main

import (
	"fmt"
)

type WorldCellContent int
type SnakeMovementDirection int

const (
	WorldCellEmpty WorldCellContent = iota
	WorldCellSnakeHead
	WorldCellSnakeMovingUp
	WorldCellSnakeMovingRight
	WorldCellSnakeMovingDown
	WorldCellSnakeMovingLeft
	WorldCellWall
	WorldCellFood
)

const (
	SnakeMoveUp SnakeMovementDirection = iota
	SnakeMoveRight
	SnakeMoveDown
	SnakeMoveLeft
)

const WORLD_SIZE = 10

var world [WORLD_SIZE][WORLD_SIZE]WorldCellContent

type Position struct {
	x int
	y int
}

var currentSnakeDirection = SnakeMoveRight
var currentSnakeHead = Position{y: 1, x: 4}
var currentSnakeTail = Position{y: 1, x: 2}
var isSnakeAlive = true
var currentScore = 0
var tickIntervalId = 0

func loadStage() {
	worldHtml := ""
	for i := 0; i < WORLD_SIZE; i++ {
		worldHtml += "\t<div class=\"row\">\n"
		for j := 0; j < WORLD_SIZE; j++ {
			world[i][j] = WorldCellEmpty
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
	currentSnakeHead = Position{y: 1, x: 2}
	currentSnakeTail = Position{y: 1, x: 0}
	world[1][0] = WorldCellSnakeMovingRight
	world[1][1] = WorldCellSnakeMovingRight
	world[1][2] = WorldCellSnakeMovingRight
	for i := 0; i < 10; i++ {
		world[6][i] = WorldCellFood
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
	if currentSnakeHead.x < 0 || currentSnakeHead.x >= WORLD_SIZE ||
		currentSnakeHead.y < 0 || currentSnakeHead.y >= WORLD_SIZE ||
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
	for i := 0; i < WORLD_SIZE; i++ {
		content += "-"
	}
	content += "|\n"
	for i := 0; i < len(world); i++ {
		content += "|"
		for j := 0; j < len(world[i]); j++ {
			content += getHtmlCellContent(world[i][j])
		}
		content += "|\n"
	}
	content += "|"
	for i := 0; i < WORLD_SIZE; i++ {
		content += "-"
	}
	content += "|\n"
	fmt.Println(content)
}

func getCellId(y, x int) string {
	return fmt.Sprintf("cell%d%d", y, x)
}

func updateHtmlWorld() {
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			cell, err := getElementById(getCellId(y, x))
			if err != nil {
				fmt.Printf("Failed to get worldDiv: %s\n", err)
				break
			}
			cell.Set("innerHTML", getHtmlCellContent(world[y][x]))
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

func getHtmlCellContent(content WorldCellContent) string {
	switch {
	case content == WorldCellSnakeMovingUp:
		return "^"
	case content == WorldCellSnakeMovingDown:
		return "v"
	case content == WorldCellSnakeMovingRight:
		return ">"
	case content == WorldCellSnakeMovingLeft:
		return "<"
	case content == WorldCellWall:
		return "X"
	case content == WorldCellFood:
		return "0"
	default:
		return " "
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
		loadStage()
		printWorld()
		intervalId, err := setInterval("tickGame", 500)
		if err != nil {
			fmt.Printf("Failed to set interval! %s\n", err)
		}
		tickIntervalId = intervalId
	}
}
