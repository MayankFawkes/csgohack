package csgo

import (
	"embed"
	"fmt"
	"time"

	"github.com/mayankfawkes/csgohack/pkg/device"
)

var (
	SENS        float32
	ZOONSENS    float32
	MODIFIER    float32
	GUN_PRE_SET map[string]Gun
	Sound       embed.FS
)

func StartRecoil(g Gun, mouse device.ProcMouse) {
	if len(g.Points) == 0 {
		return
	}

	mouse.LeftKeyDown()

	time.Sleep(g.Before)
	for _, p := range g.Points {
		if State.IsLeftKeyDown() {
			mouse.Move(p.x, p.y)
			time.Sleep(time.Duration(p.sleep) * time.Millisecond)
			// time.Sleep(84 * time.Millisecond)
		} else {
			mouse.LeftKeyUp()
			break
		}
	}

	mouse.LeftKeyUp()
}

func BotStart() {

	SENS = 2.52
	ZOONSENS = 1.0

	MODIFIER = 2.52 / SENS

	initGunPoints()

	btn_press := make(chan bool)
	mouse := device.Mouse()

	// Init State
	State = &Storage{
		Gun:      "",
		BtnPress: btn_press,
		Mouse:    MOUSE_STATE,
	}

	go Keyboard()
	go Mouse()

	for <-btn_press {
		if State.Gun == "" {
			continue
		}

		fmt.Println("Resoil start for ", State.Gun)
		go StartRecoil(GUN_PRE_SET[State.Gun], mouse)
	}

}
