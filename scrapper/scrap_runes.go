package scrapper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

var runePageCount int = -1
var moreRunesCount int
var skipRune int = 0
var boostCount int = 0
var runePageBoostCount int = 0

// GetChampRunesType get the runes types
func (s *Scrapper) GetChampRunesType(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()
	first := true
	e.ForEach(".perk-page", func(i int, h *colly.HTMLElement) {
		img := h.DOM.ChildrenFiltered(".perk-page__row").First().ChildrenFiltered("div").Children()

		title, _ := img.Attr("title")
		src, _ := img.Attr("src")

		r := strings.NewReader(title)

		doc, _ := goquery.NewDocumentFromReader(r)

		if first {
			s.Champ.Runes[e.Index].Primary.Type.Name = doc.Find("b").Text()
      s.Champ.Runes[e.Index].Primary.Type.ImageURL = "https:" + src
			first = false
		} else {
			s.Champ.Runes[e.Index].Secondary.Type.Name = doc.Find("b").Text()
      s.Champ.Runes[e.Index].Secondary.Type.ImageURL = "https:" + src
		}
	})
}

// GetChampRunesMain get the main runes
func (s *Scrapper) GetChampRunesMain(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()

	page := e.DOM.ChildrenFiltered(".perk-page").First()
	img := page.Find(".perk-page__row:nth-child(2)").Find(".perk-page__item--active").Find("img")

	alt, _ := img.Attr("alt")
	src, _ := img.Attr("src")

	s.Champ.Runes[e.Index].Primary.Main.Name = alt
  s.Champ.Runes[e.Index].Primary.Main.ImageURL = "https:" + src

}

// GetChampMoreRunes get the more runes in the page
func (s *Scrapper) GetChampMoreRunes(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()

	if e.Index == skipRune {
		skipRune = skipRune + 6
		runePageCount++
		moreRunesCount = 0
		return
	}

	img := e.DOM.ChildrenFiltered("div").ChildrenFiltered("img")
	alt, _ := img.Attr("alt")
	src, _ := img.Attr("src")

	if moreRunesCount > 2 {
		s.Champ.Runes[runePageCount].Secondary.More[moreRunesCount-3].Name = alt
    s.Champ.Runes[runePageCount].Secondary.More[moreRunesCount-3].ImageURL = "https:" + src
		moreRunesCount++
		return
	}

	s.Champ.Runes[runePageCount].Primary.More[moreRunesCount].Name = alt
  s.Champ.Runes[runePageCount].Primary.More[moreRunesCount].ImageURL = "https:" + src
	moreRunesCount++
}

// GetChampBoostRunes get the boost runes in the page
func (s *Scrapper) GetChampBoostRunes(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()

	if boostCount == 3 {
		runePageBoostCount++
		boostCount = 0
	}

	e.ForEach("img.active", func(i int, h *colly.HTMLElement) {
		s.Champ.Runes[runePageBoostCount].Boosts[boostCount].Name = h.Attr("alt")
    s.Champ.Runes[runePageBoostCount].Boosts[boostCount].ImageURL = "https:" + h.Attr("src")
		t := h.Attr("title")
		r := strings.NewReader(t)
		doc, _ := goquery.NewDocumentFromReader(r)

		s.Champ.Runes[runePageBoostCount].Boosts[boostCount].Description = doc.Find("span").Text()

		boostCount++
	})
}
