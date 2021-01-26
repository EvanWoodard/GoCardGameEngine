package cards

import ( 
	"testing"
)

// Testing to make sure that all cards, leads and trumps are supported properly
func TestCompareCards(t *testing.T) {
	aceOfHearts := Card{Suit: SuitHearts, Value: 14}
	queenOfSpades := Card{Suit: SuitSpades, Value: 12}
	queenOfDiamonds := Card{Suit: SuitDiamonds, Value: 12}
	fourOfClubs := Card{Suit: SuitClubs, Value: 4}

	testArr := []*Card{&aceOfHearts, &queenOfDiamonds, &queenOfSpades, &fourOfClubs}

	// If the trump and led suit is hearts, the aceOfHearts should win
	got := CompareCards(SuitHearts, SuitHearts, testArr)
	if got.Suit != SuitHearts {
		t.Errorf("CompareCards with trump/led hearts should have the Ace of Hearts win, instead got the %v of %v", got.Value, got.Suit)
	}

	// If only trump is hearts, the aceOfHearts should still win
	got = CompareCards(SuitHearts, SuitClubs, testArr)
	if got.Suit != SuitHearts {
		t.Errorf("CompareCards with trump hearts should have the Ace of Hearts win, instead got the %v of %v", got.Value, got.Suit)
	}

	// If trump is clubs, the little four should win
	got = CompareCards(SuitClubs, SuitHearts, testArr)
	if got.Suit != SuitClubs {
		t.Errorf("CompareCards with trump of clubs should have the four of clubs win, instead got %v of %v", got.Value, got.Suit)
	}

	// If there is no trump, the leading suit should win
	got = CompareCards(SuitNoTrump, SuitDiamonds, testArr)
	if got.Suit != SuitDiamonds {
		t.Errorf("CompareCards with noTrump and a lead of diamonds should have the queen of diamonds win, instead got %v of %v", got.Value, got.Suit)
	}

	fourOfHearts := Card{Suit: SuitHearts, Value: 4}
	sevenOfHearts := Card{Suit: SuitHearts, Value: 7}
	testArr2 := []*Card{&fourOfHearts, &aceOfHearts, &sevenOfHearts, &fourOfClubs}

	// When the trump or lead is hearts, the ace should always win
	got = CompareCards(SuitHearts, SuitDiamonds, testArr2)
	if got.Suit != SuitHearts || got.Value != 14 {
		t.Errorf("CompareCards with trump of hearts should have the highest heart win, instead got %v of %v", got.Value, got.Suit)
	}
	got = CompareCards(SuitHearts, SuitHearts, testArr2) 
	if got.Suit != SuitHearts || got.Value != 14 {
		t.Errorf("CompareCards with trump and lead of hearts should have the highest heart win, instead got %v of %v", got.Value, got.Suit)
	}
	got = CompareCards(SuitNoTrump, SuitHearts, testArr2)
	if got.Suit != SuitHearts || got.Value != 14 {
		t.Errorf("CompareCards with no trump and lead of hearts should have the highest heart win, instead got %v of %v", got.Value, got.Suit)
	}

	// The Joker should win in any scenario
	joker := Card{Suit: SuitNoTrump, Value: 15}
	
	testArr3 := []*Card{&fourOfClubs, &queenOfDiamonds, &joker, &aceOfHearts}

	got = CompareCards(SuitHearts, SuitHearts, testArr3)
	if got.Suit != SuitNoTrump {
		t.Errorf("CompareCards should always let the joker win, instead got %v of %v", got.Value, got.Suit)
	}
}