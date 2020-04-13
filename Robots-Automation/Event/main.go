package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	ok := robotgo.AddEvent("q")
	if ok {
		fmt.Println("add events...")
	}
	keve := robotgo.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}
	mleft := robotgo.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}
