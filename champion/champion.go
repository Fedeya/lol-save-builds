package champion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
)

// Field is the type for the champion fields struct
type Field struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// Champion is a type for League of Legends Champeons Data
type Champion struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Line        string     `json:"line"`
	Skills      [3]Field   `json:"skills"`
	SkillsOrder [15]string `json:"skill_order"`
	Summoners   [4]Field   `json:"summoners"`
	Items       [][]Field  `json:"items"`
	Runes       [4]struct {
		Primary struct {
			Type Field    `json:"type"`
			Main Field    `json:"main"`
			More [3]Field `json:"more"`
		} `json:"primary"`
		Secondary struct {
			Type Field    `json:"type"`
			More [2]Field `json:"more"`
		} `json:"secondary"`
		Boosts [3]struct {
			Field
			Description string `json:"description"`
		} `json:"boosts"`
	} `json:"runes"`
}

// PrintData show the champion data in the console
func (c Champion) PrintData() {
	fmt.Printf("Name: %s \n", c.Name)
	fmt.Printf("Line: %s \n", c.Line)
	fmt.Printf("Skills: %s \n", c.Skills)
	fmt.Printf("Skills Order: %s \n", c.SkillsOrder)
	fmt.Printf("Summoners: %s \n", c.Summoners)
	fmt.Printf("\nItems:\n")
	for i, arr := range c.Items {
		fmt.Println("List -> ", i)
		for _, item := range arr {
			fmt.Printf("%s \n", item.Name)
		}
		fmt.Println()
	}
	fmt.Println("Runes:")
	for i, r := range c.Runes {
		fmt.Printf("----------------- [%d] ----------------- \n", i)
		fmt.Printf("Primary Type: %s \n", r.Primary.Type)
		fmt.Printf("Primary Main Rune: %s \n", r.Primary.Main)
		fmt.Printf("Primary More Runes: %s \n", r.Primary.More)

		fmt.Println()
		fmt.Printf("Secondary Type: %s \n", r.Secondary.Type)
		fmt.Printf("Secondary More Runes: %s \n", r.Secondary.More)

		fmt.Printf("Boosts: %s \n", r.Boosts)
		fmt.Printf("----------------- [%d] ----------------- \n", i)
		fmt.Println()
	}
}

// SaveJSONChamp Save the champion in the builds.json file
func (c Champion) SaveJSONChamp() {
	file := readFile()

	var champion []Champion
	id, _ := uuid.NewRandom()
	c.ID = id

	json.Unmarshal(file, &champion)
	champion = append(champion, c)
	text, _ := json.Marshal(champion)

	ioutil.WriteFile("builds.json", text, 0644)
}

// GetBuildsWithJSON transform champion builds.json and return champion array
func GetBuildsWithJSON() []Champion {
	file := readFile()
	var champions []Champion
	json.Unmarshal(file, &champions)

	return champions
}

func readFile() []byte {
	file, err := ioutil.ReadFile("builds.json")
	if err != nil {
		ioutil.WriteFile("builds.json", []byte("[]"), 0644)
	}

	return file
}
