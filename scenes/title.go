package scenes

import (
	"github.com/darkliquid/mindthegap/engine"
	termbox "github.com/nsf/termbox-go"
)

const (
	TitleNewGame = 0
	TitleExit    = 1
)

// Title contains all the logic for the title screen
type Title struct {
	selectedOption int
}

const (
	mind = `███╗   ███╗██╗███╗   ██╗██████╗
████╗ ████║██║████╗  ██║██╔══██╗
██╔████╔██║██║██╔██╗ ██║██║  ██║
██║╚██╔╝██║██║██║╚██╗██║██║  ██║
██║ ╚═╝ ██║██║██║ ╚████║██████╔╝
╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚═════╝`

	the = `████████╗██╗  ██╗███████╗
╚══██╔══╝██║  ██║██╔════╝
   ██║   ███████║█████╗
   ██║   ██╔══██║██╔══╝
   ██║   ██║  ██║███████╗
   ╚═╝   ╚═╝  ╚═╝╚══════╝`
	gap = ` ██████╗  █████╗ ██████╗
██╔════╝ ██╔══██╗██╔══██╗
██║  ███╗███████║██████╔╝
██║   ██║██╔══██║██╔═══╝
╚██████╔╝██║  ██║██║
 ╚═════╝ ╚═╝  ╚═╝╚═╝`
)

// Render renders the scene
func (title *Title) Render(delta float64) error {
	m := engine.NewBoxFromString(mind, termbox.ColorGreen, termbox.ColorDefault)
	t := engine.NewBoxFromString(the, termbox.ColorGreen, termbox.ColorDefault)
	g := engine.NewBoxFromString(gap, termbox.ColorGreen, termbox.ColorDefault)
	m.HCenter()
	t.HCenter()
	g.HCenter()
	totalHeight := m.H + t.H + g.H
	_, sh := termbox.Size()
	m.Y = sh/2 - totalHeight/2
	t.Y = m.Y + m.H
	g.Y = t.Y + t.H
	m.Render()
	t.Render()
	g.Render()

	ngFg := termbox.ColorDefault
	ngBg := termbox.ColorDefault
	exFg := termbox.ColorDefault
	exBg := termbox.ColorDefault

	switch title.selectedOption {
	case TitleNewGame:
		ngFg = termbox.ColorBlack
		ngBg = termbox.ColorGreen
	case TitleExit:
		exFg = termbox.ColorBlack
		exBg = termbox.ColorGreen
	}

	ng := engine.NewBoxFromString("New Game", ngFg, ngBg)
	ng.Y = g.Y + g.H + 1
	ng.HCenter()
	ng.Render()
	ex := engine.NewBoxFromString("Exit", exFg, exBg)
	ex.Y = ng.Y + ng.H + 1
	ex.HCenter()
	ex.Render()

	return nil
}

// Event handles events for the title screen
func (title *Title) Event(ev termbox.Event) error {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			title.selectedOption--
			if title.selectedOption < 0 {
				title.selectedOption = 1
			}
		case termbox.KeyArrowDown:
			title.selectedOption++
			if title.selectedOption > 1 {
				title.selectedOption = 0
			}
		case termbox.KeyEnter:
			switch title.selectedOption {
			case TitleNewGame:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				return engine.SetScene("intro")
			case TitleExit:
				engine.Exit(nil)
			}
		}
	}
	return nil
}

// FPS returns the desired FPS for this scene
func (title *Title) FPS() float64 {
	return 8
}
