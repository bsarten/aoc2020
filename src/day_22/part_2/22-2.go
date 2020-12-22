package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type empty struct{}

func readCards(input string) (cards *list.List) {
	cards = list.New()
	first := true
	for _, card := range strings.Split(input, "\n") {
		if card == "" {
			continue
		}
		if first {
			first = false
			continue
		}
		cardNum, _ := strconv.Atoi(card)
		cards.PushBack(cardNum)
	}
	return
}

func moveCards(cardsWinner *list.List, cardsLoser *list.List) {
	cardsWinner.MoveToBack(cardsWinner.Front())
	cardsWinner.PushBack(cardsLoser.Front().Value.(int))
	cardsLoser.Remove(cardsLoser.Front())
}

func playGame(cardsP1 *list.List, cardsP2 *list.List) (int, *list.List, *list.List) {
	p1History := make(map[string]empty)
	p2History := make(map[string]empty)
	p1History[createHistoryEntry(cardsP1)] = empty{}
	p2History[createHistoryEntry(cardsP2)] = empty{}
	for cardsP1.Len() > 0 && cardsP2.Len() > 0 {
		currentP1 := cardsP1.Front()
		currentP2 := cardsP2.Front()
		if cardsP1.Front().Value.(int) <= cardsP1.Len()-1 && currentP2.Value.(int) <= cardsP2.Len()-1 {
			newDeckP1 := copyCards(currentP1.Next(), currentP1.Value.(int))
			newDeckP2 := copyCards(currentP2.Next(), currentP2.Value.(int))
			winner, _, _ := playGame(newDeckP1, newDeckP2)
			if winner == 0 {
				moveCards(cardsP1, cardsP2)
			} else {
				moveCards(cardsP2, cardsP1)
			}
		} else if currentP1.Value.(int) > currentP2.Value.(int) {
			moveCards(cardsP1, cardsP2)
		} else {
			moveCards(cardsP2, cardsP1)
		}
		p1he := createHistoryEntry(cardsP1)
		p2he := createHistoryEntry(cardsP2)
		if _, exists := p1History[p1he]; exists {
			if _, exists := p2History[p2he]; exists {
				return 0, cardsP1, cardsP2
			}
		}
		p2History[p2he] = empty{}
		p1History[p1he] = empty{}
	}

	if cardsP1.Len() > 0 {
		return 0, cardsP1, cardsP2
	} else {
		return 1, cardsP1, cardsP2
	}
}

func copyCards(from *list.Element, numCards int) (newDeck *list.List) {
	newDeck = list.New()
	for i := 0; i < numCards; i++ {
		newDeck.PushBack(from.Value.(int))
		from = from.Next()
	}
	return
}

func createHistoryEntry(cards *list.List) (entry string) {
	for i := cards.Front(); i != nil; i = i.Next() {
		entry += fmt.Sprintf("%d,", i.Value.(int))
	}
	return
}

func scoreGame(cards *list.List) int {
	value := cards.Len()
	score := 0
	for i := cards.Front(); i != nil; i = i.Next() {
		score += i.Value.(int) * value
		value--
	}

	return score
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	playerSplit := strings.Split(string(b), "\n\n")
	cardsP1 := readCards(playerSplit[0])
	cardsP2 := readCards(playerSplit[1])
	_, cardsP1, cardsP2 = playGame(cardsP1, cardsP2)
	var score int
	if cardsP1.Len() > 0 {
		score = scoreGame(cardsP1)
	} else {
		score = scoreGame(cardsP2)
	}
	fmt.Println(score)
}
