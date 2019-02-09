package hand

import (
	"math"

	"github.com/enceeobee/crib/suit"
)

// Count returns the number of points in a hand
func (h *Hand) Count() int {
	return fifteensCount(h) +
		pairsCount(h) +
		runsCount(h) +
		flushCount(h) +
		nobsCount(h)
}

func fifteensCount(h *Hand) int {
	points := 0
	allCardsCount := 0

	for a := 0; a < 5; a++ {
		allCardsCount += h.Cards[a].Value

		for b := a + 1; b < 5; b++ {
			if h.Cards[a].Value+h.Cards[b].Value == 15 {
				points += 2
			}

			for c := b + 1; c < 5; c++ {
				if h.Cards[a].Value+h.Cards[b].Value+h.Cards[c].Value == 15 {
					points += 2
				}

				for d := c + 1; d < 5; d++ {
					if h.Cards[a].Value+h.Cards[b].Value+h.Cards[c].Value+h.Cards[d].Value == 15 {
						points += 2
					}
				}
			}
		}
	}

	if allCardsCount == 15 {
		points += 2
	}

	return points
}

func pairsCount(h *Hand) int {
	points := 0

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if h.Cards[i].SortOrder == h.Cards[j].SortOrder {
				points += 2
			}
		}
	}

	return points
}

func runsCount(h *Hand) int {
	runLen := 1
	multiplier := 0
	duplicateNumbers := make(map[int]int)

	h.Sort()

	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i].SortOrder == h.Cards[i+1].SortOrder-1 {
			runLen++
		} else if h.Cards[i].SortOrder == h.Cards[i+1].SortOrder {
			if _, ok := duplicateNumbers[h.Cards[i].SortOrder]; !ok {
				duplicateNumbers[h.Cards[i].SortOrder] = 1
			}
			duplicateNumbers[h.Cards[i].SortOrder]++
		} else if runLen < 3 {
			runLen = 1
		}
	}

	if runLen < 3 {
		return 0
	} else if len(duplicateNumbers) == 0 {
		return runLen
	}

	for _, val := range duplicateNumbers {
		multiplier += val
	}

	return runLen * multiplier
}

func flushCount(h *Hand) int {
	var clubCount,
		diamondCount,
		heartCount,
		spadeCount,
		highestQuantityOfSuitsCount float64

	for _, card := range h.Cards {
		switch card.Suit {
		case suit.Club:
			clubCount++
			highestQuantityOfSuitsCount = math.Max(highestQuantityOfSuitsCount, clubCount)
		case suit.Diamond:
			diamondCount++
			highestQuantityOfSuitsCount = math.Max(highestQuantityOfSuitsCount, diamondCount)
		case suit.Heart:
			heartCount++
			highestQuantityOfSuitsCount = math.Max(highestQuantityOfSuitsCount, heartCount)
		default:
			spadeCount++
			highestQuantityOfSuitsCount = math.Max(highestQuantityOfSuitsCount, spadeCount)
		}
	}

	if highestQuantityOfSuitsCount < 4 ||
		(highestQuantityOfSuitsCount == 4 && h.IsCrib) {
		return 0
	}

	return int(highestQuantityOfSuitsCount)
}

func nobsCount(h *Hand) int {
	var cutSuit suit.Suit
	jackSuits := make(map[suit.Suit]bool)

	for _, card := range h.Cards {
		if card.IsCut {
			cutSuit = card.Suit
			continue
		}

		if card.SortOrder == 11 {
			jackSuits[card.Suit] = true
		}
	}

	if _, ok := jackSuits[cutSuit]; ok {
		return 1
	}

	return 0
}
