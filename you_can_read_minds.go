package main

import (
	"fmt"
	"math/rand"
	"time"
)

type suit int

const (
	heart suit = iota
	diamond
	club
	spade
)

type card struct {
	suit   suit
	number int
}

func newRandCard() card {
	n := rand.Intn(52)
	quotient := n / 13
	remainder := n % 13
	return card{
		suit:   suit(quotient),
		number: remainder + 1,
	}
}

// less returns true if c should proceed d.
func (c card) less(d card) bool {
	if c.number != d.number {
		return c.number < d.number
	} else {
		return c.suit < d.suit
	}
}

// distance returns clock-wise distance from c to d.
func (c card) distance(d card) int {
	dist := d.number - c.number
	if dist < 0 {
		dist += 13
	}
	return dist
}

func (c card) String() string {
	return fmt.Sprintf("(%v_%v)", c.suit, c.number)
}

// getFiveRandomCards returns a slice of five unique cards.
func getFiveRandomCards() []card {
	cardSet := make(map[card]bool)
	for len(cardSet) != 5 {
		cardSet[newRandCard()] = true
	}
	var cards []card
	for c := range cardSet {
		cards = append(cards, c)
	}
	return cards
}

func encode(cards []card) (card, card, card, card) {
	// Pick the suit.
	count := make(map[suit]int)
	for _, c := range cards {
		count[c.suit]++
	}
	var targetSuit suit
	for s, c := range count {
		if c >= 2 {
			targetSuit = s
			break
		}
	}

	// Separate target card and other cards.
	var candidate1Flag, candidate2Flag bool
	var candidate1, candidate2 card
	var others []card
	for _, c := range cards {
		if c.suit == targetSuit {
			if !candidate1Flag {
				candidate1Flag = true
				candidate1 = c
				continue
			} else if !candidate2Flag {
				candidate2Flag = true
				candidate2 = c
				continue
			}
		}
		others = append(others, c)
	}

	var first card
	var dist int
	if candidate1.distance(candidate2) <= 6 {
		dist = candidate1.distance(candidate2)
		first = candidate1
	} else {
		dist = candidate2.distance(candidate1)
		first = candidate2
	}
	second, third, forth := encodeNumber(dist, others[0], others[1], others[2])
	return first, second, third, forth
}

func encodeNumber(n int, c1, c2, c3 card) (card, card, card) {
	var s, m, l card
	switch {
	case c1.less(c2) && c2.less(c3):
		s, m, l = c1, c2, c3
	case c1.less(c3) && c3.less(c2):
		s, m, l = c1, c3, c2
	case c2.less(c1) && c1.less(c3):
		s, m, l = c2, c1, c3
	case c2.less(c3) && c3.less(c1):
		s, m, l = c2, c3, c1
	case c3.less(c1) && c1.less(c2):
		s, m, l = c3, c1, c2
	case c3.less(c2) && c2.less(c1):
		s, m, l = c3, c2, c1
	}

	switch n {
	case 1:
		return s, m, l
	case 2:
		return s, l, m
	case 3:
		return m, s, l
	case 4:
		return m, l, s
	case 5:
		return l, s, m
	case 6:
		return l, m, s
	default:
		return s, m, l
	}
}

// decode decodes the fifth card based on four cards.
func decode(c1, c2, c3, c4 card) card {
	var result card
	result.suit = c1.suit

	n := decodeNumber(c2, c3, c4)
	n += c1.number
	if n > 13 {
		n -= 13
	}
	result.number = n

	return result
}

// decodeNumber decodes a number from 3 cards.
func decodeNumber(c1, c2, c3 card) int {
	switch {
	case c1.less(c2) && c2.less(c3):
		return 1
	case c1.less(c3) && c3.less(c2):
		return 2
	case c2.less(c1) && c1.less(c3):
		return 3
	case c3.less(c1) && c1.less(c2):
		return 4
	case c2.less(c3) && c3.less(c1):
		return 5
	case c3.less(c2) && c2.less(c1):
		return 6
	default:
		return 0
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := getFiveRandomCards()
	fmt.Println(cards)
	c1, c2, c3, c4 := encode(cards)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(decode(c1, c2, c3, c4))
}
