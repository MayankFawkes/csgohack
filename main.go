package main

import (
	"embed"
	"fmt"

	"github.com/mayankfawkes/csgohack/pkg/csgo"
)

//go:embed sound/*
var sound embed.FS

func main() {
	var inputSENS float32
	fmt.Printf("Enter Sensitivity: ")
	n, err := fmt.Scanf("%f\n", &inputSENS)
	if err != nil || n != 1 {
		fmt.Println("Error", n, err)
		return
	}

	fmt.Println("Successfully Started")

	csgo.SENS = inputSENS
	csgo.MODIFIER = 2.52 / inputSENS
	csgo.Sound = sound
	csgo.BotStart()
}
