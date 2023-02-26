package main

import (
	"embed"

	"github.com/mayankfawkes/csgohack/pkg/csgo"
)

//go:embed sound/*
var sound embed.FS

func main() {
	// Mouse()

	// SendMouse := device.Mouse()

	// SendMouse.LeftKeyDown()
	// SendMouse.Move(10, 10)
	// SendMouse.Move(10, 10)
	// SendMouse.Move(10, 10)

	// time.Sleep(1 * time.Second)

	// SendMouse.LeftKeyUp()
	csgo.Sound = sound
	csgo.BotStart()

	// mouse := device.Mouse()

	// slp := 94 * time.Millisecond

	// time.Sleep(5000 * time.Millisecond)

	// mouse.LeftKeyDown()
	// time.Sleep(50 * time.Millisecond)

	// mouse.Move(-4, 7)
	// time.Sleep(slp)
	// mouse.Move(4, 19)
	// time.Sleep(slp)
	// mouse.Move(-3, 29)
	// time.Sleep(slp)
	// mouse.Move(-1, 31)
	// time.Sleep(slp)
	// mouse.Move(13, 31)
	// time.Sleep(slp)
	// mouse.Move(8, 28)
	// time.Sleep(slp)
	// mouse.Move(13, 21)
	// time.Sleep(slp)
	// mouse.Move(-17, 12)
	// time.Sleep(slp)
	// mouse.Move(-42, -3)
	// time.Sleep(slp)
	// mouse.Move(-21, 2)
	// time.Sleep(slp)
	// mouse.Move(12, 11)
	// time.Sleep(slp)
	// mouse.Move(-15, 7)
	// time.Sleep(slp)
	// mouse.Move(-26, -8)
	// time.Sleep(slp)
	// mouse.Move(-3, 4)
	// time.Sleep(slp)
	// mouse.Move(40, 1)
	// time.Sleep(slp)
	// mouse.Move(19, 7)
	// time.Sleep(slp)
	// mouse.Move(14, 10)
	// time.Sleep(slp)
	// mouse.Move(27, 0)
	// time.Sleep(slp)
	// mouse.Move(33, -10)
	// time.Sleep(slp)
	// mouse.Move(-21, -2)
	// time.Sleep(slp)
	// mouse.Move(7, 3)
	// time.Sleep(slp)
	// mouse.Move(-7, 9)
	// time.Sleep(slp)
	// mouse.Move(-8, 4)
	// time.Sleep(slp)
	// mouse.Move(19, -3)
	// time.Sleep(slp)
	// mouse.Move(5, 6)
	// time.Sleep(slp)
	// mouse.Move(-20, -1)
	// time.Sleep(slp)
	// mouse.Move(-33, -4)
	// time.Sleep(slp)
	// mouse.Move(-45, -21)
	// time.Sleep(slp)
	// mouse.Move(-14, 1)
	// time.Sleep(slp)

	// mouse.LeftKeyUp()

}
