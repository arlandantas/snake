package main

import (
	"fmt"
	"syscall/js"
)

func exportJsFunctions() {
	js.Global().Set("startGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		startGame()
		return true
	}))
	js.Global().Set("tickGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		tick()
		return true
	}))
	js.Global().Set("clickUp", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(SnakeMoveUp)
		return true
	}))
	js.Global().Set("clickDown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(SnakeMoveDown)
		return true
	}))
	js.Global().Set("clickLeft", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(SnakeMoveLeft)
		return true
	}))
	js.Global().Set("clickRight", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setSnakeHeadDirection(SnakeMoveRight)
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
					setSnakeHeadDirection(SnakeMoveLeft)
				case 119:
					setSnakeHeadDirection(SnakeMoveUp)
				case 115:
					setSnakeHeadDirection(SnakeMoveDown)
				case 100:
					setSnakeHeadDirection(SnakeMoveRight)
				}
				moveSnake(true)
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
	loadStage(0)
	printWorld()
	<-make(chan bool)
}
