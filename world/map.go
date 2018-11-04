package world

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/darkliquid/mindthegap/engine"
	termbox "github.com/nsf/termbox-go"
)

// rawMaps represent the world map
//
// Each set of 50 lines is a different subway line.
// Numbers represent colours, letters represent the
// different stations that can be stopped at or used
// to switch line.
var rawMaps = []string{`



▒▒▒A▒▒▒
       ▒                 D▒▒▒
        ▒               ▒    ▒
         ▒             ▒      ▒
          B▒▒▒▒▒▒▒    ▒        ▒            ▒▒▒▒
                  C▒▒▒          E▒▒▒▒▒▒▒▒▒▒▒    F
                                             ▒▒
                                          ▒▒▒
                                       ▒G▒
                                     ▒▒
                                   ▒▒
                                   ▒
                                   ▒
                                   H
                                  ▒
                                 ▒
                                ▒
                                ▒
               K▒               ▒
               ▒ ▒              ▒
               ▒  ▒             I
              ▒    ▒     ▒▒▒▒▒▒▒
             ▒      ▒▒▒▒J
            ▒
         ▒▒L
        ▒
       ▒
      ▒
     M
      ▒
       ▒
        ▒▒▒▒▒▒N▒▒▒















`, `
















                                  ▒H
                   T▒▒▒▒▒▒▒▒▒▒▒▒▒▒  ▒
                  ▒                 ▒
                 ▒                  ▒
                ▒                   ▒
              ▒K                   ▒
             ▒                    ▒
            ▒             ▒▒▒▒▒▒I▒
           ▒             ▒
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
▒        ▒     ▒        ▒
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
                            ▒▒▒▒E               F▒      g
                           ▒     ▒           ▒▒▒       ▒
                          ▒      ▒        ▒▒▒         ▒
                         ▒        ▒▒▒▒▒▒G▒           ▒
                        U                           ▒
                        ▒                          ▒
                     ▒▒▒                          ▒
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

// mapNames are the names of each subway line
var mapNames = []string{"Livewire", "Spanner", "Neck", "Jug", "Soot & Glass"}

// stationNames map station map codes to full names
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
	X, Y  int
	Lines []bool
}

// Stations maps co-ordinates to station reference
var Stations map[[2]int]*Station

var lineColours = []termbox.Attribute{
	termbox.ColorYellow,
	termbox.ColorGreen,
	termbox.ColorRed,
	termbox.ColorBlue,
	termbox.ColorMagenta,
}

// Map is a named map
type Map struct {
	Name string
	*engine.Box
}

// Maps is the array of all maps in the game
var Maps []*Map

// Since maps are complicated and contain a bunch of info, first we extract it and cache
// it as doing it every time is a pain
func init() {
	Maps = make([]*Map, len(rawMaps))
	Stations = make(map[[2]int]*Station)
	re := regexp.MustCompile("[a-zA-Z]")

	// Process the raw map data
	for i, m := range rawMaps {
		b := engine.NewBoxFromString(m, lineColours[i], termbox.ColorDefault)
		b.Mode = engine.BoxModeTransparent
		Maps[i] = &Map{Name: mapNames[i], Box: b}
		lines := strings.Split(m, "\n")
		for y, line := range lines {
			indexes := re.FindAllStringIndex(line, -1)
			for _, index := range indexes {
				runeOffset := utf8.RuneCountInString(line[:index[0]])
				stationCode := rune(line[index[0]])
				if _, ok := Stations[[2]int{runeOffset, y}]; !ok {
					Stations[[2]int{runeOffset, y}] = &Station{Lines: make([]bool, len(rawMaps))}
				}
				Stations[[2]int{runeOffset, y}].Name = stationNames[stationCode]
				Stations[[2]int{runeOffset, y}].X = runeOffset
				Stations[[2]int{runeOffset, y}].Y = y
				Stations[[2]int{runeOffset, y}].Lines[i] = true
				b.SetCell(runeOffset, y, stationCode, termbox.ColorBlack|termbox.AttrBold, lineColours[i])
			}
		}
	}
}
