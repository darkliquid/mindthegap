package scenes

import (
	"strings"

	"github.com/darkliquid/mindthegap/engine"
	termbox "github.com/nsf/termbox-go"
)

// Intro is the introduction scene when starting a new game
type Intro struct {
	cursor int
}

var introTxt = []rune(`Nobody remembers how or why. Maybe nobody ever did. 
The oldest living survivors have childhood memories painted in shades
of fire, blood and shit. Chaos and death and an endless howling as
the world split open and something came crawling out. Nothing real,
something in the mind. The worlds subconscious hate and terror and
lust, spilling out of an open wound into a psychic storm from which
none of us had shelter.

They called it the Gap. The space between places. Between heaven and
hell, love and hate, life and death. The edge of the platform and subway
train turning your face into paste. And out of the Gap poured pain and
misery and hunger made manifest. Out of the Gap came the Wolves.

But that was long ago, and while the Wolves are real, mankind lives on,
safe, secure, in the Underground. The ones before lined this lands
with tunnels, with engines that tore between them and wires to speak
secret messages rendered in lightning and fire. It is a harsh world,
but this group of strange people find themselves banded together by
need, by fate, by the only things left that matter.
                                                                      
                                                   And so it begins...`)

var width int

func init() {
	for _, line := range strings.Split(string(introTxt), "\n") {
		l := len(line)
		if l > width {
			width = l
		}
	}
}

// Render renders the intro text
func (i *Intro) Render(delta float64) error {
	output := string(introTxt[:i.cursor])
	ib := engine.NewBoxFromString(output, termbox.ColorDefault, termbox.ColorDefault)
	ib.W = width
	ib.HCenter()
	ib.VCenter()

	if i.cursor == 0 {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		i.cursor++
	} else if i.cursor < len(introTxt) {
		i.cursor++
	} else {
		button := engine.NewBoxFromString("Press enter to continue", termbox.ColorBlack, termbox.ColorGreen)
		button.Y = ib.Y + ib.H + 1
		button.X = ib.X + width - button.W
		button.Render()
	}

	ib.Render()

	return nil
}

// FPS sets the desired FPS
func (i *Intro) FPS() float64 {
	return 32
}
