package cards

import (
	"math/rand"
	"time"
)

var r *rand.Rand

type Deck struct {
	Cards []*Card
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (d *Deck) Shuffle() {
	perm := rand.Perm(len(d.Cards))
	temp := make([]*Card, len(d.Cards))

	for i, v := range perm {
		temp[v] = d.Cards[i]
	}

	d.Cards = temp
}

func (d *Deck) Draw(n int) []*Card {
	cards := make([]*Card, 0)

	for i := 0; i < n; i++ {
		x := d.Cards[0]
		d.Cards = d.Cards[1:]
		cards = append(cards, x)
	}

	return cards
}


// GetNewDeck will return you a deck, right now it only gives you a deck for 500
// Will look into ways to pass in deck metadata
func GetNewDeck() *Deck {
	d := Deck{Cards: make([]*Card, 0)}

	cardValues := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for suit := 0; suit < 4; suit++ {
		for _, val := range cardValues {
			newCard := Card{Suit: suit, Value: val}
			d.Cards = append(d.Cards, &newCard)
		}
	}

	joker := Card{Suit: SuitNoTrump, Value: 15}
	d.Cards = append(d.Cards, &joker)

	return &d
}