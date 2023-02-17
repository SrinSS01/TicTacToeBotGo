package game

import "github.com/bwmarrin/discordgo"

type MatchData struct {
	Challenger discordgo.User
	Player     discordgo.User
	Accepted   bool
	Game       *Game
}

var MatchDataMap = map[string]*MatchData{}

func (m *MatchData) GetEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Game Started!",
			IconURL: m.Challenger.AvatarURL(""),
		},
		Description: m.Challenger.Mention() + "'s turn!",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://tic-tac-toe-next-app-three.vercel.app/api/og?" + "x=0&o=0",
		},
		Color: 0x2b2d31,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "vs " + m.Player.Username,
			IconURL: m.Player.AvatarURL(""),
		},
	}
}
