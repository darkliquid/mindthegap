package engine

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

var (
	exit chan struct{}
	err  error
)

// Init initialises the engine
func Init() error {
	exit = make(chan struct{})
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	return nil
}

// Close closes the engine
func Close() error {
	<-exit
	termbox.Close()
	return err
}

func handleInput(ch chan termbox.Event) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				Exit(nil)
				return
			}
			ch <- ev
		case termbox.EventError, termbox.EventInterrupt:
			Exit(ev.Err)
			return
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

// Loop runs the main game loop
func Loop() error {
	inputCh := make(chan termbox.Event)
	go handleInput(inputCh)

	clock := time.Now()
	timer := time.NewTimer(0)
	for {
		select {
		case newScene := <-changeScene:
			scene = scenes[newScene]
		case ev := <-inputCh:
			if handler, ok := scene.(Evented); ok {
				if err := handler.Event(ev); err != nil {
					Exit(err)
					return err
				}
			}
		case <-exit:
			timer.Stop()
			return err
		case update := <-timer.C:
			delta := update.Sub(clock).Seconds()
			clock = update
			if renderer, ok := scene.(Renderer); ok {
				if err := renderer.Render(delta); err != nil {
					Exit(err)
					return err
				}
				termbox.Flush()
			}
			sleep := time.Duration((time.Until(update).Seconds()*1000.0)+1000.0/FPS()) * time.Millisecond
			timer.Reset(sleep)
		}
	}
}

// FPS returns the currently desired FPS
func FPS() float64 {
	if limit, ok := scene.(FPSLimited); ok {
		return limit.FPS()
	}
	return 1
}

// Exit sets the engine to exit
func Exit(e error) {
	select {
	case <-exit:
		return
	default:
		err = e
		close(exit)
	}
}
