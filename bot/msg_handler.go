package bot

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Run() {
	bot_token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("BOT_TOKEN not found")
	}

	discord, err := discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal(err)
	}

	// message handler
	discord.AddHandler(onRecvHandler)
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	err = discord.Open()
	if err != nil {
		log.Fatal("session failed to open")
	}
	defer discord.Close()

	log.Println("running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}

func onRecvHandler(discord *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID != discord.State.User.ID {
		log.Println(msg.Content)
		if strings.Contains(msg.Content, "ip") {
			res := fetchPublicIP()
			discord.ChannelMessageSend(msg.ChannelID, res)
		}
	}

}

func fetchPublicIP() string {
	res, _ := http.Get("https://api.ipify.org")
	ipByte, _ := ioutil.ReadAll(res.Body)
	ipStr := string(ipByte)

	return ipStr
}
