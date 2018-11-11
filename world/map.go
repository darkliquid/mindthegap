package world

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/darkliquid/mindthegap/engine"
	termbox "github.com/nsf/termbox-go"
)

// rawLines represent all the line in the world
//
// Each set of 50 lines is a different subway line.
// Numbers represent colours, letters represent the
// different stations that can be stopped at or used
// to switch line.
var rawLines = []string{`



   A▒▒▒
       ▒                 D▒▒▒
        ▒               ▒    ▒
         ▒             ▒      ▒
          B▒▒▒▒▒▒▒    ▒        ▒            ▒▒▒▒
                  C▒▒▒          E▒▒▒▒▒▒▒▒▒▒▒    F
                                             ▒▒▒
                                          ▒▒▒
                                       ▒G▒
                                     ▒▒
                                    ▒
                                   ▒
                                   ▒
                                   H
                                  ▒
                                 ▒
                                ▒
                                ▒
               K▒               ▒
              ▒  ▒              ▒
              ▒  ▒             I
              ▒   ▒▒     ▒▒▒▒▒▒
             ▒      ▒▒▒▒J
            ▒
         ▒▒L
        ▒
       ▒
      ▒
     M
      ▒
       ▒
        ▒▒▒▒▒▒N















`, `
















                                  ▒H
                   T▒▒▒▒▒▒▒▒▒▒▒▒▒▒  ▒
                  ▒                ▒
                 ▒                ▒
                ▒                ▒
              ▒K                 ▒
             ▒                  ▒
            ▒                  I
           ▒             ▒▒▒▒▒▒
           ▒            J
           ▒           ▒
           L          ▒
           ▒         ▒
          ▒          ▒
         ▒           ▒
    ▒M▒▒▒            S
   ▒                 ▒
  ▒                   ▒
 ▒        ▒▒▒▒N        ▒
▒        ▒     ▒        ▒▒
▒        ▒     ▒          R
O ▒▒▒▒▒▒▒      ▒           ▒▒▒▒▒▒▒▒▒
 ▒             ▒                    Q
              ▒  ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
              ▒ ▒
               P







`, `
                                                     ▒
                                                    i ▒▒▒▒▒▒▒▒
                                                    ▒         ▒
                                                    ▒          h
                                                    ▒       ▒▒▒
                                                    ▒      ▒
                                                   ▒      ▒
                                                  ▒      ▒
                             ▒▒▒E               F▒      g
                            ▒    ▒           ▒▒▒       ▒
                           ▒     ▒        ▒▒▒         ▒
                          ▒       ▒▒▒▒▒▒G▒           ▒
                        U▒                          ▒
                       ▒                           ▒
                     ▒▒                           ▒
                   ▒▒               ▒▒▒▒         ▒
                  ▒                H    ▒       f
                   T▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒      ▒       ▒
                                          ▒       ▒▒▒▒▒
                                           ▒           e
                                            ▒V▒        ▒
                                               ▒       ▒
                                                ▒▒      ▒
                                                  ▒      ▒
                                             ▒▒▒▒W        d
                                            ▒            ▒
                                           X            ▒
                                           ▒            ▒
                                           ▒            ▒
                                            ▒           ▒
                                             ▒          c
                                              Y          ▒
                                              ▒           ▒
                                              ▒            ▒
                                              ▒             ▒
                                              ▒             b
                                              ▒  ▒▒▒▒a▒▒▒▒▒▒
                                            ▒▒  ▒
                                           ▒   ▒
                                            Z▒▒










`, `
             q
            ▒ ▒
            ▒  ▒▒
            ▒    ▒                                            ▒h
            ▒    ▒       D                                   ▒  ▒
            ▒    ▒      ▒ ▒                                 ▒   ▒
           ▒     ▒     ▒  ▒                                ▒    ▒
          B      ▒    ▒  ▒                                ▒     ▒
          ▒       C▒▒▒  ▒    ▒▒▒E               F▒▒▒▒▒▒▒g▒      ▒
          ▒             ▒   ▒    ▒           ▒▒▒                ▒
         ▒              ▒  ▒     ▒        ▒▒▒                   ▒
        k              ▒  ▒       ▒▒▒▒▒▒G▒                      ▒
       ▒                U▒                                      j
      ▒      o                                               ▒▒▒
     ▒      ▒ ▒                                             ▒
  ▒▒▒      ▒   ▒                                           ▒
 l        ▒     ▒▒▒                            ▒f         ▒
  ▒      n         T                          ▒  ▒       ▒
   ▒    ▒          ▒                         ▒    ▒     ▒
    ▒  ▒            ▒                        ▒     ▒▒▒▒e
     m▒              ▒                       V
                      ▒                     ▒
                       ▒                    ▒
                        ▒                   ▒
                        ▒                  ▒
                        J                  ▒
                         ▒                 X
                          ▒                 ▒▒▒▒▒
                           ▒                     ▒
                            ▒                     ▒
                             ▒▒▒p▒▒▒▒▒▒▒▒▒▒▒▒▒     ▒
                                              Y     ▒
                                              ▒      ▒
                                              ▒      ▒
                                              ▒      ▒
                                              ▒      ▒
                                              ▒  ▒▒▒▒a
                                            ▒▒  ▒
                                           ▒   ▒
                                            Z▒▒










`, `
             q                                       ▒
            ▒ ▒                                     i ▒▒▒▒▒▒▒▒
            ▒  ▒▒                                   ▒         ▒
   A▒▒▒     ▒    ▒                                  ▒          h
       ▒    ▒    ▒       D                          ▒       ▒▒▒
        ▒   ▒    ▒      ▒ ▒                         ▒      ▒
         ▒ ▒     ▒     ▒  ▒                        ▒      ▒
          B      ▒    ▒  ▒                        ▒      ▒
                  C▒▒▒  ▒    ▒▒▒E               F▒      g
                        ▒   ▒    ▒           ▒▒▒       ▒
                        ▒  ▒     ▒        ▒▒▒         ▒
                       ▒  ▒       ▒▒▒▒▒▒G▒           ▒  ▒▒▒▒▒▒▒▒
                        U▒                          ▒  ▒        j
                                                   ▒  ▒         ▒
                                                  ▒  ▒          ▒
                                                 ▒  ▒           ▒
                                                f  ▒            ▒
                                                 ▒▒             ▒
                                             ▒▒▒                r
                                            ▒   ▒▒▒▒▒▒▒e▒▒▒▒▒▒▒▒
                                             V▒
                                               ▒
                                                ▒▒
                                                  ▒   ▒▒▒▒
                                                 W   ▒    d
                                                  ▒▒▒    ▒
                                           X            ▒
                                           ▒            ▒
                                           ▒            ▒
                                            ▒           ▒
                                             ▒          c
                                              Y          ▒
                                              ▒           ▒
                                              ▒            ▒
                                              ▒             ▒
                                              ▒             b
                                              ▒  ▒▒▒▒a▒▒▒▒▒▒
                                            ▒▒  ▒
                                           ▒   ▒
                                            Z▒▒










`}

