package duke

import "github.com/dustywilson/coup/role"

// Duke is a Role
type Duke struct {
}

func init() {
	role.Add(Duke{})
}

// ID is a stable ID for this Role
func (d Duke) ID() string {
	return "DUKE"
}

// Category is Finance
func (d Duke) Category() string {
	return role.Finance
}

// Name is Duke
func (d Duke) Name() string {
	return "Duke"
}

// IsAdvanced is false
func (d Duke) IsAdvanced() bool {
	return false
}

// IsAction is true
func (d Duke) IsAction() bool {
	return true
}

// ProvidesExtraAction is false
func (d Duke) ProvidesExtraAction() bool {
	return false
}

// BlocksAnytime is false
func (d Duke) BlocksAnytime() []string {
	return []string{}
}

// BlocksIfTarget is false
func (d Duke) BlocksIfTarget() []string {
	return []string{}
}

// BlockedByAnytime is false
func (d Duke) BlockedByAnytime() []string {
	return []string{}
}

// BlockedByIfTarget is false
func (d Duke) BlockedByIfTarget() []string {
	return []string{}
}

// PayToBank is 0
func (d Duke) PayToBank() int {
	return 0
}

// EarnFromBank is 3
func (d Duke) EarnFromBank() int {
	return 3
}

// TakeFromTargetPlayer is 0
func (d Duke) TakeFromTargetPlayer() int {
	return 0
}

// TakeFromAllPlayers is 0
func (d Duke) TakeFromAllPlayers() int {
	return 0
}

// GiveToTargetPlayer is 0
func (d Duke) GiveToTargetPlayer() int {
	return 0
}

// GiveToAllPlayers is 0
func (d Duke) GiveToAllPlayers() int {
	return 0
}
