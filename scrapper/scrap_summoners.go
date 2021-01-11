package scrapper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

// GetChampSummoners get the recomended champion summoners
func (s *Scrapper) GetChampSummoners(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()
	e.ForEachWithBreak("img.tip", func(i int, e *colly.HTMLElement) bool {
		title := e.Attr("title")
		r := strings.NewReader(title)
		doc, _ := goquery.NewDocumentFromReader(r)
		skill := doc.Find("b").Text()
		src := e.Attr("src")
		strings.TrimSpace(skill)
		if s.Champ.Summoners[1].Name != "" {
			if i == 0 {
				s.Champ.Summoners[2].Name = skill
        s.Champ.Summoners[2].ImageURL = "https:" + src
			} else if i == 1 {
				s.Champ.Summoners[3].Name = skill
        s.Champ.Summoners[3].ImageURL = "https:" + src
				return false
			}
		}
		s.Champ.Summoners[i].Name = skill
    s.Champ.Summoners[i].ImageURL = "https:" + src
		return true
	})
}
