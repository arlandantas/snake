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
	for i := range world {
		worldHtml += "\t<div class=\"row\">\n"
		for j := range world[i] {
			worldHtml += fmt.Sprintf("\t\t<div class=\"cell\" id=\"cell%d%d\"></div>\n", i, j)
		}
		worldHtml += "</div>"
	}
	div, err := getElementById("world")
	if err != nil {
		fmt.Printf("Failed to get world div: %s\n", err)
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

func getCellClasses(y, x int) string {
	classes := "cell"
	cell := world[y][x]
	cellPosition := Position{x, y}
	isNotOnTop := y > 0
	isNotOnBottom := y < worldH-1
	isNotOnLeft := x > 0
	isNotOnRight := x < worldW-1
	if currentSnakeHead == cellPosition {
		classes += " snake  head"
		switch cell {
		case WorldCellSnakeMovingUp:
			classes += " bottom"
		case WorldCellSnakeMovingDown:
			classes += " top"
		case WorldCellSnakeMovingRight:
			classes += " left"
		case WorldCellSnakeMovingLeft:
			classes += " right"
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
		if isNotOnTop && (world[y-1][x] == WorldCellSnakeMovingDown || cell == WorldCellSnakeMovingUp) {
			classes += " top"
		}
		if isNotOnBottom && (world[y+1][x] == WorldCellSnakeMovingUp || cell == WorldCellSnakeMovingDown) {
			classes += " bottom"
		}
		if isNotOnLeft && (world[y][x-1] == WorldCellSnakeMovingRight || cell == WorldCellSnakeMovingLeft) {
			classes += " left"
		}
		if isNotOnRight && (world[y][x+1] == WorldCellSnakeMovingLeft || cell == WorldCellSnakeMovingRight) {
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
	// printWorldString()
	updateHtmlWorld()
}
