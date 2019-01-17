package main

import (
	"os"
	"os/exec"

	termbox "github.com/nsf/termbox-go"
)

func inputLoop() {
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyCtrlQ || ev.Key == termbox.KeyEsc {
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()
				os.Exit(3)
			}
		}
	}
}
