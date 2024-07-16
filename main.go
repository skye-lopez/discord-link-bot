package main

import (
	"fmt"
	"net"
	nu "net/url"
	"os"
	"os/signal"
	"strings"
	"database/sql"

	_ "github.com/lib/pq"
	migrator "github.com/skye-lopez/go-pq-migrator"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/skye-lopez/go-query"
)

var gq goquery.GoQuery

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    botToken := os.Getenv("BOT_TOKEN")
    gq = GetGoQuery() 
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

func storeUserLink(guildID string, channelID string, userID string, url string)  {
    // TODO: store it in a smart way
    linkID := guildID + channelID + userID
    //linkID, userID, channelID, guildID, link 
    _, err := gq.QueryString("SELECT insert_link($1, $2, $3, $4, $5)", linkID, userID, channelID, guildID, url)
    if err != nil {
        fmt.Println(err)
    }
}

func GetDB() *sql.DB {
    // TODO: This is just a testing db, eventually will need a real one.
    connStr := "postgres://me:me@localhost/tft_tracker"
    db, connErr := sql.Open("postgres", connStr)
    if connErr != nil {
        fmt.Println("Error opening PG", connErr)
    }
    return db
}

func GetGoQuery() goquery.GoQuery {
    db := GetDB()
    gq := goquery.NewGoQuery(db)
    gq.AddQueriesToMap("q")
    return gq
}

func GetMigrator() migrator.Migrator {
    db := GetDB()
    m, err := migrator.NewMigrator(db)
    if err != nil {
        //TODO: Err
    }
    m.AddQueriesToMap("migrations")
    return m
}
