package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	str, err := robotgo.TypeStr("Hello World")
	fmt.Println(str)
	// robotgo.TypeStr("It's Awesome to play with Keyboard", 1.0)
	// robotgo.TypeStr("Hi Galaxy. Awesome.")
	// robotgo.Sleep(1)
	// robotgo.KeyTap("enter")
	// robotgo.KeyTap("c", "ctrl", "command")
	// arr := []string{"ctrl", "command"}
	// robotgo.KeyTap("c", arr)
	// robotgo.WriteAll("Test")
	// text, err := robotgo.ReadAll()
	// if err == nil {
	// 	fmt.Println(text)
	// }
}
