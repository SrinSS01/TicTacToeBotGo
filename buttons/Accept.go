package buttons

import (
	"TicTacToeBot/game"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

var Logger, _ = zap.NewProduction()

type AcceptButton struct {
	Button *discordgo.Button
}

var Accept = AcceptButton{
	Button: &discordgo.Button{
		Label:    "Accept",
		CustomID: "accept",
		Style:    discordgo.SuccessButton,
	},
}

func (a *AcceptButton) Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	matchData := game.MatchDataMap[interaction.Message.ID]
	if matchData == nil {
		deferEdit(session, interaction)
		return
	}
	if interaction.Member.User.ID != matchData.Player.ID {
		deferEdit(session, interaction)
		return
	}
	matchData.Accepted = true
	matchData.Game = game.NewGame(&matchData.Challenger, &matchData.Player)
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{matchData.GetEmbed()},
			Components: matchData.Game.Cells,
		},
	})
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}

func deferEdit(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}
