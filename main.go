package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		auth      = os.Getenv("DISCORD_AUTH")
		GuildID   = os.Getenv("GUILD_ID")
		ChannelID = os.Getenv("CHANNEL_ID")
	)

	discord, err := discordgo.New("Bot " + auth)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	dgv, err := discord.ChannelVoiceJoin(GuildID, ChannelID, false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/rob", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		go dgvoice.PlayAudioFile(dgv, "robs.mp3", make(chan bool))
	})

	http.HandleFunc("/brandon", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
		go dgvoice.PlayAudioFile(dgv, "chances.mp3", make(chan bool))
	})

	http.HandleFunc("/dan", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
		go dgvoice.PlayAudioFile(dgv, "odins.mp3", make(chan bool))
	})

	http.HandleFunc("/default", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
		go dgvoice.PlayAudioFile(dgv, "sads.mp3", make(chan bool))
	})

	defer dgv.Close()
	defer discord.Close()

	log.Fatal(http.ListenAndServe(":8081", nil))
	fmt.Println("hi")
}
