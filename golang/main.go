package main

import "fmt"

type WorldCellContent int

const (
	WorldCellEmpty WorldCellContent = iota
	WorldCellSnakeHead
	WorldCellSnakeMovingLeft
	WorldCellSnakeMovingUp
	WorldCellSnakeMovingRight
	WorldCellSnakeMovingDown
	WorldCellWall
	WorldCellFood
)

const WORLD_SIZE = 10

var world [WORLD_SIZE][WORLD_SIZE]WorldCellContent

func update() {
	world[1][1] = WorldCellSnakeMovingUp
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
			if world[i][j] == WorldCellEmpty {
				fmt.Print(" ")
			} else {
				fmt.Print(world[i][j])
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
	update()
	printWorld()
}
