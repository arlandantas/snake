package main

import (
	"fmt"
	"time"
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

func start() {
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
		world[currentSnakeHead.y][currentSnakeHead.x] == WorldCellWall {
		fmt.Println("A cobra morreu!")
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
		fmt.Println("A cobra se alimentou!")
		currentScore++
	}
	world[currentSnakeHead.y][currentSnakeHead.x] = previousHeadDirection
}

func printWorld() {
	fmt.Printf("Score: %d\n", currentScore)
	fmt.Print("|")
	for i := 0; i < WORLD_SIZE; i++ {
		fmt.Print("-")
	}
	fmt.Println("|")
	for i := 0; i < len(world); i++ {
		fmt.Print("|")
		for j := 0; j < len(world[i]); j++ {
			switch {
			case world[i][j] == WorldCellSnakeMovingUp:
				fmt.Print("^")
			case world[i][j] == WorldCellSnakeMovingDown:
				fmt.Print("v")
			case world[i][j] == WorldCellSnakeMovingRight:
				fmt.Print(">")
			case world[i][j] == WorldCellSnakeMovingLeft:
				fmt.Print("<")
			case world[i][j] == WorldCellWall:
				fmt.Print("X")
			case world[i][j] == WorldCellFood:
				fmt.Print("0")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
	fmt.Print("|")
	for i := 0; i < WORLD_SIZE; i++ {
		fmt.Print("-")
	}
	fmt.Println("|")
}

func main() {
	start()
	printWorld()
	directions := []SnakeMovementDirection{
		SnakeMoveRight,
		SnakeMoveDown,
		SnakeMoveLeft,
		SnakeMoveUp,
	}
	for j := 0; j < 4; j++ {
		setSnakeHeadDirection(directions[j])
		for i := 0; i < 6; i++ {
			time.Sleep(250 * time.Millisecond)
			moveSnake()
			printWorld()
			if !isSnakeAlive {
				break
			}
		}
		if !isSnakeAlive {
			break
		}
	}
	fmt.Println("Game Over")
}
