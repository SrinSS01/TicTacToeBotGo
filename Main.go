package main

import (
	"TicTacToeBot/buttons"
	"TicTacToeBot/commands"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	config    = Config{}
	Logger    *zap.Logger
	discord   *discordgo.Session
	_commands = []*discordgo.ApplicationCommand{
		commands.Play.Command,
	}
	commandHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		commands.Play.Command.Name: commands.Play.Execute,
	}
	componentHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		buttons.Accept.Button.CustomID:  buttons.Accept.Execute,
		buttons.Decline.Button.CustomID: buttons.Decline.Execute,
		buttons.Stop.Button.CustomID:    buttons.Stop.Execute,
		buttons.Zero.Button.CustomID:    buttons.Zero.Execute,
		buttons.One.Button.CustomID:     buttons.One.Execute,
		buttons.Two.Button.CustomID:     buttons.Two.Execute,
		buttons.Three.Button.CustomID:   buttons.Three.Execute,
		buttons.Four.Button.CustomID:    buttons.Four.Execute,
		buttons.Five.Button.CustomID:    buttons.Five.Execute,
		buttons.Six.Button.CustomID:     buttons.Six.Execute,
		buttons.Seven.Button.CustomID:   buttons.Seven.Execute,
		buttons.Eight.Button.CustomID:   buttons.Eight.Execute,
	}
)

// args check
func init() {
	// check for TOKEN in environment variables
	if token := os.Getenv("TOKEN"); token != "" {
		config.Token = token
		return
	}
	if len(os.Args) != 2 {
		content, err := os.ReadFile("config.json")
		if err != nil {
			fmt.Print("Enter bot token: ")
			if _, err := fmt.Scanln(&config.Token); err != nil {
				log.Fatal("Error during Scanln(): ", err)
			}
			configJson()
			return
		}
		if err := json.Unmarshal(content, &config); err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		return
	}
	config.Token = os.Args[1]
	configJson()
}

func configJson() {
	marshal, err := json.Marshal(&config)
	if err != nil {
		log.Fatal("Error during Marshal(): ", err)
		return
	}
	if err := os.WriteFile("config.json", marshal, 0644); err != nil {
		log.Fatal("Error during WriteFile(): ", err)
	}
}

// init logger
func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// init discord
func init() {
	var err error
	discord, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		Logger.Fatal("Error creating Discord session", zap.Error(err))
		return
	}
	discord.Identify.Intents = discordgo.IntentsGuildMembers
}

// init discord handlers
func init() {
	discord.AddHandler(onReady)
	discord.AddHandler(slashCommandInteraction)
	discord.AddHandler(buttonInteraction)
}
func main() {
	if err := discord.Open(); err != nil {
		Logger.Fatal("Error opening connection", zap.Error(err))
		return
	}
	for _, command := range _commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", command)
		if err != nil {
			Logger.Fatal("Error creating slash command", zap.Error(err))
			return
		}
	}
	Logger.Info("Bot is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	if err := discord.Close(); err != nil {
		Logger.Fatal("Error closing connection", zap.Error(err))
		return
	}
	Logger.Info("Bot is shutting down")
}

func onReady(session *discordgo.Session, _ *discordgo.Ready) {
	Logger.Info(session.State.User.Username + " is ready")
}

func slashCommandInteraction(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}
	commandHandlers[interaction.ApplicationCommandData().Name](session, interaction)
}

func buttonInteraction(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionMessageComponent {
		return
	}
	componentHandlers[interaction.MessageComponentData().CustomID](session, interaction)
}
