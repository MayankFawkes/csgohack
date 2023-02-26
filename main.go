package main

import (
	"embed"

	"github.com/mayankfawkes/csgohack/pkg/csgo"
)

//go:embed sound/*
var sound embed.FS

func main() {

	csgo.Sound = sound
	csgo.BotStart()

}
