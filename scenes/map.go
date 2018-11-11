package scenes

import (
	"fmt"

	"github.com/darkliquid/mindthegap/engine"
	"github.com/darkliquid/mindthegap/world"
	termbox "github.com/nsf/termbox-go"
)

// Map represents the world map and travel screen
type Map struct {
	lineState   []bool
	playerPos   world.Coord
	currentLine *world.Line
	blink       float64
	showStatus  bool
	travelMode  bool
	travelIndex int
	lastStation *world.Station
}

// NewMap returns an initialised Map scene
func NewMap() *Map {
	return &Map{
		lineState:   []bool{true, true, true, true, true},
		currentLine: world.Lines[0],
		playerPos:   world.StationsByName["Livewire"].Pos,
		showStatus:  true,
		lastStation: world.StationsByName["Livewire"],
	}
}

// renderPlayerPos blinks the cell at the player position
func (m *Map) renderPlayerPos(delta float64) error {
	m.blink += delta
	sw, _ := termbox.Size()
	cells := termbox.CellBuffer()
	pX := m.playerPos.X()
	pY := m.playerPos.Y()
	playerCell := cells[pX+pY*sw]
	if m.blink > 1 {
		m.blink = 0
	} else if m.blink > 0.5 {
		termbox.SetCell(pX, pY, playerCell.Ch, termbox.ColorBlack|termbox.AttrBold, m.currentLine.Color)
	} else if m.blink > 0 {
		termbox.SetCell(pX, pY, playerCell.Ch, m.currentLine.Color|termbox.AttrBold, termbox.ColorBlack)
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
	if m.travelIndex >= len(m.lastStation.Next[m.currentLine]) {
		m.travelIndex = 0
	}
	cells := termbox.CellBuffer()
	to := m.lastStation.Next[m.currentLine][m.travelIndex].To.Pos
	sw, _ := termbox.Size()
	toCell := cells[to.X()+to.Y()*sw]
	if m.blink > 0.5 {
		termbox.SetCell(to.X(), to.Y(), toCell.Ch, termbox.ColorBlack|termbox.AttrBold, m.currentLine.Color)
	} else if m.blink > 0 {
		termbox.SetCell(to.X(), to.Y(), toCell.Ch, m.currentLine.Color|termbox.AttrBold, termbox.ColorBlack)
	}
	for _, coord := range m.lastStation.Next[m.currentLine][m.travelIndex].Path {
		if m.blink > 0.5 {
			termbox.SetCell(coord.X(), coord.Y(), '█', m.currentLine.Color, termbox.ColorBlack)
		} else if m.blink > 0 {
			termbox.SetCell(coord.X(), coord.Y(), '▒', m.currentLine.Color, termbox.ColorBlack)
		}
	}

	m.renderStatusLine(fmt.Sprintf("│ Travel to %v │", m.lastStation.Next[m.currentLine][m.travelIndex].To.Name))
	return nil
}

// Render draws the map layers
func (m *Map) Render(delta float64) error {
	for i, wm := range world.Lines {
		if m.lineState[i] {
			wm.Render()
		}
	}
	if err := m.currentLine.Render(); err != nil {
		return err
	}

	if m.travelMode {
		if err := m.renderTravel(); err != nil {
			return err
		}
	} else if m.showStatus {
		status := fmt.Sprintf("│ Line: %v │", m.currentLine.Name)
		if station, ok := world.StationsByCoord[m.playerPos]; ok {
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
				m.travelIndex = len(m.lastStation.Next[m.currentLine]) - 1
			}
		case termbox.KeyArrowRight, termbox.KeyTab:
			m.travelIndex++
			if m.travelIndex > len(m.lastStation.Next[m.currentLine])-1 {
				m.travelIndex = 0
			}
		case termbox.KeyEnter:
			// TODO: Begin travel to selected destination
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
			if station, ok := world.StationsByCoord[m.playerPos]; ok {
				newLineIdx := 0
				for i, line := range station.Lines {
					if m.currentLine == line {
						newLineIdx = i + 1
						if newLineIdx >= len(station.Lines) {
							newLineIdx = 0
						}
						break
					}
				}
				m.currentLine = station.Lines[newLineIdx]
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
