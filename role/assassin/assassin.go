package assassin

import "github.com/dustywilson/coup/role"

// Assassin is a Role
type Assassin struct {
}

func init() {
	role.Add(Assassin{})
}

// ID is a stable ID for this Role
func (a Assassin) ID() string {
	return "ASSASSIN"
}

// Category is Finance
func (a Assassin) Category() string {
	return role.Force
}

// Name is Assassin
func (a Assassin) Name() string {
	return "Assassin"
}

// IsAdvanced is false
func (a Assassin) IsAdvanced() bool {
	return false
}

// IsAction is true
func (a Assassin) IsAction() bool {
	return true
}

// ProvidesExtraAction is false
func (a Assassin) ProvidesExtraAction() bool {
	return false
}

// BlocksAnytime is false
func (a Assassin) BlocksAnytime() []string {
	return []string{}
}

// BlocksIfTarget is false
func (a Assassin) BlocksIfTarget() []string {
	return []string{}
}

// BlockedByAnytime is false
func (a Assassin) BlockedByAnytime() []string {
	return []string{}
}

// BlockedByIfTarget is false
func (a Assassin) BlockedByIfTarget() []string {
	return []string{"CONTESSA"}
}

// PayToBank is 3
func (a Assassin) PayToBank() int {
	return 3
}

// EarnFromBank is 0
func (a Assassin) EarnFromBank() int {
	return 0
}

// TakeFromTargetPlayer is 0
func (a Assassin) TakeFromTargetPlayer() int {
	return 0
}

// TakeFromAllPlayers is 0
func (a Assassin) TakeFromAllPlayers() int {
	return 0
}

// GiveToTargetPlayer is 0
func (a Assassin) GiveToTargetPlayer() int {
	return 0
}

// GiveToAllPlayers is 0
func (a Assassin) GiveToAllPlayers() int {
	return 0
}
