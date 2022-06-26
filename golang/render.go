package main

import (
	"fmt"
)

var SnakeWorldCellContents = []WorldCellContent{
	WorldCellSnakeMovingUp,
	WorldCellSnakeMovingRight,
	WorldCellSnakeMovingDown,
	WorldCellSnakeMovingLeft,
}

func renderInitialWorld() {
	worldHtml := ""
	if currentStage.boundary {
		worldHtml += "<div class=\"row boundary\">"
		for j := 0; j < worldW+2; j++ {
			worldHtml += "<div class=\"cell\"></div>"
		}
		worldHtml += "</div>"
	}
	for i := range world {
		worldHtml += "<div class=\"row\">"
		if currentStage.boundary {
			worldHtml += "<div class=\"cell boundary\"></div>"
		}
		for j := range world[i] {
			worldHtml += fmt.Sprintf("<div class=\"cell\" id=\"%s\"></div>", getCellId(i, j))
		}
		if currentStage.boundary {
			worldHtml += "<div class=\"cell boundary\"></div>"
		}
		worldHtml += "</div>"
	}
	if currentStage.boundary {
		worldHtml += "<div class=\"row boundary\">"
		for j := 0; j < worldW+2; j++ {
			worldHtml += "<div class=\"cell\"></div>"
		}
		worldHtml += "</div>"
	}
	div, err := getElementById("world")
	if err != nil {
		fmt.Printf("Failed to get world div: %s", err)
	} else {
		div.Set("innerHTML", worldHtml)
	}
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
	return fmt.Sprintf("cell-%d-%d", y, x)
}

func printWorld(ignoreGameOver ...bool) {
	for y := range world {
		for x := range world[y] {
			cell, err := getElementById(getCellId(y, x))
			if err != nil {
				fmt.Printf("Failed to get cellDiv: %s\n", err)
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
	scoreText := fmt.Sprintf("Score: %d", currentScore)
	if isSnakeAlive || (len(ignoreGameOver) > 0 && ignoreGameOver[0]) {
		bMessage.Set("innerHTML", scoreText)
	} else {
		bMessage.Set("innerHTML", fmt.Sprintf("GAME OVER! %s", scoreText))
		setElementDisplay("start_game", "")
		setElementDisplay("stop_game", "none")
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

func getCellClasses(y, x int) string {
	classes := "cell"
	cell := world[y][x]
	cellPosition := Position{x, y}
	isOnTop := y == 0
	isOnBottom := y == worldH-1
	isOnLeft := x == 0
	isOnRight := x == worldW-1
	if currentSnakeHead == cellPosition {
		classes += " snake  head"
		switch cell {
		case WorldCellSnakeMovingUp:
			classes += " bottom"
			if (!isOnTop && world[y-1][x] == WorldCellFood) ||
				(isOnTop && !currentStage.boundary && world[worldH-1][x] == WorldCellFood) {
				classes += " tongue"
			}
		case WorldCellSnakeMovingDown:
			classes += " top"
			if (!isOnBottom && world[y+1][x] == WorldCellFood) ||
				(isOnBottom && !currentStage.boundary && world[0][x] == WorldCellFood) {
				classes += " tongue"
			}
		case WorldCellSnakeMovingRight:
			classes += " left"
			if (!isOnRight && world[y][x+1] == WorldCellFood) ||
				(isOnRight && !currentStage.boundary && world[y][0] == WorldCellFood) {
				classes += " tongue"
			}
		case WorldCellSnakeMovingLeft:
			classes += " right"
			if (!isOnLeft && world[y][x-1] == WorldCellFood) ||
				(isOnLeft && !currentStage.boundary && world[y][worldW-1] == WorldCellFood) {
				classes += " tongue"
			}
		}
	} else if currentSnakeTail == cellPosition {
		classes += " snake"
		switch cell {
		case WorldCellSnakeMovingUp:
			classes += " top"
		case WorldCellSnakeMovingDown:
			classes += " bottom"
		case WorldCellSnakeMovingRight:
			classes += " right"
		case WorldCellSnakeMovingLeft:
			classes += " left"
		}
	} else if isValidContentType(SnakeWorldCellContents, cell) {
		classes += " snake"
		if cell == WorldCellSnakeMovingUp || (!isOnTop && world[y-1][x] == WorldCellSnakeMovingDown) ||
			(isOnTop && !currentStage.boundary && world[worldH-1][x] == WorldCellSnakeMovingDown) {
			classes += " top"
		}
		if cell == WorldCellSnakeMovingDown || (!isOnBottom && world[y+1][x] == WorldCellSnakeMovingUp) ||
			(isOnBottom && !currentStage.boundary && world[0][x] == WorldCellSnakeMovingUp) {
			classes += " bottom"
		}
		if cell == WorldCellSnakeMovingLeft || (!isOnLeft && world[y][x-1] == WorldCellSnakeMovingRight) ||
			(isOnLeft && !currentStage.boundary && world[y][worldW-1] == WorldCellSnakeMovingRight) {
			classes += " left"
		}
		if cell == WorldCellSnakeMovingRight || (!isOnRight && world[y][x+1] == WorldCellSnakeMovingLeft) ||
			(isOnRight && !currentStage.boundary && world[y][0] == WorldCellSnakeMovingLeft) {
			classes += " right"
		}
	} else if cell == WorldCellWall {
		classes += " wall"
		if isOnTop || (!isOnTop && world[y-1][x] == WorldCellWall) {
			classes += " top"
		}
		if isOnBottom || (!isOnBottom && world[y+1][x] == WorldCellWall) {
			classes += " bottom"
		}
		if isOnLeft || (!isOnLeft && world[y][x-1] == WorldCellWall) {
			classes += " left"
		}
		if isOnRight || (!isOnRight && world[y][x+1] == WorldCellWall) {
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
	case "x":
		return WorldCellWall
	case "0":
		return WorldCellFood
	default:
		return WorldCellEmpty
	}
}
