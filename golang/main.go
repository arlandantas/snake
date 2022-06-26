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
	jsDoc, err := getDocument()
	if err == nil {
		jsDoc.Set("onkeypress", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			keyCode := args[0].Get("charCode").Int()
			if keyCode == 97 || keyCode == 119 || keyCode == 115 || keyCode == 100 {
				switch keyCode {
				case 97:
					setSnakeHeadDirection(WorldCellSnakeMovingLeft)
				case 119:
					setSnakeHeadDirection(WorldCellSnakeMovingUp)
				case 115:
					setSnakeHeadDirection(WorldCellSnakeMovingDown)
				case 100:
					setSnakeHeadDirection(WorldCellSnakeMovingRight)
				}
			}
			return true
		}))
	} else {
		fmt.Printf("Failed to get document: %s\n", err)
	}
}

func main() {
	exportJsFunctions()
	initilizeHandlers()
	loadStage(currentStageIndex)
	<-make(chan bool)
}
