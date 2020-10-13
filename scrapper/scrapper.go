package scrapper

import (
	"sync"

	"github.com/fedeya/lol-save-builds/champeon"
	"github.com/gocolly/colly/v2"
)

// Scrapper is a type for collect data
type Scrapper struct {
	url   string
	c     *colly.Collector
	Champ champeon.Champeon
	wg    sync.WaitGroup
}

// New Create a new Instance for Scrapper
func New(url string) *Scrapper {
	c := colly.NewCollector()
	champ := champeon.Champeon{}

	return &Scrapper{
		url:   url,
		c:     c,
		Champ: champ,
		wg:    sync.WaitGroup{},
	}
}

// Visit the champ url for collect data
func (s *Scrapper) Visit() {
	s.c.Visit(s.url)
}

// Listeners is all Cully Listeners
func (s *Scrapper) Listeners() {
	go s.c.OnHTML("h1.champion-stats-header-info__name", s.GetChampName)
	go s.c.OnHTML("span.champion-stats-header__position__role", s.GetChampLine)
	go s.c.OnHTML("ul.champion-stats__list", s.GetChampSkills)
	go s.c.OnHTML("table.champion-skill-build__table > tbody", s.GetChampSkillsOrder)

	s.wg.Wait()
}
