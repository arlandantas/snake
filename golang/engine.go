package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WorldCellContent int

const (
	WorldCellEmpty WorldCellContent = iota
	WorldCellSnakeMovingUp
	WorldCellSnakeMovingRight
	WorldCellSnakeMovingDown
	WorldCellSnakeMovingLeft
	WorldCellWall
	WorldCellFood
)

type Position struct {
	x int
	y int
}

var world [][]WorldCellContent
var worldW = 0
var worldH = 0
var currentSpeed = 0
var currentSnakeHeadDirection = WorldCellEmpty
var currentSnakeHead = Position{y: 0, x: 0}
var currentSnakeTail = Position{y: 0, x: 0}
var isSnakeAlive = false
var currentScore = 0
var timeoutId = 0
var currentStageIndex = 0
var currentStage Stage
var randomSource = rand.New(rand.NewSource(time.Now().Unix()))

func loadStage(stageIndex int) {
	currentStage = stages[stageIndex]
	world = currentStage.initialWorld
	worldH = len(world)
	worldW = len(world[0])
	currentSnakeTail = currentStage.initialSnakeTail
	currentSnakeHead = currentStage.initialSnakeHead
	currentSnakeHeadDirection = world[currentSnakeHead.y][currentSnakeHead.x]
	currentSpeed = currentStage.initialSpeed
	renderInitialWorld()
}

func setSnakeHeadDirection(direction WorldCellContent) {
	if !isSnakeAlive {
		return
	}
	currentDirection := world[currentSnakeHead.y][currentSnakeHead.x]
	moved := false
	if direction == currentDirection {
		clearTickTimeout()
		tick()
		printWorld()
	} else if direction == WorldCellSnakeMovingUp && currentSnakeHeadDirection != WorldCellSnakeMovingDown {
		currentDirection = WorldCellSnakeMovingUp
		moved = true
	} else if direction == WorldCellSnakeMovingRight && currentSnakeHeadDirection != WorldCellSnakeMovingLeft {
		currentDirection = WorldCellSnakeMovingRight
		moved = true
	} else if direction == WorldCellSnakeMovingDown && currentSnakeHeadDirection != WorldCellSnakeMovingUp {
		currentDirection = WorldCellSnakeMovingDown
		moved = true
	} else if direction == WorldCellSnakeMovingLeft && currentSnakeHeadDirection != WorldCellSnakeMovingRight {
		currentDirection = WorldCellSnakeMovingLeft
		moved = true
	}
	if moved {
		world[currentSnakeHead.y][currentSnakeHead.x] = currentDirection
		printWorld()
	}
}

func moveSnake() {
	if !isSnakeAlive {
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
		if currentSpeed > 50 {
			currentSpeed -= 10
		}
		currentScore++
	}
	currentSnakeHeadDirection = previousHeadDirection
	world[currentSnakeHead.y][currentSnakeHead.x] = previousHeadDirection
	printWorld()
}

func clearTickTimeout() {
	clearTimeout(timeoutId)
	timeoutId = 0
}

func tick() {
	moveSnake()
	if isSnakeAlive {
		generatedTimeoutId, err := setTimeout("tickGame", currentSpeed)
		if err != nil {
			fmt.Printf("Failed to set interval! %s\n", err)
		}
		timeoutId = generatedTimeoutId
	}
}

func createFood() {
	y := randomSource.Intn(worldH)
	x := randomSource.Intn(worldW)
	if world[y][x] == WorldCellEmpty {
		world[y][x] = WorldCellFood
		return
	}
	createFood()
}

func startGame() {
	if !isSnakeAlive {
		fmt.Println("starting game...")
		isSnakeAlive = true
		currentScore = 0
		loadStage(currentStageIndex)
		randomSource = rand.New(rand.NewSource(time.Now().Unix()))
		createFood()
		printWorld()
		tick()
		startBtn, err := getElementById("start_game")
		if err != nil {
			fmt.Printf("Failed to get startGame button %s", err)
		} else {
			startBtn.Get("style").Set("display", "none")
		}
	}
}
