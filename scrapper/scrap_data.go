package scrapper

import (
	"github.com/gocolly/colly/v2"
)

// GetChampName get the champion name
func (s *Scrapper) GetChampName(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()
	s.Champ.Name = e.Text
}

// GetChampLine get the champion build line
func (s *Scrapper) GetChampLine(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()
	s.Champ.Line = e.Attr("data-position")
}

// GetChampImg get the champion build line
func (s *Scrapper) GetChampImg(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()
	s.Champ.ImageURL = e.ChildAttr("img", "src")
}
