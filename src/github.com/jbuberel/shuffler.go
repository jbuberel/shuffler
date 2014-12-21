package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Deck []int

var (
	numDecks int = 1000
	numCards int = 52
	trials int = 100
	r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

)

func main() {

	var fairCount int

	fairCount = 0
	for i := 0; i < trials; i++ {
		if isFair(randomShuffle) {
			fairCount++
		}
	}
	fmt.Printf("randomShuffler %v/%v with %v decks\n", fairCount, trials, numDecks)
	
	fairCount = 0
	for i := 0; i < trials; i++ {
		if isFair(halfAssedShuffler) {
			fairCount++
		}
	}
	fmt.Printf("halfAssedShuffler %v/%v with %v decks\n", fairCount, trials, numDecks)
	
	
	fairCount = 0
	for i := 0; i < trials; i++ {
		if isFair(nullShuffle) {
			fairCount++
		}
	}
	fmt.Printf("nullShuffle %v/%v with %v decks\n", fairCount, trials, numDecks)
	
}

func randomShuffle(cards Deck) {

	for i := 0; i < numCards; i++ {
		j := r.Intn(numCards)
		t := cards[i]
		cards[i] = cards[j]
		cards[j] = t
	}

}

func nullShuffle(cards Deck) {
	return
}

func halfAssedShuffler(cards Deck) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numCards; i += 2 {
		j := r.Intn(numCards)
		t := cards[i]
		cards[i] = cards[j]
		cards[j] = t
	}
	
}

type shuffle func( cards Deck)



// For each position in the set of decks, this will compute the
// average value and the standard deviation. A shuffle is considered
// unfair if the average for any position is greater than half
// the standard deviation away from the nominal mean (numCards/2).
//func isFair (decks []Deck) (bool) {
func isFair(shuf shuffle) bool {
	decks := make([]Deck, numDecks)

	// generate 1000 decks, and shuffle them
	for d := 0; d < len(decks); d++ {

		cards := make(Deck, numCards)

		// populate the deck with sequential entries
		for i := 0; i < numCards; i++ {
			cards[i] = i
		}

		shuf(cards)

		// store each shuffled deck
		decks[d] = cards

	}

	for i := 0; i < numCards; i++ {
		sum := 0
		stdDev := 0.0
		for j := 0; j < numDecks; j++ {
			sum += decks[j][i]
			stdDev += (float64(numCards/2) - float64(decks[j][i])) * (float64(numCards/2) - float64(decks[j][i]))
		}
		stdDev = math.Sqrt(stdDev / float64(numDecks))
		avg := float64(sum / numDecks)

		if math.Abs(avg-float64(numCards/2)) > stdDev/2 {
			return false
		}
	}

	return true
}
