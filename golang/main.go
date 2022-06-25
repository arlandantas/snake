package main

import (
	"fmt"
	"syscall/js"
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

func loadStage() {
	worldHtml := ""
	for i := 0; i < WORLD_SIZE; i++ {
		worldHtml += "<div class=\"row\">"
		for j := 0; j < WORLD_SIZE; j++ {
			world[i][j] = WorldCellEmpty
			worldHtml += "<div class=\"cell\"></div>"
		}
		worldHtml += "</div>"
		div, err := getElementById("world")
		if err != "" {
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
		world[currentSnakeHead.y][currentSnakeHead.x] == WorldCellWall {
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
			switch {
			case world[i][j] == WorldCellSnakeMovingUp:
				content += "^"
			case world[i][j] == WorldCellSnakeMovingDown:
				content += "v"
			case world[i][j] == WorldCellSnakeMovingRight:
				content += ">"
			case world[i][j] == WorldCellSnakeMovingLeft:
				content += "<"
			case world[i][j] == WorldCellWall:
				content += "X"
			case world[i][j] == WorldCellFood:
				content += "0"
			default:
				content += " "
			}
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

func updateHtmlWorld() {
	div, err := getElementById("worldDiv")
	if err != "" {
		fmt.Printf("Failed to get worldDiv: %s\n", err)
	} else {
		div.Set("innerHTML", "Atualizado!!")
	}
}

func printWorld() {
	printWorldString()
}

func startGame() {
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

func getElementById(elementId string) (js.Value, string) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return js.Null(), "Document is invalid"
	}
	jsonElement := jsDoc.Call("getElementById", elementId)
	if !jsonElement.Truthy() {
		return js.Null(), "Element is invalid"
	}
	return jsonElement, ""
}

func exportJsFunctions() {
	js.Global().Set("startGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		startGame()
		return true
	}))
	// js.Global().Set("getWorldString", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	return getWorldString()
	// }))
}

func main() {
	loadStage()
	exportJsFunctions()
	printWorld()
	<-make(chan bool)
}
