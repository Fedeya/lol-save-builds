package main

import (
	"github.com/fedeya/lol-save-builds/champion"
	"github.com/fedeya/lol-save-builds/scrapper"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func scrap(url string) {
	s := scrapper.New(url)
	s.Scrap()
}

func getBuilds() []champion.Champion {
	return champion.GetBuildsWithJSON()
}

func main() {
	js := mewn.String("./client/build/static/js/main.js")
	css := mewn.String("./client/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  800,
		Height: 600,
		Title:  "LoL Save Builds",
		JS:     js,
		CSS:    css,
	})

	app.Bind(scrap)
	app.Bind(getBuilds)

	app.Run()
}
