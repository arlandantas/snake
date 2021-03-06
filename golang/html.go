package main

import (
	"errors"
	"fmt"
	"syscall/js"
)

func getDocument() (js.Value, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return js.Null(), errors.New("Document is invalid")
	}
	return jsDoc, nil
}

func querySelector(selector string) (js.Value, error) {
	jsDoc, jsDocErr := getDocument()
	if jsDocErr != nil {
		return js.Null(), jsDocErr
	}
	querySelector := jsDoc.Call("querySelector", selector)
	if !querySelector.Truthy() {
		return js.Null(), errors.New("Element is invalid")
	}
	return querySelector, nil
}

func getElementById(elementId string) (js.Value, error) {
	return querySelector("#" + elementId)
}

func setElementDisplay(elementId, display string) bool {
	elmt, err := getElementById(elementId)
	if err != nil {
		fmt.Printf("Failed to get elment %s %s", elementId, err)
		return false
	}
	elmt.Get("style").Set("display", display)
	return true
}

func setTimeout(functionName string, interval int) (int, error) {
	timeoutId := js.Global().Call("setTimeout", js.Global().Get(functionName), interval)
	if !timeoutId.Truthy() {
		return 0, errors.New("Failed to set timeout")
	}
	return timeoutId.Int(), nil
}

func clearTimeout(timeoutId int) {
	js.Global().Call("clearTimeout", timeoutId)
}

func setInterval(functionName string, interval int) (int, error) {
	intervalId := js.Global().Call("setInterval", js.Global().Get(functionName), interval)
	if !intervalId.Truthy() {
		return 0, errors.New("Failed to set interval")
	}
	return intervalId.Int(), nil
}

func clearInterval(intervalId int) {
	js.Global().Call("clearInterval", intervalId)
}
