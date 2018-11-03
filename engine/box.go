package engine

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// Box is a box of terminal cells that can be rendered at an offset
type Box struct {
	X, Y int

	W, H  int
	cells [][]rune
}

// NewBoxFromString builds a new Box from a multiline string
func NewBoxFromString(s string) *Box {
	box := &Box{}
	lines := strings.Split(s, "\n")
	box.cells = make([][]rune, len(lines))
	for i, line := range lines {
		l := len(line)
		box.cells[i] = []rune(line)
		if l > box.W {
			box.W = l
		}
	}
	for i, line := range box.cells {
		newLine := make([]rune, box.W)
		copy(newLine, line)
		box.cells[i] = newLine
	}
	box.H = len(lines)
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

// Render renders the box to the terinal at its specified position
func (b *Box) Render() error {
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			termbox.SetCell(b.X+x, b.Y+y, b.cells[y][x], termbox.ColorGreen, termbox.ColorDefault)
		}
	}

	return nil
}
