package ambassador

import "github.com/dustywilson/coup/role"

// Ambassador is a Role
type Ambassador struct {
}

func init() {
	role.Add(Ambassador{})
}

// ID is a stable ID for this Role
func (a Ambassador) ID() string {
	return "AMBASSASDOR"
}

// Category is Finance
func (a Ambassador) Category() string {
	return role.Communications
}

// Name is Ambassador
func (a Ambassador) Name() string {
	return "Ambassador"
}

// IsAdvanced is false
func (a Ambassador) IsAdvanced() bool {
	return false
}

// IsAction is true
func (a Ambassador) IsAction() bool {
	return true
}

// ProvidesExtraAction is false
func (a Ambassador) ProvidesExtraAction() bool {
	return false
}

// BlocksAnytime is false
func (a Ambassador) BlocksAnytime() []string {
	return []string{}
}

// BlocksIfTarget is false
func (a Ambassador) BlocksIfTarget() []string {
	return []string{"CAPTAIN"}
}

// BlockedByAnytime is false
func (a Ambassador) BlockedByAnytime() []string {
	return []string{}
}

// BlockedByIfTarget is false
func (a Ambassador) BlockedByIfTarget() []string {
	return []string{}
}

// PayToBank is 0
func (a Ambassador) PayToBank() int {
	return 0
}

// EarnFromBank is 0
func (a Ambassador) EarnFromBank() int {
	return 0
}

// TakeFromTargetPlayer is 0
func (a Ambassador) TakeFromTargetPlayer() int {
	return 0
}

// TakeFromAllPlayers is 0
func (a Ambassador) TakeFromAllPlayers() int {
	return 0
}

// GiveToTargetPlayer is 0
func (a Ambassador) GiveToTargetPlayer() int {
	return 0
}

// GiveToAllPlayers is 0
func (a Ambassador) GiveToAllPlayers() int {
	return 0
}
