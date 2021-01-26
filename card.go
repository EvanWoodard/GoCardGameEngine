package cards

// Card ...
type Card struct {
	Suit  int `json:"suit"`
	Value int `json:"value"`
}

// CompareCards will return the highest value card
func CompareCards(trump int, led int, cards []*Card) *Card {
	var c = cards[0]
	var cc = cards[1:]

	for _, card := range cc {
		// If the joker is in, it should always win
		if c.Suit == SuitNoTrump {
			return c
		}
		if card.Suit == SuitNoTrump {
			return card
		}

		// If there is no trump, the highest value card wins
		// If both cards are trump, the highest value card wins
		if c.Suit == trump && card.Suit == trump {
			if c.Value < card.Value {
				c = card
			}
			continue
		}

		// If only one card is trump, it wins
		if c.Suit != trump && card.Suit == trump {
			c = card
			continue
		}
		if c.Suit == trump && card.Suit != trump {
			continue
		}

		// Now that trump is done, we have to deal with the led suit
		// In this case, the led suit acts as trump
		// Here the only time the led suit matters is if only one card is of the led suit
		// If both or neither card are of the led suit, they just compare value
		if c.Suit != led && card.Suit == led {
			c = card
			continue
		}
		if c.Suit == led && card.Suit != led {
			continue
		}

		// If there is no trump, and both cards are either led or not, all that is needed is a straight comparison
		if c.Value < card.Value {
			c = card
		}
	}

	return c
}