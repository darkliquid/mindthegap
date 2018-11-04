package scenes

import (
	"github.com/darkliquid/mindthegap/world"
	termbox "github.com/nsf/termbox-go"
)

// Map represents the world map and travel screen
type Map struct {
	mapState   []bool
	playerX    int
	playerY    int
	currentMap *world.Map
}

// NewMap returns an initialised Map scene
func NewMap() *Map {
	return &Map{
		mapState:   []bool{true, true, true, true, true},
		currentMap: world.Maps[0],
	}
}

// Render draws the map layers
func (m *Map) Render(delta float64) error {
	for i, wm := range world.Maps {
		if m.mapState[i] {
			wm.Render()
		}
	}
	m.currentMap.Render()
	return nil
}

// Event handles input events for the map scene
func (m *Map) Event(ev termbox.Event) error {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Ch {
		case '1':
			m.mapState[0] = !m.mapState[0]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '2':
			m.mapState[1] = !m.mapState[1]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '3':
			m.mapState[2] = !m.mapState[2]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '4':
			m.mapState[3] = !m.mapState[3]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		case '5':
			m.mapState[4] = !m.mapState[4]
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	return nil
}

// FPS sets the desired fps
func (m *Map) FPS() float64 {
	return 32
}
