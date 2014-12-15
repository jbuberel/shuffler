package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Creating deck of cards\n")
	cards := make([]string, 52)
	fmt.Printf("Initializing values\n")
	
	for i := 0; i < len(cards); i++ {
		cards[i] = fmt.Sprintf("%v", i)
	}
	
	fmt.Printf("Cards is now:\n")
	for i, v := range cards {
		fmt.Printf("%v - %v\n", i, v)
	}
	
	fmt.Printf("About to shuffle\n")
	shuffle(cards)
	
	fmt.Printf("Cards is now\n")
	for i, v := range cards {
		fmt.Printf("%v - %v\n", i, v)
	}
	
	
}

func shuffle (cards []string) {
	
	r := rand.New(rand.NewSource(99))
	l := len(cards)
	for i := range cards {
		j := r.Intn(l)
		t := cards[i]
		cards[i] = cards[j]
		cards[j] = t
	}
	
}