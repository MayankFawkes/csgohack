package device

import "golang.org/x/sys/windows"

var (
	user32 = windows.NewLazySystemDLL("user32.dll")
	// procKeyBd = user32.NewProc("keybd_event")
	mouse_event = user32.NewProc("mouse_event")
)

// see http://msdn.microsoft.com/en-us/library/ms646260(VS.85).aspx for details
const (
	MOUSEEVENTF_LEFTDOWN   = 0x0002 // left button down
	MOUSEEVENTF_LEFTUP     = 0x0004 // left button up
	MOUSEEVENTF_RIGHTDOWN  = 0x0008 // right button down
	MOUSEEVENTF_RIGHTUP    = 0x0010 // right button up
	MOUSEEVENTF_MIDDLEDOWN = 0x0020 // middle button down
	MOUSEEVENTF_MIDDLEUP   = 0x0040 // middle button up
	MOUSEEVENTF_MOVE       = 0x0001 // Movement occurred.
)
