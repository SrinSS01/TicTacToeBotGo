package game

import (
	"github.com/bwmarrin/discordgo"
)

type Game struct {
	Cells      []discordgo.MessageComponent
	TicTacToe  *TicTacToe
	Challenger *discordgo.User
	Player     *discordgo.User
	playerMap  map[PlayerType]*discordgo.User
}

func (g *Game) GetCurrentPlayer() *discordgo.User {
	return g.playerMap[*g.TicTacToe.CurrentPlayer.Type]
}

func (g *Game) GetCurrentPlayerOpponent() *discordgo.User {
	return g.playerMap[*g.TicTacToe.CurrentPlayer.OpponentType]
}

func NewGame(challenger *discordgo.User, player *discordgo.User) *Game {
	return &Game{
		playerMap: map[PlayerType]*discordgo.User{
			CROSS:  challenger,
			NOUGHT: player,
		},
		Cells: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "8",
						Style:    discordgo.SecondaryButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "7",
						Style:    discordgo.SecondaryButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "6",
						Style:    discordgo.SecondaryButton,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "5",
						Style:    discordgo.SecondaryButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "4",
						Style:    discordgo.SecondaryButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "3",
						Style:    discordgo.SecondaryButton,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "2",
						Style:    discordgo.SecondaryButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "1",
						Style:    discordgo.SecondaryButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "0",
						Style:    discordgo.SecondaryButton,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "filler1",
						Style:    discordgo.SecondaryButton,
						Disabled: true,
					},
					discordgo.Button{
						Label:    "stop",
						CustomID: "stop",
						Style:    discordgo.DangerButton,
					},
					discordgo.Button{
						Label:    "\u0000",
						CustomID: "filler2",
						Style:    discordgo.SecondaryButton,
						Disabled: true,
					},
				},
			},
		},
		TicTacToe: NewTicTacToe(),
	}
}
