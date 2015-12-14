package role

import "sync"

var roles = make(map[string]Role)
var roleLock sync.RWMutex

// Role Categories
const (
	Communications = "Communications"
	Finance        = "Finance"
	Force          = "Force"
	SI             = "Special Interests"
)

// Role is a Coup role
type Role interface {
	ID() string // must be a stable ID, no matter what; used in hard-coded references between roles, etc; not exposed to end users
	Category() string
	Name() string // should be internationalized somehow; does not need to be a permanently stable ID
	IsAdvanced() bool
	IsAction() bool
	ProvidesExtraAction() bool
	BlocksAnytime() []string     // Block* values are recorded separately (blocker role vs blocked role)
	BlocksIfTarget() []string    // ...
	BlockedByAnytime() []string  // ...
	BlockedByIfTarget() []string // ... because any are legitimate when matched, even if not in between blocker/blocked role (Roles by different authors, etc)
	PayToBank() int
	EarnFromBank() int
	TakeFromTargetPlayer() int
	TakeFromAllPlayers() int
	GiveToTargetPlayer() int
	GiveToAllPlayers() int
}

// Roles returns a list of Role
func Roles() []Role {
	roleLock.RLock()
	defer roleLock.RUnlock()
	rList := make([]Role, len(roles))
	i := 0
	for _, r := range roles {
		rList[i] = r
		i++
	}
	return rList
}

// Add adds a Role registration
func Add(role Role) {
	roleLock.Lock()
	defer roleLock.Unlock()
	if _, ok := roles[role.ID()]; !ok {
		roles[role.ID()] = role
	}
}
