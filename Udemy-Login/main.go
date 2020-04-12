package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// x, y := robotgo.GetMousePos()
	// robotgo.ScrollMouse(10, "up")

	robotgo.MoveClick(19, 150, "left", false)
	robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	time.Sleep(time.Second * 3)
	robotgo.MoveClick(296, 77, "left", false)
	robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	time.Sleep(time.Second * 3)
	robotgo.MoveClick(354, 110, "left", false)
	time.Sleep(time.Second * 3)
	robotgo.TypeStr("Udemy")
	time.Sleep(time.Second * 3)
	robotgo.KeyTap("enter")
	robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	time.Sleep(time.Second * 10)
	robotgo.MoveClick(1194, 159, "left", false)
	robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	time.Sleep(time.Second * 15)
	robotgo.MoveClick(730, 410, "left", false)
	robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	time.Sleep(time.Second * 15)
	robotgo.MoveClick(582, 495, "left", false)
	robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	time.Sleep(time.Second * 15)
	robotgo.MoveClick(503, 376, "left", false)
	time.Sleep(time.Second * 10)
	var password string
	if password != "" {
		robotgo.TypeStr(password)
		time.Sleep(time.Second * 5)
		robotgo.KeyTap("enter")
		time.Sleep(time.Second * 5)
		robotgo.MoveMouseSmooth(600, 400, 1.0, 20.0)
	}
	fmt.Println("Error please provide the password")

}

