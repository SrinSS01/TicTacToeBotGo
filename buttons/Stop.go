package buttons

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type StopButton struct {
	Button *discordgo.Button
}

var Stop = StopButton{
	Button: &discordgo.Button{
		Label:    "stop",
		CustomID: "stop",
		Style:    discordgo.DangerButton,
	},
}

func (s *StopButton) Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	interaction.Message.Embeds[0].Description = interaction.Member.User.Mention() + " resigned!"
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: interaction.Message.Embeds,
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
