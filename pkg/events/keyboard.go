package events

func StartKeyboard(c chan<- KeyboardEvent) {
	// defer user32.Release()
	keyboardHook = SetWindowsHookEx(WH_KEYBOARD_LL, (HOOKPROC)(DefaultKeyboardHookHandler(c)), 0, 0)
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
