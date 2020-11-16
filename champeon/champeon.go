package champeon

import "fmt"

// Item is a type for Items in the Champion
type Item struct {
	Name     string
	ImageURL string
}

type Field struct {
	Name     string
	ImageURL string
}

// Champeon is a type for League of Legends Champeons Data
type Champeon struct {
	Name        string
	Line        string
	Skills      [3]Field
	SkillsOrder [15]string
	Summoners   [4]Field
	Items       [][]Item
	Runes       [4]struct {
		Primary struct {
			Type Field
			Main Field
			More [3]Field
		}
		Secondary struct {
			Type Field
			More [2]Field
		}
		Boosts [3]struct {
			Field
			Description string
		}
	}
}

// PrintData show the champion data in the console
func (c Champeon) PrintData() {
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
