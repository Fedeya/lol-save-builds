package champeon

import "fmt"

// Champeon is a type for League of Legends Champeons Data
type Champeon struct {
	Name        string
	Line        string
	Skills      [3]string
	SkillsOrder [15]string
}

// PrintData show the champion data in the console
func (c Champeon) PrintData() {
	fmt.Printf("Name: %s \n", c.Name)
	fmt.Printf("Line: %s \n", c.Line)
	fmt.Printf("Skills: %s \n", c.Skills)
	fmt.Printf("Skills Order: %s \n", c.SkillsOrder)
}
