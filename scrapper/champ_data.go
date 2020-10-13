package scrapper

import (
	"github.com/gocolly/colly/v2"
)

// GetChampName get the champion name
func (s *Scrapper) GetChampName(e *colly.HTMLElement) {
	s.Champ.Name = e.Text
}

// GetChampLine get the champion build line
func (s *Scrapper) GetChampLine(e *colly.HTMLElement) {
	s.Champ.Line = e.Text
}
