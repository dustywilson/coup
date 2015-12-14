package main

import (
	"fmt"

	"github.com/dustywilson/coup/role"
	_ "github.com/dustywilson/coup/role/ambassador"
	_ "github.com/dustywilson/coup/role/assassin"
	_ "github.com/dustywilson/coup/role/captain"
	_ "github.com/dustywilson/coup/role/contessa"
	_ "github.com/dustywilson/coup/role/duke"

	"github.com/dustywilson/coup/player"
	_ "github.com/dustywilson/coup/player/dummy"
)

func main() {
	fmt.Println("\n[Roles]")
	for _, r := range role.Roles() {
		fmt.Printf("\n\t%s\n", r.Name())
		fmt.Printf("\t\tID:\t\t\t%s\n", r.ID())
		fmt.Printf("\t\tCategory:\t\t%s\n", r.Category())
		if r.IsAdvanced() {
			fmt.Printf("\t\tIsAdvanced:\t\t%t\n", r.IsAdvanced())
		}
		if len(r.BlocksAnytime()) > 0 {
			fmt.Printf("\t\tBlocksAnytime:\t\t%+v\n", r.BlocksAnytime())
		}
		if len(r.BlocksIfTarget()) > 0 {
			fmt.Printf("\t\tBlocksIfTarget:\t\t%+v\n", r.BlocksIfTarget())
		}
		if len(r.BlockedByAnytime()) > 0 {
			fmt.Printf("\t\tBlockedByAnytime:\t%+v\n", r.BlockedByAnytime())
		}
		if len(r.BlockedByIfTarget()) > 0 {
			fmt.Printf("\t\tBlockedByIfTarget:\t%+v\n", r.BlockedByIfTarget())
		}
		if r.IsAction() {
			fmt.Printf("\t\tIsAction:\t\t%t\n", r.IsAction())
		}
		if r.ProvidesExtraAction() {
			fmt.Printf("\t\tProvidesExtraAction:\t%t\n", r.ProvidesExtraAction())
		}
		if r.PayToBank() != 0 {
			fmt.Printf("\t\tPayToBank:\t\t%d\n", r.PayToBank())
		}
		if r.EarnFromBank() != 0 {
			fmt.Printf("\t\tEarnFromBank:\t\t%d\n", r.EarnFromBank())
		}
		if r.TakeFromTargetPlayer() != 0 {
			fmt.Printf("\t\tTakeFromTargetPlayer:\t%d\n", r.TakeFromTargetPlayer())
		}
		if r.TakeFromAllPlayers() != 0 {
			fmt.Printf("\t\tTakeFromAllPlayers:\t%d\n", r.TakeFromAllPlayers())
		}
		if r.GiveToTargetPlayer() != 0 {
			fmt.Printf("\t\tGiveToTargetPlayer:\t%d\n", r.GiveToTargetPlayer())
		}
		if r.GiveToAllPlayers() != 0 {
			fmt.Printf("\t\tGiveToAllPlayers:\t%d\n", r.GiveToAllPlayers())
		}
	}
	fmt.Println("\n[Players]")
	for _, r := range player.Players() {
		fmt.Printf("\t%-20s\n", r.TypeName())
	}
	fmt.Println()
}
