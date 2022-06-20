package main

import "fmt"

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

func start() {
	world[1][2] = WorldCellSnakeMovingRight
	world[1][3] = WorldCellSnakeMovingRight
	world[1][4] = WorldCellSnakeMovingRight
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
	world[currentSnakeHead.y][currentSnakeHead.x] = previousHeadDirection
	if currentSnakeHead.x < 0 || currentSnakeHead.x >= WORLD_SIZE ||
		currentSnakeHead.y < 0 || currentSnakeHead.y >= WORLD_SIZE ||
		world[currentSnakeHead.y][currentSnakeHead.x] == WorldCellWall {
		fmt.Println("A cobra morreu!")
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
	}
}

func printWorld() {
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
}
