package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/queue"
)

func readCards(input string) (cards *queue.Queue) {
	cards = queue.New()
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
		cards.Enqueue(cardNum)
	}
	return
}

func playGame(cardsP1 *queue.Queue, cardsP2 *queue.Queue) (int, *queue.Queue, *queue.Queue) {
	for cardsP1.Len() > 0 && cardsP2.Len() > 0 {
		topCardP1 := cardsP1.Dequeue().(int)
		topCardP2 := cardsP2.Dequeue().(int)
		if topCardP1 > topCardP2 {
			cardsP1.Enqueue(topCardP1)
			cardsP1.Enqueue(topCardP2)
		} else {
			cardsP2.Enqueue(topCardP2)
			cardsP2.Enqueue(topCardP1)
		}
	}

	if cardsP1.Len() > 0 {
		return 0, cardsP1, cardsP2
	} else {
		return 1, cardsP1, cardsP2
	}
}

func scoreGame(cards *queue.Queue) int {
	value := cards.Len()
	score := 0
	for card := cards.Dequeue().(int); cards.Len() != 0; card = cards.Dequeue().(int) {
		score += card * value
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
