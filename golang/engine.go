package main

import (
	"fmt"
	"math/rand"
	"time"
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
var currentStageIndex = 0
var currentStage Stage
var randomSource = rand.New(rand.NewSource(time.Now().Unix()))

func loadStage(stageIndex int) {
	currentStage = stages[stageIndex]
	world = currentStage.initialWorld
	worldH = len(world)
	worldW = len(world[0])
	currentSnakeHead = currentStage.initialSnakeHead
	currentSnakeTail = currentStage.initialSnakeTail
	renderInitialWorld()
}

func setSnakeHeadDirection(direction SnakeMovementDirection) {
	currentDirection := world[currentSnakeHead.y][currentSnakeHead.x]
	moved := false
	if direction == SnakeMoveUp && currentDirection != WorldCellSnakeMovingDown {
		currentDirection = WorldCellSnakeMovingUp
		moved = true
	} else if direction == SnakeMoveRight && currentDirection != WorldCellSnakeMovingLeft {
		currentDirection = WorldCellSnakeMovingRight
		moved = true
	} else if direction == SnakeMoveDown && currentDirection != WorldCellSnakeMovingUp {
		currentDirection = WorldCellSnakeMovingDown
		moved = true
	} else if direction == SnakeMoveLeft && currentDirection != WorldCellSnakeMovingRight {
		currentDirection = WorldCellSnakeMovingLeft
		moved = true
	}
	if moved {
		world[currentSnakeHead.y][currentSnakeHead.x] = currentDirection
		moveSnake(true)
		printWorld()
	}
}

func moveSnake(userInput bool) {
	if tickIntervalId == 0 {
		return
	}
	previousHeadDirection := world[currentSnakeHead.y][currentSnakeHead.x]
	newY, newX := currentSnakeHead.y, currentSnakeHead.x
	switch previousHeadDirection {
	case WorldCellSnakeMovingRight:
		newX += 1
		if !currentStage.boundary && newX == worldW {
			newX = 0
		}
	case WorldCellSnakeMovingLeft:
		newX -= 1
		if !currentStage.boundary && newX == -1 {
			newX = worldW - 1
		}
	case WorldCellSnakeMovingUp:
		newY -= 1
		if !currentStage.boundary && newY == -1 {
			newY = worldH - 1
		}
	case WorldCellSnakeMovingDown:
		newY += 1
		if !currentStage.boundary && newY == worldH {
			newY = 0
		}
	}
	if newX < 0 || newX >= worldW ||
		newY < 0 || newY >= worldH ||
		(world[newY][newX] != WorldCellEmpty &&
			world[newY][newX] != WorldCellFood &&
			(currentSnakeTail.x != newX || currentSnakeTail.y != newY)) {
		isSnakeAlive = false
		return
	}
	currentSnakeHead.y, currentSnakeHead.x = newY, newX
	if world[currentSnakeHead.y][currentSnakeHead.x] != WorldCellFood {
		previousTailDirection := world[currentSnakeTail.y][currentSnakeTail.x]
		world[currentSnakeTail.y][currentSnakeTail.x] = WorldCellEmpty
		switch previousTailDirection {
		case WorldCellSnakeMovingRight:
			currentSnakeTail.x += 1
			if !currentStage.boundary && currentSnakeTail.x == worldW {
				currentSnakeTail.x = 0
			}
		case WorldCellSnakeMovingLeft:
			currentSnakeTail.x -= 1
			if !currentStage.boundary && currentSnakeTail.x == -1 {
				currentSnakeTail.x = worldW - 1
			}
		case WorldCellSnakeMovingUp:
			currentSnakeTail.y -= 1
			if !currentStage.boundary && currentSnakeTail.y == -1 {
				currentSnakeTail.y = worldH - 1
			}
		case WorldCellSnakeMovingDown:
			currentSnakeTail.y += 1
			if !currentStage.boundary && currentSnakeTail.y == worldH {
				currentSnakeTail.y = 0
			}
		}
	} else {
		createFood()
		currentScore++
	}
	world[currentSnakeHead.y][currentSnakeHead.x] = previousHeadDirection
	skipNextTickMove = userInput
}

func tick() bool {
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
	randomSource = rand.New(rand.NewSource(time.Now().Unix()))
	y := randomSource.Intn(worldH)
	x := randomSource.Intn(worldW)
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
		loadStage(currentStageIndex)
		createFood()
		printWorld()
		intervalId, err := setInterval("tickGame", 500)
		if err != nil {
			fmt.Printf("Failed to set interval! %s\n", err)
		}
		tickIntervalId = intervalId
	}
}
