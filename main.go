package main

import (
	"fmt"
	"os"
	"os/signal"
    nu "net/url"
    "net"
    "strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    botToken := os.Getenv("BOT_TOKEN")
    RunBot(botToken)
}

func RunBot(token string) {
    discord, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Println("Could not connect to discord bot", err)
        return
    }
    
    discord.AddHandler(newMessage)

    // Open
    discord.Open()
    defer discord.Close()

    // Run until we get an interrupt.
    fmt.Println("Bot is running")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
    // Ignore messages from self
    if message.Author.ID == discord.State.User.ID {
        return
    }

    // Confirm the latest message is infact a URL
    url, messageIsUrl := confirmURL(message.Content)
    if messageIsUrl {
        storeUserLink(message.GuildID, message.ChannelID, message.Author.ID, url)
    }
}

func confirmURL(provided string) (string, bool) {
    u, err := nu.ParseRequestURI(provided)
    if err != nil {
        return "", false
    }

    adress := net.ParseIP(u.Host)
    if adress == nil {
        valid := strings.Contains(u.Host, ".")
        return provided, valid
    }

    return provided, true
}

func storeUserLink(guildID string, channelID string, userID string, url string) {
    // TODO: store it in a smart way
    fmt.Println(guildID, channelID, userID, url)
}
