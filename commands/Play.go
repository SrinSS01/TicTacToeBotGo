package commands

import (
	"TicTacToeBot/buttons"
	"TicTacToeBot/game"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

type PlayCommand struct {
	Command *discordgo.ApplicationCommand
}

var Play = PlayCommand{
	Command: &discordgo.ApplicationCommand{
		Name:        "play",
		Description: "Play a game of Tic Tac Toe",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to play with",
			},
		},
	},
}

func (c *PlayCommand) Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var err error
	options := interaction.ApplicationCommandData().Options
	if len(options) == 0 {
		playWithBot(session, interaction)
		return
	}
	user := options[0].UserValue(session)
	if user.Bot {
		playWithBot(session, interaction)
		return
	}
	member := interaction.Member.User
	if user.ID == member.ID {
		err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You can't play with yourself!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			return
		}
		return
	}
	err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: member.Mention() + " is challenging " + user.Mention() + " to a game of Tic Tac Toe!",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						buttons.Accept.Button,
						buttons.Decline.Button,
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	response, err := session.InteractionResponse(interaction.Interaction)
	if err != nil {
		return
	}
	game.MatchDataMap[response.ID] = &game.MatchData{
		Challenger: *member,
		Player:     *user,
		Accepted:   false,
	}
	time.Sleep(10 * time.Minute)
	matchData := game.MatchDataMap[response.ID]
	if matchData == nil || matchData.Accepted {
		return
	}
	content := user.Mention() + " didn't accept the challenge!"
	_, err = session.InteractionResponseEdit(
		interaction.Interaction,
		&discordgo.WebhookEdit{
			Content:    &content,
			Components: &[]discordgo.MessageComponent{},
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	game.MatchDataMap[response.ID] = nil
}

func playWithBot(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Not implemented yet!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		return
	}
}
