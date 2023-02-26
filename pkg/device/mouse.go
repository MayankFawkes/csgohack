package device

import (
	"time"
)

func SendRawMouseEvent(dwFlags, dx, dy, dwData, dwExtraInfo uintptr) {
	mouse_event.Call(dwFlags, dx, dy, dwData, dwExtraInfo)
}

type ProcMouse struct{}

func Mouse() ProcMouse {
	return ProcMouse{}
}

func SendMouseEvent() {
	mouse_event.Call(MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 0)

	time.Sleep(1 * time.Second)

	mouse_event.Call(MOUSEEVENTF_LEFTUP, 0, 0, 0, 0)
}

func (p ProcMouse) Move(x, y float32) {
	SendRawMouseEvent(MOUSEEVENTF_MOVE, uintptr(x), uintptr(y), 0, 0)
}

func (p ProcMouse) LeftKeyDown() {
	SendRawMouseEvent(MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 1)
}

func (p ProcMouse) LeftKeyUp() {
	SendRawMouseEvent(MOUSEEVENTF_LEFTUP, 0, 0, 0, 1)
}

func (p ProcMouse) RightKeyDown() {
	SendRawMouseEvent(MOUSEEVENTF_RIGHTDOWN, 0, 0, 0, 0)
}

func (p ProcMouse) RightKeyUp() {
	SendRawMouseEvent(MOUSEEVENTF_RIGHTUP, 0, 0, 0, 0)
}
