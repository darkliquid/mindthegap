package state

import "github.com/darkliquid/mindthegap/world"

var (
	// Player contains the global player state
	Player *playerState
	// World contains the global world state
	World *worldState
)

// init sets up the New Game state
func init() {
	Player = &playerState{}
	World = &worldState{
		CurrentLine:    world.Lines[0],
		PlayerPos:      world.StationsByName["Livewire"].Pos,
		CurrentSegment: world.StationsByName["Livewire"].Next[world.Lines[0]][0],
	}
}

type playerState struct {
}

type worldState struct {
	PlayerPos      world.Coord
	CurrentLine    *world.Line
	CurrentSegment *world.Segment
}
