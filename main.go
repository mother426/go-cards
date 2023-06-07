package main

import "fmt"

func main() {
	cards := []string{newCard(), "ace of spades"}
	cards = append(cards, "appended card")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "five of hearts"
}
