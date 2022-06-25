package main

import (
	"fmt"
	"math/rand"
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

const (
	SnakeMoveUp SnakeMovementDirection = iota
	SnakeMoveRight
	SnakeMoveDown
	SnakeMoveLeft
)

var world [][]WorldCellContent
var worldW = 0
var worldH = 0

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
var skipNextTickMove = false

func loadStage(stageIndex int) {
	stage := stages[stageIndex]
	world = stage.initialWorld
	worldH = len(world)
	worldW = len(world[0])
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

func moveSnake(userInput bool) {
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
	if currentSnakeHead.x < 0 || currentSnakeHead.x >= worldW ||
		currentSnakeHead.y < 0 || currentSnakeHead.y >= worldH ||
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
		createFood()
		currentScore++
	}
	world[currentSnakeHead.y][currentSnakeHead.x] = previousHeadDirection
	skipNextTickMove = userInput
}

func tick() bool {
	fmt.Println("ticking...")
	if !skipNextTickMove {
		moveSnake(false)
	}
	skipNextTickMove = false
	printWorld()
	if !isSnakeAlive {
		clearInterval(tickIntervalId)
		tickIntervalId = 0
	}
	return isSnakeAlive
}

func createFood() {
	y := rand.Intn(worldH)
	x := rand.Intn(worldW)
	if world[y][x] == WorldCellEmpty {
		world[y][x] = WorldCellFood
		return
	}
	createFood()
}

func startGame() {
	if tickIntervalId == 0 {
		fmt.Println("starting game...")
		isSnakeAlive = true
		currentScore = 0
		loadStage(0)
		createFood()
		printWorld()
		intervalId, err := setInterval("tickGame", 500)
		if err != nil {
			fmt.Printf("Failed to set interval! %s\n", err)
		}
		tickIntervalId = intervalId
	}
}
