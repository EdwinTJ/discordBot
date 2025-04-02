package bot

import (
	"discordbot/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start(){
	
	goBot, err := discordgo.New("Bot " + config.Token)
    goBot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	u ,err := goBot.User("@me")
	if err != nil {
		fmt.Println("Error retrieving account,", err)
		return
	}
	BotID = u.ID
	channelID := config.ChannelID
	goBot.AddHandler(messageHandler)


	err = goBot.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	_, err = goBot.ChannelMessageSend(channelID, "Hola Panita")
	if err != nil {
		fmt.Println("Error sending initial message,", err)
		}

	


	fmt.Println("Bot is now running. Press CTRL+C to exit.")
} 

// Function to handle messages
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	const prefix string = "!bot"
	args := strings.Split(m.Content, " ")

	if args[0] != prefix {
		fmt.Println("No prefix")
		return
	}

	if args[1] == "jugar" {
		s.ChannelMessageSend(m.ChannelID, "Conectate R6!!!!!!")
		return
	}
}


// Function to get all messages
func getAllChannelMessages(s *discordgo.Session, channelID string) ([]*discordgo.Message, error) {
	var allMessages []*discordgo.Message
	var lastMessageID string
	
	// Discord limit = 100 msg per request
	limit := 25
	
	for {
		messages, err := s.ChannelMessages(channelID, limit, lastMessageID, "", "")
		if err != nil {
			return nil, fmt.Errorf("failed to get messages: %v", err)
		}
		
		if len(messages) == 0 {
			break
		}
		
		allMessages = append(allMessages, messages...)
		
		lastMessageID = messages[len(messages)-1].ID
		
		if len(messages) < limit {
			break
		}
	}
	
	return allMessages, nil
}