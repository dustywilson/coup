package player

import "sync"

var players = make(map[Player]Player)
var playerLock sync.RWMutex

// Player is a Coup player (probably AI)
type Player interface {
	TypeName() string
	Name() string
}

// Players returns a list of Player
func Players() []Player {
	playerLock.RLock()
	defer playerLock.RUnlock()
	rList := make([]Player, len(players))
	i := 0
	for _, r := range players {
		rList[i] = r
		i++
	}
	return rList
}

// Add adds a Player registration
func Add(player Player) {
	playerLock.Lock()
	defer playerLock.Unlock()
	if _, ok := players[player]; !ok {
		players[player] = player
	}
}
