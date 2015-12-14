package contessa

import "github.com/dustywilson/coup/role"

// Contessa is a Role
type Contessa struct {
}

func init() {
	role.Add(Contessa{})
}

// ID is a stable ID for this Role
func (c Contessa) ID() string {
	return "CONTESSA"
}

// Category is Finance
func (c Contessa) Category() string {
	return role.SI
}

// Name is Contessa
func (c Contessa) Name() string {
	return "Contessa"
}

// IsAdvanced is false
func (c Contessa) IsAdvanced() bool {
	return false
}

// IsAction is false
func (c Contessa) IsAction() bool {
	return false
}

// ProvidesExtraAction is false
func (c Contessa) ProvidesExtraAction() bool {
	return false
}

// BlocksAnytime is false
func (c Contessa) BlocksAnytime() []string {
	return []string{}
}

// BlocksIfTarget is false
func (c Contessa) BlocksIfTarget() []string {
	return []string{"ASSASSIN"}
}

// BlockedByAnytime is false
func (c Contessa) BlockedByAnytime() []string {
	return []string{}
}

// BlockedByIfTarget is false
func (c Contessa) BlockedByIfTarget() []string {
	return []string{}
}

// PayToBank is 0
func (c Contessa) PayToBank() int {
	return 0
}

// EarnFromBank is 0
func (c Contessa) EarnFromBank() int {
	return 0
}

// TakeFromTargetPlayer is 0
func (c Contessa) TakeFromTargetPlayer() int {
	return 0
}

// TakeFromAllPlayers is 0
func (c Contessa) TakeFromAllPlayers() int {
	return 0
}

// GiveToTargetPlayer is 0
func (c Contessa) GiveToTargetPlayer() int {
	return 0
}

// GiveToAllPlayers is 0
func (c Contessa) GiveToAllPlayers() int {
	return 0
}
