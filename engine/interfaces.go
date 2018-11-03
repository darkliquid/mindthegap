package engine

import (
	termbox "github.com/nsf/termbox-go"
)

// Renderer is something that can render to the terminal
type Renderer interface {
	Render(float64) error
}

// Evented is something that processes events
type Evented interface {
	Event(termbox.Event) error
}

// FPSLimited is something that limits the FPS
type FPSLimited interface {
	FPS() float64
}
