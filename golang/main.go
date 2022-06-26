package main

import (
	"fmt"
	"syscall/js"
)

func exportJsFunctions() {
	js.Global().Set("startGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			fmt.Println("You must choose a stage to play!")
			return false
		}
		startGame(args[0].Int())
		return true
	}))
	js.Global().Set("giveup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		giveup()
		return true
	}))
	js.Global().Set("pauseGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		pauseGame()
		return true
	}))
	js.Global().Set("resumeGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resumeGame()
		return true
	}))
	js.Global().Set("tickGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		tick()
		return true
	}))
	js.Global().Set("clickUp", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(WorldCellSnakeMovingUp)
		return true
	}))
	js.Global().Set("clickDown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(WorldCellSnakeMovingDown)
		return true
	}))
	js.Global().Set("clickLeft", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(WorldCellSnakeMovingLeft)
		return true
	}))
	js.Global().Set("clickRight", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(WorldCellSnakeMovingRight)
		return true
	}))
}

func initilizeHandlers() {
	js.Global().Set("onkeyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyCode := args[0].Get("keyCode").Int()
		if keyCode == 37 || keyCode == 38 || keyCode == 39 || keyCode == 40 ||
			keyCode == 65 || keyCode == 87 || keyCode == 83 || keyCode == 68 {
			switch keyCode {
			case 37:
				setSnakeHeadDirection(WorldCellSnakeMovingLeft)
			case 65:
				setSnakeHeadDirection(WorldCellSnakeMovingLeft)
			case 38:
				setSnakeHeadDirection(WorldCellSnakeMovingUp)
			case 87:
				setSnakeHeadDirection(WorldCellSnakeMovingUp)
			case 40:
				setSnakeHeadDirection(WorldCellSnakeMovingDown)
			case 83:
				setSnakeHeadDirection(WorldCellSnakeMovingDown)
			case 39:
				setSnakeHeadDirection(WorldCellSnakeMovingRight)
			case 68:
				setSnakeHeadDirection(WorldCellSnakeMovingRight)
			}
		}
		return true
	}))
}

func main() {
	exportJsFunctions()
	initilizeHandlers()
	loadStage(0)
	<-make(chan bool)
}
