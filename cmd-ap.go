package main

import(
        irc "github.com/fluffle/goirc/client"
	"github.com/kless/goconfig/config"
	"http"
	"strings"
)

const apConfigFile = "ap.conf"

func apUserExists(username string) (isuser bool) {
	url := "http://anime-planet.com/users/" + username

	r, _, err := http.Get(url)

	if err != nil {
		return false
	}

	if r.StatusCode == 200 {
		return true
	}

	r.Body.Close()

	return false
}

func apReadConfig(nick *irc.Nick) (apnick string) {
	c, _ := config.ReadDefault(apConfigFile)
	
	hostmask := user(nick)

	apnick, _ = c.String(hostmask, "nick")

	return apnick
}

func apProfile(conn *irc.Conn, nick *irc.Nick, arg string, channel string) {
	username := strings.TrimSpace(arg)

	if username  == "" {
		username = apReadConfig(nick)
	}

	if username == "" {
		username = nick.Nick
	}

	if apUserExists(username) {
		say(conn, channel, "%s's profile: http://anime-planet.com/users/%s", username, username)
		return
	}
	
	say(conn, channel, "The user '%s' doesn't exist. Try again.", username)
}

func apAnimeList(conn *irc.Conn, nick *irc.Nick, arg string, channel string) {
        username := strings.TrimSpace(arg)

        if username == "" {
                username = apReadConfig(nick)
        }

        if username == "" {
                username = nick.Nick
        }

	if apUserExists(username) {
		say(conn, channel, "%s's anime list: http://anime-planet.com/users/%s/anime", username, username)
		return
	}

	say(conn, channel, "The user '%s' doesn't exist. Try again.", username)
}

func apMangaList(conn *irc.Conn, nick *irc.Nick, arg string, channel string) {
        username := strings.TrimSpace(arg)

        if username == "" {
                username = apReadConfig(nick)
        }

        if username == "" {
                username = nick.Nick
        }

	if apUserExists(username) {
		say(conn, channel, "%s's manga list: http://anime-planet.com/users/%s/manga", username, username)
		return
	}

	say(conn, channel, "The user '%s' doesn't exist. Try again.", username)
}

func apSetNick(conn *irc.Conn, nick *irc.Nick, arg string, channel string) {
	arg = strings.TrimSpace(arg)

	if arg == "" {
		say(conn, channel, "Format is !apnick <nickname>")
		return
	}

	if apUserExists(arg) {
		c, _ := config.ReadDefault(apConfigFile)
	
		hostmask := user(nick)

		c.AddOption(hostmask, "nick", arg)
		c.WriteFile(apConfigFile, 0644, "")

		say(conn, channel, "Your anime-planet.com username has been recorded as '%s'", arg)
		return
	}

	say(conn, channel, "The user '%s' doesn't exist. Try again.", arg)
}

func apMyNick(conn *irc.Conn, nick *irc.Nick, arg string, channel string) {
	username := apReadConfig(nick)

	if username == "" {
		say(conn, channel, "You haven't set your username. You can do so with !apnick <username>")
		return
	}

	say(conn, channel, "Your anime-planet.com username has been recorded as '%s'.", username)
}