package console

import (
	"syscall"
)

var (
	getWin  = syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	showWin = syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
)

func Show() {
	const SW_RESTORE uintptr = 9

	hWnd, _, _ := getWin.Call()
	if hWnd == 0 {
		return
	}

	showWin.Call(hWnd, SW_RESTORE)
}

func Hide() {
	const SW_HIDE uintptr = 0

	hWnd, _, _ := getWin.Call()
	if hWnd == 0 {
		return
	}

	showWin.Call(hWnd, SW_HIDE)
}
