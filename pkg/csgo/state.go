package csgo

import (
	"fmt"
	"sync"

	"github.com/mayankfawkes/csgohack/pkg/events"
)

var (
	MOUSE_STATE = map[events.Message]bool{
		events.WM_LBUTTONDOWN: false,
		events.WM_LBUTTONUP:   false,
	}

	GUN_BY_KEY = map[events.VKCode]string{
		events.VK_F5: "m4a4",
		events.VK_F6: "m4a1-s",
		events.VK_F7: "ak-47",
		events.VK_F8: "galil",
		events.VK_F9: "famas",
	}
)

var (
	MOUSE_LOCK    = sync.RWMutex{}
	KEYBOARD_LOCK = sync.RWMutex{}
)

type Storage struct {
	Gun       string
	BtnPress  chan bool
	Mouse     map[events.Message]bool
	MouseLock sync.RWMutex
}

var State *Storage

func (s *Storage) LeftKeyDown() {
	s.MouseLock.Lock()
	defer s.MouseLock.Unlock()

	s.Mouse[events.WM_LBUTTONDOWN] = true
	s.Mouse[events.WM_LBUTTONUP] = false

	s.BtnPress <- true
}

func (s *Storage) LeftKeyUp() {
	s.MouseLock.Lock()
	defer s.MouseLock.Unlock()

	s.Mouse[events.WM_LBUTTONUP] = true
	s.Mouse[events.WM_LBUTTONDOWN] = false
}

func (s *Storage) IsLeftKeyDown() bool {
	s.MouseLock.RLock()
	defer s.MouseLock.RUnlock()

	return s.Mouse[events.WM_LBUTTONDOWN]
}

func (s *Storage) UpdateGun(g string) bool {
	if len(g) == 0 {
		if len(s.Gun) > 0 {
			go Play(fmt.Sprintf("%s-deactivated", s.Gun))
			s.Gun = ""
			return false
		} else {
			s.Gun = ""
			return false
		}
	}
	if g == s.Gun {
		s.Gun = ""
		go Play(fmt.Sprintf("%s-deactivated", g))
		return false
	} else {
		s.Gun = g
		go Play(fmt.Sprintf("%s-activated", g))
		return true
	}

}

func (s *Storage) Close() {
	s.BtnPress <- false
}