// LineNames are the names of each subway line
var LineNames = []string{"Livewire", "Spanner", "Neck", "Jug", "Soot & Glass"}

// stationNames map station line codes to full names
var stationNames = map[rune]string{
	'A': "Fuse",
	'B': "Waterloo",
	'C': "Porterville",
	'D': "Old Trick Tavern",
	'E': "Four-point Station",
	'F': "Lower Bridge",
	'G': "The Bridge",
	'H': "Livewire",
	'I': "Cablegate",
	'J': "Empire",
	'K': "Deadmans Dock",
	'L': "Vagrants Point",
	'M': "Armoury",
	'N': "Spawns End",
	'O': "The Spanner",
	'P': "Pentville",
	'Q': "Devils Switchback",
	'R': "Lonesome",
	'S': "Runaway Split",
	'T': "Ducks Back",
	'U': "Guillotine",
	'V': "Xenmouth",
	'W': "Shithole",
	'X': "Sandbank",
	'Y': "Rivers Reach",
	'Z': "Bohemian Alley",
	'a': "Oldhaven",
	'b': "Stains",
	'c': "Newhaven",
	'd': "Boiardi Joint",
	'e': "Breakers Crossing",
	'f': "Quayside",
	'g': "Lower Mead",
	'h': "Upper Mead",
	'i': "Trent",
	'j': "Diamondback",
	'k': "Lid Island",
	'l': "Spout",
	'm': "Jughead",
	'n': "Pitcher Point",
	'o': "Handle",
	'p': "Old Dog",
	'q': "Highpoint",
	'r': "Soot & Glass",
}

// Station represents a train station
type Station struct {
	Name  string
	Code  rune
	Pos   Coord
	Lines []*Line
	Next  map[*Line][]*Segment
}

// Segment represents the path between two stations
type Segment struct {
	Path []Coord
	From *Station
	To   *Station
}

// Coord is an X/Y coordinate
type Coord [2]int

// X is the X coordinate
func (c Coord) X() int {
	return c[0]
}

// Y is the Y coordinate
func (c Coord) Y() int {
	return c[1]
}

// Equal tests coord equality
func (c Coord) Equal(c2 Coord) bool {
	return c.X() == c2.X() && c.Y() == c2.Y()
}

