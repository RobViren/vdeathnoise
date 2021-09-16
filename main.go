package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func connect(auth, GuildID, ChannelID string) (session *discordgo.Session, voice *discordgo.VoiceConnection, err error) {
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
	return discord, dgv, nil
}

func playit(log_entry, path string) {
	discord, dgv, err := connect(os.Getenv("DISCORD_AUTH"), os.Getenv("GUILD_ID"), os.Getenv("CHANNEL_ID"))
	if err != nil {
		panic(err)
	}
	defer dgv.Close()
	defer discord.Close()
	if strings.Contains(log_entry, "Sbeve The Dim") {
		dgvoice.PlayAudioFile(dgv, path+"/audio/robs.mp3", make(chan bool))
	} else if strings.Contains(log_entry, "Dan") {
		dgvoice.PlayAudioFile(dgv, path+"/audio/odins.mp3", make(chan bool))
	} else if strings.Contains(log_entry, "Brandon") {
		dgvoice.PlayAudioFile(dgv, path+"/audio/chances.mp3", make(chan bool))
	} else {
		dgvoice.PlayAudioFile(dgv, path+"/audio/sads.mp3", make(chan bool))
	}
}

func main() {
	working_dir, err := os.Executable()
	working_dir = filepath.Dir(working_dir)
	fmt.Println(working_dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = godotenv.Load(working_dir + "/.env")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log_entry := string(body)
		parts := strings.Split(log_entry, ":")
		alive_time, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			panic(err)
		}
		fmt.Println(alive_time)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if alive_time > 30 {
			go playit(log_entry, working_dir)
		}
	})
	fmt.Println("Starting Sever")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
