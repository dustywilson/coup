package captain

import "github.com/dustywilson/coup/role"

// Captain is a Role
type Captain struct {
}

func init() {
	role.Add(Captain{})
}

// ID is a stable ID for this Role
func (c Captain) ID() string {
	return "CAPTAIN"
}

// Category is Finance
func (c Captain) Category() string {
	return role.SI
}

// Name is Captain
func (c Captain) Name() string {
	return "Captain"
}

// IsAdvanced is false
func (c Captain) IsAdvanced() bool {
	return false
}

// IsAction is true
func (c Captain) IsAction() bool {
	return true
}

// ProvidesExtraAction is false
func (c Captain) ProvidesExtraAction() bool {
	return false
}

// BlocksAnytime is false
func (c Captain) BlocksAnytime() []string {
	return []string{}
}

// BlocksIfTarget is false
func (c Captain) BlocksIfTarget() []string {
	return []string{"CAPTAIN"}
}

// BlockedByAnytime is false
func (c Captain) BlockedByAnytime() []string {
	return []string{}
}

// BlockedByIfTarget is false
func (c Captain) BlockedByIfTarget() []string {
	return []string{"AMBASSADOR", "CAPTAIN"}
}

// PayToBank is 0
func (c Captain) PayToBank() int {
	return 0
}

// EarnFromBank is 0
func (c Captain) EarnFromBank() int {
	return 0
}

// TakeFromTargetPlayer is 2
func (c Captain) TakeFromTargetPlayer() int {
	return 2
}

// TakeFromAllPlayers is 0
func (c Captain) TakeFromAllPlayers() int {
	return 0
}

// GiveToTargetPlayer is 0
func (c Captain) GiveToTargetPlayer() int {
	return 0
}

// GiveToAllPlayers is 0
func (c Captain) GiveToAllPlayers() int {
	return 0
}
