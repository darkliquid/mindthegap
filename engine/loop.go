package engine

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

var exit bool

func Init() error {
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

func Close() error {
	termbox.Close()
	return nil
}

func input(ch chan termbox.Event) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			ch <- ev
		case termbox.EventError, termbox.EventInterrupt:
			ch <- ev
			return
		}
	}
}

func Loop() error {
	inputEvent := make(chan termbox.Event)
	go input(inputEvent)

	clock := time.Now()
	for {
		update := time.Now()
		delta := update.Sub(clock).Seconds()
		clock = update

		select {
		case ev := <-inputEvent:
			if handler, ok := scene.(Evented); ok {
				if err := handler.Event(ev); err != nil {
					return err
				}
			}
			switch ev.Type {
			case termbox.EventError:
				return ev.Err
			case termbox.EventInterrupt:
				return nil
			}
		default:
			if renderer, ok := scene.(Renderer); ok {
				if err := renderer.Render(delta); err != nil {
					return err
				}
				termbox.Flush()
			}
		}

		if exit {
			return nil
		}

		sleep := time.Duration((time.Until(update).Seconds()*1000.0)+1000.0/FPS()) * time.Millisecond
		time.Sleep(sleep)
	}
}

func FPS() float64 {
	if limit, ok := scene.(FPSLimited); ok {
		return limit.FPS()
	}
	return 1
}

func Exit() {
	exit = true
}