// NewCoord returns a Coord
func NewCoord(x, y int) Coord {
	return Coord([2]int{x, y})
}

// StationsByCoord Lines co-ordinates to station reference
var StationsByCoord map[Coord]*Station

// StationsByName Lines names to station reference
var StationsByName map[string]*Station

var lineColours = []termbox.Attribute{
	termbox.ColorYellow,
	termbox.ColorGreen,
	termbox.ColorRed,
	termbox.ColorBlue,
	termbox.ColorMagenta,
}

// Line is a named Line
type Line struct {
	Name  string
	Color termbox.Attribute
	*engine.Box
}

// Lines is the array of all Lines in the game
var Lines []*Line

var re = regexp.MustCompile("[a-zA-Z]")

// Since Lines are complicated and contain a bunch of info, first we extract it and cache
// it as doing it every time is a pain
func init() {
	Lines = make([]*Line, len(rawLines))
	StationsByCoord = make(map[Coord]*Station)
	StationsByName = make(map[string]*Station)

	// Process the raw Line data
	for i, m := range rawLines {
		b := engine.NewBoxFromString(m, lineColours[i], termbox.ColorDefault)
		b.Mode = engine.BoxModeTransparent
		Lines[i] = &Line{Name: LineNames[i], Box: b, Color: lineColours[i]}
		lines := strings.Split(m, "\n")
		for y, line := range lines {
			indexes := re.FindAllStringIndex(line, -1)
			for _, index := range indexes {
				runeOffset := utf8.RuneCountInString(line[:index[0]])
				stationCode := rune(line[index[0]])
				coord := NewCoord(runeOffset, y)
				if _, ok := StationsByCoord[coord]; !ok {
					StationsByCoord[coord] = &Station{Lines: make([]*Line, 0), Next: make(map[*Line][]*Segment)}
					StationsByName[stationNames[stationCode]] = StationsByCoord[coord]
				}
				StationsByCoord[coord].Name = stationNames[stationCode]
				StationsByCoord[coord].Code = stationCode
				StationsByCoord[coord].Pos = coord
				StationsByCoord[coord].Lines = append(StationsByCoord[coord].Lines, Lines[i])
				b.SetCell(runeOffset, y, stationCode, termbox.ColorBlack|termbox.AttrBold, lineColours[i])
			}
		}
	}

	// Build next/prev mappings
	for _, station := range StationsByCoord {
		for _, line := range station.Lines {
			station.Next[line] = FindSegments(station, line)
		}
	}
}

// FindSegments scans the line for the adjacent stations
func FindSegments(current *Station, line *Line) []*Segment {
	x, y := current.Pos.X(), current.Pos.Y()

	dirs := []Coord{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

	s := make([]*Segment, 0)
	for _, dir := range dirs {
		newCoord := NewCoord(x+dir.X(), y+dir.Y())
		cell := line.GetCell(newCoord.X(), newCoord.Y())
		if cell == nil {
			continue
		}
		if cell.Ch == rune('▒') {
			s = append(s, &Segment{
				From: current,
				Path: []Coord{newCoord},
			})
		}

		if re.MatchString(string(cell.Ch)) {
			s = append(s, &Segment{
				From: current,
				To:   StationsByCoord[newCoord],
			})
		}
	}

	for _, segment := range s {
		done := false
		if segment.To != nil {
			continue
		}
		last := segment.Path[0]
		for !done {
			for i, dir := range dirs {
				newCoord := NewCoord(last.X()+dir.X(), last.Y()+dir.Y())
				if newCoord.Equal(segment.From.Pos) {
					segment.Path = segment.Path[:1]
					if i == len(dirs)-1 {
						done = true
					}
					continue
				}

				if len(segment.Path) > 1 {
					if newCoord.Equal(segment.Path[len(segment.Path)-2]) {
						if i == len(dirs)-1 {
							done = true
						}
						continue
					}
				}

				cell := line.GetCell(newCoord.X(), newCoord.Y())
				if cell == nil {
					if i == len(dirs)-1 {
						done = true
					}
					continue
				}

				if cell.Ch == rune('▒') {
					segment.Path = append(segment.Path, newCoord)
					last = newCoord
					break
				}

				if re.MatchString(string(cell.Ch)) {
					segment.To = StationsByCoord[newCoord]
					done = true
					last = segment.From.Pos
					break
				}

				if i == len(dirs)-1 {
					done = true
				}
			}
		}

		if segment.To == nil {
			// All segments should have Froms and Tos, so if not, the map is at fault
			// and needs fixing
			panic(fmt.Sprint(line.Name, segment.From.Name, segment.Path))
		}
	}

	return s
}
