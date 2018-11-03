package scenes

import (
	"github.com/darkliquid/mindthegap/engine"
	termbox "github.com/nsf/termbox-go"
)

// Title contains all the logic for the title screen
type Title struct{}

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
	m := engine.NewBoxFromString(mind)
	t := engine.NewBoxFromString(the)
	g := engine.NewBoxFromString(gap)
	m.HCenter()
	t.HCenter()
	g.HCenter()
	t.Y = m.H
	g.Y = m.H + t.H
	m.Render()
	t.Render()
	g.Render()

	return nil
}

func (title *Title) Event(ev termbox.Event) error {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyEsc {
			engine.Exit()
		}
	}
	return nil
}
