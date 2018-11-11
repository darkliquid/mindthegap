package scenes

import (
	"fmt"

	"github.com/darkliquid/mindthegap/engine"
	"github.com/darkliquid/mindthegap/state"
	termbox "github.com/nsf/termbox-go"
)

const train = `
___________   __________________________________________
 ___   ___ |||  ___   ___   ___    ___ ___  |   __  ,----\
|   | |   |||| |   | |   | |   |  |   |   | |  |  | |_____\
|___| |___|||| |___| |___| |___|  | O | O | |  |  |        \
           |||                    |___|___| |  |__|         )
___________|||______________________________|______________/
           |||                                        /
-----------'''---------------------------------------'`

const tunnel = `┬────────┬────────
│        º
│
│
│
│
│
│
│
│
╬════════╬════════`

// Travel is the travelling mode screen
type Travel struct {
	progress      float64
	tunnelCounter float64
}

// Render draws the travel scene
func (t *Travel) Render(delta float64) error {
	last := state.World.CurrentSegment.From
	next := state.World.CurrentSegment.To

	lastBox := engine.NewBoxFromString(last.Name, state.World.CurrentLine.Color, termbox.ColorDefault)
	lastBox.HCenter()
	lastBox.Y = 2

	toBox := engine.NewBoxFromString("to", termbox.ColorDefault, termbox.ColorDefault)
	toBox.HCenter()
	toBox.Y = lastBox.Y + lastBox.H + 1

	nextBox := engine.NewBoxFromString(next.Name, state.World.CurrentLine.Color, termbox.ColorDefault)
	nextBox.HCenter()
	nextBox.Y = toBox.Y + toBox.H + 1

	lastBox.Render()
	toBox.Render()
	nextBox.Render()

	t.progress += delta
	length := float64(len(state.World.CurrentSegment.Path) - 1)
	if t.progress >= length {
		t.progress = 0
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		state.World.CurrentSegment = state.World.CurrentSegment.To.Next[state.World.CurrentLine][0]
		state.World.PlayerPos = state.World.CurrentSegment.From.Pos
		return engine.SetScene("map")
	}

	state.World.PlayerPos = state.World.CurrentSegment.Path[int(t.progress)]

	progressBox := engine.NewBoxFromString(fmt.Sprintf("Journey %d%% complete", int((t.progress/length)*100)), termbox.ColorDefault, termbox.ColorDefault)
	progressBox.HCenter()
	progressBox.Y = nextBox.Y + nextBox.H + 1
	progressBox.Render()

	t.tunnelCounter += delta * 32
	tunBox := engine.NewBoxFromString(tunnel, termbox.ColorWhite, termbox.ColorDefault)
	sw, _ := termbox.Size()
	tunBox.HRepeat(sw)
	tunBox.VCenter()
	tunBox.HCycle(int(t.tunnelCounter))
	if t.tunnelCounter > 18 {
		t.tunnelCounter = 0
	}
	tunBox.Render()

	trainBox := engine.NewBoxFromString(train, termbox.ColorWhite, termbox.ColorDefault)
	trainBox.Mode = engine.BoxModeTransparent
	trainBox.Y = tunBox.Y + 1
	trainBox.Render()

	return nil
}

// FPS defines the FPS for this scene
func (t *Travel) FPS() float64 {
	return 32
}
