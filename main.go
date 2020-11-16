package main

import (
	"fmt"

	"github.com/fedeya/lol-save-builds/scrapper"
)

func main() {
	url := "https://las.op.gg/champion/sett/statistics/top"

	if url == "" {
		fmt.Print("Enter a url: ")
		fmt.Scanln(&url)
	}

	s := scrapper.New(url)
	s.Scrap()

}
