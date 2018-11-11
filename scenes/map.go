package scenes

import (
	"fmt"

	"github.com/darkliquid/mindthegap/engine"
	"github.com/darkliquid/mindthegap/state"
	"github.com/darkliquid/mindthegap/world"
	termbox "github.com/nsf/termbox-go"
)

// Map represents the world map and travel screen
type Map struct {
	lineState   []bool
	blink       float64
	showStatus  bool
	travelMode  bool
	travelIndex int
}

// NewMap returns an initialised Map scene
func NewMap() *Map {
	return &Map{
		lineState:  []bool{true, true, true, true, true},
		showStatus: true,
	}
}

// renderPlayerPos blinks the cell at the player position
func (m *Map) renderPlayerPos(delta float64) error {
	m.blink += delta
	sw, _ := termbox.Size()
	cells := termbox.CellBuffer()
	pX := state.World.PlayerPos.X()
	pY := state.World.PlayerPos.Y()
	playerCell := cells[pX+pY*sw]
	if m.blink > 1 {
		m.blink = 0
	} else if m.blink > 0.5 {
		termbox.SetCell(pX, pY, playerCell.Ch, termbox.ColorBlack|termbox.AttrBold, state.World.CurrentLine.Color)
	} else if m.blink > 0 {
		termbox.SetCell(pX, pY, playerCell.Ch, state.World.CurrentLine.Color|termbox.AttrBold, termbox.ColorBlack)
	}
	return nil
}

// renderStatusLine draws the status line/info bar on the screen
func (m *Map) renderStatusLine(msg string) error {
	sw, sh := termbox.Size()
	w := sw - 8
	msg = fmt.Sprintf("█▓▒░%[1]*s░▒▓█", -w, fmt.Sprintf("%[1]*s", (w+len(msg))/2, msg))
	b := engine.NewBoxFromString(msg, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
	b.Y = sh - 1
	return b.Render()
}

// renderTravel highlights the path to the selected destination
func (m *Map) renderTravel() error {
	cells := termbox.CellBuffer()
	to := state.World.CurrentSegment.To.Pos
	sw, _ := termbox.Size()
	toCell := cells[to.X()+to.Y()*sw]
	if m.blink > 0.5 {
		termbox.SetCell(to.X(), to.Y(), toCell.Ch, termbox.ColorBlack|termbox.AttrBold, state.World.CurrentLine.Color)
	} else if m.blink > 0 {
		termbox.SetCell(to.X(), to.Y(), toCell.Ch, state.World.CurrentLine.Color|termbox.AttrBold, termbox.ColorBlack)
	}
	for _, coord := range state.World.CurrentSegment.Path {
		if m.blink > 0.5 {
			termbox.SetCell(coord.X(), coord.Y(), '█', state.World.CurrentLine.Color, termbox.ColorBlack)
		} else if m.blink > 0 {
			termbox.SetCell(coord.X(), coord.Y(), '▒', state.World.CurrentLine.Color, termbox.ColorBlack)
		}
	}

	m.renderStatusLine(fmt.Sprintf("│ Travel to %v │", state.World.CurrentSegment.To.Name))
	return nil
}

// Render draws the map layers
func (m *Map) Render(delta float64) error {
	for i, wm := range world.Lines {
		if m.lineState[i] {
			wm.Render()
		}
	}
	if err := state.World.CurrentLine.Render(); err != nil {
		return err
	}

	if m.travelMode {
		if err := m.renderTravel(); err != nil {
			return err
		}
	} else if m.showStatus {
		status := fmt.Sprintf("│ Line: %v │", state.World.CurrentLine.Name)
		if station, ok := world.StationsByCoord[state.World.PlayerPos]; ok {
			status += fmt.Sprintf(" Station: %v │", station.Name)
		}
		if err := m.renderStatusLine(status); err != nil {
			return err
		}
	}

	return m.renderPlayerPos(delta)
}

// Event handles input events for the map scene
func (m *Map) Event(ev termbox.Event) error {
	if m.travelMode {
		if ev.Ch == 't' || ev.Ch == 'T' {
			m.travelMode = false
			return nil
		}
		switch ev.Key {
		case termbox.KeyArrowLeft:
			m.travelIndex--
			if m.travelIndex < 0 {
				m.travelIndex = len(state.World.CurrentSegment.From.Next[state.World.CurrentLine]) - 1
			}
			state.World.CurrentSegment = state.World.CurrentSegment.From.Next[state.World.CurrentLine][m.travelIndex]
		case termbox.KeyArrowRight, termbox.KeyTab:
			m.travelIndex++
			if m.travelIndex > len(state.World.CurrentSegment.From.Next[state.World.CurrentLine])-1 {
				m.travelIndex = 0
			}
			state.World.CurrentSegment = state.World.CurrentSegment.From.Next[state.World.CurrentLine][m.travelIndex]
		case termbox.KeyEnter:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			return engine.SetScene("travel")
		}
	}
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Ch {
		// Toggle line visibility
		case '1':
			m.lineState[0] = !m.lineState[0]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '2':
			m.lineState[1] = !m.lineState[1]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '3':
			m.lineState[2] = !m.lineState[2]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '4':
			m.lineState[3] = !m.lineState[3]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '5':
			m.lineState[4] = !m.lineState[4]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		// Change current line (only at stations)
		case 'c', 'C':
			if station, ok := world.StationsByCoord[state.World.PlayerPos]; ok {
				newLineIdx := 0
				for i, line := range station.Lines {
					if state.World.CurrentLine == line {
						newLineIdx = i + 1
						if newLineIdx >= len(station.Lines) {
							newLineIdx = 0
						}
						break
					}
				}
				state.World.CurrentLine = station.Lines[newLineIdx]
				next := state.World.CurrentSegment.From.Next[state.World.CurrentLine]
				if m.travelIndex > len(next)-1 {
					m.travelIndex = 0
				}
				state.World.CurrentSegment = next[m.travelIndex]
			}
		// Toggle status line
		case 's':
			m.showStatus = !m.showStatus
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		// Enter travel mode
		case 't', 'T':
			m.travelMode = true
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	return nil
}

// FPS sets the desired fps
func (m *Map) FPS() float64 {
	return 32
}
