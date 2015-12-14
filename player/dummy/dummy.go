package dummy

import "github.com/dustywilson/coup/player"

// Dummy is a Player
type Dummy struct {
	name string
}

func init() {
	player.Add(Dummy{})
}

// Name is the Dummy player's name
func (d Dummy) Name() string {
	return d.name
}

// TypeName is the Dummy's typeName
func (d Dummy) TypeName() string {
	return "Dummy"
}
