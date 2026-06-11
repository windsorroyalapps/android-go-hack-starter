package main

import (
	"fmt"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					// Start rendering
				}
			}
		}
	})
}

// Hack mode: Add your BMW OBD or custom Android hooks here!