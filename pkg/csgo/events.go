package csgo

import (
	"fmt"
	"os"

	"github.com/mayankfawkes/csgohack/pkg/events"
)

func Keyboard() {
	c := make(chan events.KeyboardEvent)
	go events.StartKeyboard(c)

	for {
		m := <-c
		if m.Message == events.WM_KEYUP {
			if events.VK_F5 <= m.KBDLLHOOKSTRUCT.VkCode && m.KBDLLHOOKSTRUCT.VkCode <= events.VK_F9 {
				// fmt.Printf("Key pressed: %d\n", m.KBDLLHOOKSTRUCT.VkCode)

				gun := GUN_BY_KEY[m.KBDLLHOOKSTRUCT.VkCode]
				// fmt.Printf("Setting gun: %s\n", gun)

				status := State.UpdateGun(gun)

				if gun == "" {
					fmt.Printf("Gun status: %t\n", status)
				} else {
					fmt.Printf("Gun %s status: %t\n", gun, status)
				}
			}

			if events.VK_F10 == m.KBDLLHOOKSTRUCT.VkCode {
				fmt.Printf("Exiting...\n")
				os.Exit(1)
			}

			if events.VK_ESCAPE == m.KBDLLHOOKSTRUCT.VkCode {
				status := State.UpdateGun("")
				fmt.Printf("Gun status: %t\n", status)
			}
		}
	}
}

func Mouse() {
	c := make(chan events.MouseEvent)
	go events.StartMouse(c)

	for {
		m := <-c

		if m.MSLLHOOKSTRUCT.DwExtraInfo == uintptr(1) {
			continue
		}

		if m.Message == events.WM_LBUTTONDOWN {
			State.LeftKeyDown()
		}
		if m.Message == events.WM_LBUTTONUP {
			State.LeftKeyUp()
		}
	}
}
