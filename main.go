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

	app := wails.CreateApp(&wails.AppConfig{
		Width:     800,
		Height:    600,
		Title:     "LoL Save Builds",
		JS:        js,
		Resizable: true,
	})

	scrap("https://www.op.gg/champion/sett/statistics/top")

	app.Bind(scrap)
	app.Bind(getBuilds)

	// app.Run()
}
