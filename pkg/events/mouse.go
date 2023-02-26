package events

func StartMouse(c chan<- MouseEvent) {
	// defer user32.Release()
	keyboardHook = SetWindowsHookEx(WH_MOUSE_LL, (HOOKPROC)(DefaultMouseHookHandler(c)), 0, 0)
	var msg MSG
	// for GetMessage(&msg, 0, 0, 0) != 0 {
	// }
	for GetMessage(&msg, 0, 0, 0) != 0 {
		TranslateMessage(&msg)
		DispatchMessage(&msg)
		// fmt.Println("key pressed:")

	}

	UnhookWindowsHookEx(keyboardHook)
	keyboardHook = 0
}
