package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Deck []int

var (
	numDecks int = 1000
	numCards int = 52
)

func main() {
	r := rand.New(rand.NewSource(99))
	decks := make([]Deck, numDecks)

	// generate 1000 decks, and shuffle them	
	for d := 0; d < len(decks); d++ {
	
		cards := make(Deck, numCards)
		
		// populate the deck with sequential entries
		for i := 0; i < numCards; i++ {
			cards[i] = i
		}
		
		nullShuffle(r, cards)
		
		// store each shuffled deck
		decks[d] = cards
		
	}
	
	fmt.Printf("Is null shuffle fair? %v\n", isFair(decks))
	
	// generate 1000 decks, and shuffle them	
	for d := 0; d < len(decks); d++ {
	
		cards := make(Deck, numCards)
		
		// populate the deck with sequential entries
		for i := 0; i < numCards; i++ {
			cards[i] = i
		}
		
		randomShuffle(r, cards)
		
		// store each shuffled deck
		decks[d] = cards
		
	}
	
	fmt.Printf("Is random shuffle fair? %v\n", isFair(decks))
}

func randomShuffle(r *rand.Rand, cards Deck) {
	
	for i := 0; i < numCards; i++ {
		j := r.Intn(numCards)
		t := cards[i]
		cards[i] = cards[j]
		cards[j] = t
	}
	
}

func nullShuffle(r *rand.Rand, cards Deck) {
	return
}

// For each position in the set of decks, this will compute the
// average value and the standard deviation. A shuffle is considered
// unfair if the average for any position is greater than half
// the standard deviation away from the nominal mean (numCards/2).
func isFair (decks []Deck) (bool) {
	
	for i := 0; i < numCards; i++ {
		sum := 0
		stdDev := 0.0
		for j := 0; j < numDecks; j++ {
			sum += decks[j][i]
			stdDev += (float64(numCards/2) - float64(decks[j][i])) * (float64(numCards/2) - float64(decks[j][i])) 
		}
		stdDev = math.Sqrt(stdDev/float64(numDecks))
		avg := float64(sum/numDecks)
		
		if math.Abs(avg - float64(numCards/2)) > stdDev/2 {
			return false
		}
	}
	
	return true
}