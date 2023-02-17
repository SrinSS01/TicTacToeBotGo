package buttons

import (
	"TicTacToeBot/game"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
)

type Cell struct {
	Button discordgo.Button
}

var Eight = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "8",
		Style:    discordgo.SecondaryButton,
	},
}

var Seven = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "7",
		Style:    discordgo.SecondaryButton,
	},
}

var Six = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "6",
		Style:    discordgo.SecondaryButton,
	},
}

var Five = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "5",
		Style:    discordgo.SecondaryButton,
	},
}

var Four = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "4",
		Style:    discordgo.SecondaryButton,
	},
}

var Three = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "3",
		Style:    discordgo.SecondaryButton,
	},
}

var Two = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "2",
		Style:    discordgo.SecondaryButton,
	},
}

var One = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "1",
		Style:    discordgo.SecondaryButton,
	},
}

var Zero = Cell{
	Button: discordgo.Button{
		Label:    "\u0000",
		CustomID: "0",
		Style:    discordgo.SecondaryButton,
	},
}

func (c *Cell) Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	matchData := game.MatchDataMap[interaction.Message.ID]
	if matchData == nil {
		deferEdit(session, interaction)
		return
	}
	currentPlayer := matchData.Game.GetCurrentPlayer()
	opponent := matchData.Game.GetCurrentPlayerOpponent()
	if interaction.Member.User.ID != currentPlayer.ID {
		deferEdit(session, interaction)
		return
	}
	cellIndex := c.Button.CustomID
	cellIndexInt, _ := strconv.Atoi(cellIndex)
	cellIndexInt = 8 - cellIndexInt
	row := cellIndexInt / 3
	column := cellIndexInt % 3
	button := c.Button
	button.Label = matchData.Game.TicTacToe.CurrentPlayer.Type.Value
	button.Style = discordgo.SuccessButton
	button.Disabled = true
	matchData.Game.Cells[row].(discordgo.ActionsRow).Components[column] = button
	result, indices := matchData.Game.TicTacToe.Place(8 - cellIndexInt)
	xBoard := matchData.Game.TicTacToe.GetXBoard()
	oBoard := matchData.Game.TicTacToe.GetOBoard()
	interaction.Message.Embeds[0].Image = &discordgo.MessageEmbedImage{
		URL: "https://tic-tac-toe-next-app-three.vercel.app/api/og?" + "x=" + xBoard + "&o=" + oBoard + "&w=" + strconv.Itoa(indices),
	}
	switch result {
	case game.WIN:
		interaction.Message.Embeds[0].Description = currentPlayer.Mention() + " won"
	case game.DRAW:
		interaction.Message.Embeds[0].Description = "It's a draw!"
	case game.NONE:
		interaction.Message.Embeds[0].Description = opponent.Mention() + "'s turn!"
	}
	cellMapOrNot := func() []discordgo.MessageComponent {
		if result == game.NONE {
			return matchData.Game.Cells
		} else {
			return nil
		}
	}
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     interaction.Message.Embeds,
			Components: cellMapOrNot(),
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
