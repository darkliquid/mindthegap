package engine

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

const (
	// BoxModeOpaque cause boxes to overwrite characters below spaces
	BoxModeOpaque = iota
	// BoxModeTransparent cause boxes to not output spaces over other characters
	BoxModeTransparent
)

// Box is a box of terminal cells that can be rendered at an offset
type Box struct {
	X, Y, W, H int
	Mode       int

	cells [][]termbox.Cell
}

// NewBoxFromString builds a new Box from a multiline string
func NewBoxFromString(s string, fg, bg termbox.Attribute) *Box {
	box := &Box{Mode: BoxModeOpaque}
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
	sw, sh := termbox.Size()
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			if x < len(b.cells[y]) {
				if b.Mode == BoxModeOpaque || b.cells[y][x].Ch != rune(' ') {
					posX := b.X + x
					posY := b.Y + y
					if posX >= 0 && posY >= 0 && posX <= sw && posY <= sh {
						termbox.SetCell(b.X+x, b.Y+y, b.cells[y][x].Ch, b.cells[y][x].Fg, b.cells[y][x].Bg)
					}
				}
			}
		}
	}

	return nil
}

// SetCell sets a cell within the box
func (b *Box) SetCell(x, y int, chr rune, fg, bg termbox.Attribute) {
	if len(b.cells) < y {
		newCells := make([][]termbox.Cell, y)
		copy(newCells, b.cells)
		b.cells = newCells
		b.H = y + 1
	}
	if len(b.cells[y]) < x {
		newLine := make([]termbox.Cell, x)
		copy(newLine, b.cells[y])
		b.cells[y] = newLine
		b.W = x + 1
	}
	b.cells[y][x] = termbox.Cell{
		Ch: chr,
		Fg: fg,
		Bg: bg,
	}
}

// GetCell gets a cell within the box
func (b *Box) GetCell(x, y int) *termbox.Cell {
	if y < 0 || y >= len(b.cells) {
		return nil
	}
	if x < 0 || x >= len(b.cells[y]) {
		return nil
	}
	return &b.cells[y][x]
}

// HCycle cycles the box horizontally, wrapping the 'image'
func (b *Box) HCycle(offset int) {
	if offset == 0 || offset > b.W || offset < 0-b.W {
		return
	}

	if offset < 0 {
		offset = b.W + offset
	}

	for y, row := range b.cells {
		b.cells[y] = append(row[offset:], row[:offset]...)
	}
}

// HRepeat repeats the box content horizonally until width is reached
func (b *Box) HRepeat(width int) {
	b.W = width
	for y, row := range b.cells {
		target := width
		for {
			if target > len(row) {
				b.cells[y] = append(b.cells[y], row...)
				target -= len(row)
			} else {
				b.cells[y] = append(b.cells[y], row[:target]...)
				break
			}
		}
	}
}
