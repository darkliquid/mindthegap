package engine

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// Box is a box of terminal cells that can be rendered at an offset
type Box struct {
	X, Y int

	W, H  int
	cells [][]termbox.Cell
}

// NewBoxFromString builds a new Box from a multiline string
func NewBoxFromString(s string, fg, bg termbox.Attribute) *Box {
	box := &Box{}
	lines := strings.Split(s, "\n")
	box.cells = make([][]termbox.Cell, len(lines))
	for _, line := range lines {
		l := len([]rune(line))
		if l > box.W {
			box.W = l
		}
	}
	box.H = len(lines)

	for y := 0; y < box.H; y++ {
		box.cells[y] = make([]termbox.Cell, box.W)
		for x := 0; x < box.W; x++ {
			box.cells[y][x].Fg = fg
			box.cells[y][x].Bg = bg
			runes := []rune(lines[y])
			if x < len(runes) {
				box.cells[y][x].Ch = rune(runes[x])
			} else {
				box.cells[y][x].Ch = rune(' ')
			}

		}
	}

	return box
}

// HCenter centers the position of the box horizontally
func (b *Box) HCenter() {
	sw, _ := termbox.Size()
	b.X = sw/2 - b.W/2
}

// VCenter centers the position of the box horizontally
func (b *Box) VCenter() {
	_, sh := termbox.Size()
	b.Y = sh/2 - b.H/2
}

// Render renders the box to the terminal at its specified position
func (b *Box) Render() error {
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			if x < len(b.cells[y]) {
				termbox.SetCell(b.X+x, b.Y+y, b.cells[y][x].Ch, b.cells[y][x].Fg, b.cells[y][x].Bg)
			}
		}
	}

	return nil
}
