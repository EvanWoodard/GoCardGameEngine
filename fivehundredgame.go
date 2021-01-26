package cards
// // Old code to start with
// package fivehundredgame

// import (
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/websocket"
// )

// const (
// 	spades = iota
// 	clubs
// 	diamonds
// 	hearts
// 	none
// )

// // Card ...
// type Card struct {
// 	Suit  int `json:"suit"`
// 	Value int `json:"value"`
// }

// type deck struct {
// 	cards []*Card
// }

// type player struct {
// 	username string
// 	hand     []*Card
// 	score    int
// 	socket   *websocket.Conn
// }

// type fiveHundredTable struct {
// 	players []*player
// 	deck    deck
// 	blind   []*Card
// 	trump   int
// 	time    time.Time
// }

// // GameStatus ...
// type GameStatus struct {
// 	Hand  []*Card `json:"hand"`
// 	Trump int     `json:"trump"`
// 	Score int     `json:"score"`
// 	Inkle int     `json:"inkle"`
// }

// var gameRoom = make(map[int]*fiveHundredTable)

// // FiveHundredSetup sets up a web socket in an active game of 500
// func FiveHundredSetup(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	username, _ := vars["username"]
// 	gameIDparam, _ := vars["gameID"]

// 	gameID, _ := strconv.Atoi(gameIDparam)

// 	upgrader := websocket.Upgrader{}

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	newPlayer := player{username: username, hand: make([]*Card, 0), score: 0, socket: ws}

// 	desiredTable := gameRoom[gameID]

// 	if desiredTable == nil {
// 		table := fiveHundredTable{players: []*player{&newPlayer}, trump: -1, time: time.Now()}
// 		gameRoom[gameID] = &table
// 	} else {
// 		if len(desiredTable.players) < 4 {
// 			desiredTable.players = append(desiredTable.players, &newPlayer)

// 			if len(desiredTable.players) == 4 {
// 				desiredTable.StartGame()
// 			}
// 		}
// 	}

// 	// Make sure we close the connection when the function returns
// 	// defer ws.Close()
// }

// func (t *fiveHundredTable) StartGame() {
// 	newDeck := getNewDeck()
// 	newDeck.Shuffle()
// 	newDeck.Deal(t)

// 	t.UpdateStatus()
// }

// func (t *fiveHundredTable) UpdateStatus() {
// 	for _, player := range t.players {
// 		gs := GameStatus{Hand: player.hand, Trump: -1, Score: 0, Inkle: -1}
// 		player.socket.WriteJSON(gs)
// 	}
// }

// func (d *deck) Shuffle() {
// 	perm := rand.Perm(len(d.cards))
// 	temp := make([]*Card, len(d.cards))

// 	for i, v := range perm {
// 		temp[v] = d.cards[i]
// 	}

// 	d.cards = temp
// }

// func (d *deck) Deal(t *fiveHundredTable) {
// 	for _, player := range t.players {
// 		player.hand = append(player.hand, d.Draw(3)...)
// 	}

// 	t.blind = append(t.blind, d.Draw(3)...)

// 	for _, player := range t.players {
// 		player.hand = append(player.hand, d.Draw(4)...)
// 	}

// 	t.blind = append(t.blind, d.Draw(2)...)

// 	for _, player := range t.players {
// 		player.hand = append(player.hand, d.Draw(3)...)
// 	}
// }

// func (d *deck) Draw(n int) []*Card {
// 	cards := make([]*Card, 0)

// 	for i := 0; i < n; i++ {
// 		x := d.cards[0]
// 		d.cards = d.cards[1:]
// 		cards = append(cards, x)
// 	}

// 	return cards
// }

// func getNewDeck() *deck {
// 	deck := deck{cards: make([]*Card, 0)}

// 	cardValues := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

// 	for suit := 0; suit < 4; suit++ {
// 		for _, value := range cardValues {
// 			newCard := Card{Suit: suit, Value: value}
// 			deck.cards = append(deck.cards, &newCard)
// 		}
// 	}

// 	joker := Card{Suit: none, Value: 15}
// 	deck.cards = append(deck.cards, &joker)

// 	return &deck
// }
