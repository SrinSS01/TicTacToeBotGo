package buttons

import (
	"TicTacToeBot/game"
	"github.com/bwmarrin/discordgo"
	"log"
)

type DeclineButton struct {
	Button *discordgo.Button
}

var Decline = DeclineButton{
	Button: &discordgo.Button{
		Label:    "Decline",
		CustomID: "decline",
		Style:    discordgo.DangerButton,
	},
}

func (d *DeclineButton) Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	matchData := game.MatchDataMap[interaction.Message.ID]
	if matchData == nil {
		deferEdit(session, interaction)
		return
	}
	if interaction.Member.User.ID != matchData.Player.ID {
		deferEdit(session, interaction)
		return
	}
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    matchData.Player.Mention() + " declined!",
			Components: []discordgo.MessageComponent{},
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	game.MatchDataMap[interaction.Message.ID] = nil
}
