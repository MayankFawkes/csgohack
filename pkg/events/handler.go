package events

import (
	"unsafe"
)

type KeyboardEvent struct {
	Message Message
	KBDLLHOOKSTRUCT
}

type MouseEvent struct {
	Message Message
	MSLLHOOKSTRUCT
}

func DefaultMouseHookHandler(c chan<- MouseEvent) HOOKPROC {
	return func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
		if lparam != 0 {
			c <- MouseEvent{
				Message:        Message(wparam),
				MSLLHOOKSTRUCT: *(*MSLLHOOKSTRUCT)(unsafe.Pointer(lparam)),
			}
		}

		return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
	}
}

func DefaultKeyboardHookHandler(c chan<- KeyboardEvent) HOOKPROC {
	return func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
		if lparam != 0 {
			c <- KeyboardEvent{
				Message:         Message(wparam),
				KBDLLHOOKSTRUCT: *(*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam)),
			}
		}

		return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
	}
}
