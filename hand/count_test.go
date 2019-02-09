package hand

import (
	"testing"

	"github.com/enceeobee/crib/card"
	"github.com/enceeobee/crib/suit"
)

type countTest struct {
	desc string
	hand *Hand
	x    int
}

var gooseEgg = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Spade, SortOrder: 6},
		&card.Card{Value: 8, Suit: suit.Spade, SortOrder: 8},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var goodJackOnly = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Heart, SortOrder: 2, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Spade, SortOrder: 6},
		&card.Card{Value: 8, Suit: suit.Spade, SortOrder: 8},
		&card.Card{Label: "Jack", Value: 10, Suit: suit.Heart, SortOrder: 11},
	},
}
var one15 = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 3, Suit: suit.Club, SortOrder: 3, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Spade, SortOrder: 6},
		&card.Card{Value: 9, Suit: suit.Spade, SortOrder: 9},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var two15s = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Spade, SortOrder: 6},
		&card.Card{Value: 9, Suit: suit.Spade, SortOrder: 9},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var onePair = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Spade, SortOrder: 6},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var twoPair = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 10, Suit: suit.Diamond, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var threePair = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 4, Suit: suit.Diamond, SortOrder: 4},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var fourPair = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4, IsCut: true},
		&card.Card{Value: 4, Suit: suit.Spade, SortOrder: 4},
		&card.Card{Value: 4, Suit: suit.Diamond, SortOrder: 4},
		&card.Card{Value: 4, Suit: suit.Heart, SortOrder: 4},
		&card.Card{Label: "Queen", Value: 10, Suit: suit.Heart, SortOrder: 12},
	},
}
var runOf3 = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 9, Suit: suit.Club, SortOrder: 9, IsCut: true},
		&card.Card{Value: 2, Suit: suit.Spade, SortOrder: 2},
		&card.Card{Value: 1, Suit: suit.Spade, SortOrder: 1, Label: "Ace"},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Label: "Jack", Value: 10, Suit: suit.Heart, SortOrder: 11},
	},
}
var runOf4 = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 9, Suit: suit.Club, SortOrder: 9, IsCut: true},
		&card.Card{Value: 8, Suit: suit.Spade, SortOrder: 8},
		&card.Card{Value: 1, Suit: suit.Spade, SortOrder: 1, Label: "Ace"},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Label: "Jack", Value: 10, Suit: suit.Heart, SortOrder: 11},
	},
}
var runOf5 = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 9, Suit: suit.Club, SortOrder: 9, IsCut: true},
		&card.Card{Value: 8, Suit: suit.Spade, SortOrder: 8},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 12, Label: "Queen"},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Label: "Jack", Value: 10, Suit: suit.Heart, SortOrder: 11},
	},
}
var flushOf4NoCrib = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2},
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Club, SortOrder: 6},
		&card.Card{Value: 8, Suit: suit.Club, SortOrder: 8},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 10, IsCut: true},
	},
}
var flushOf5NoCrib = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2},
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Club, SortOrder: 6},
		&card.Card{Value: 8, Suit: suit.Club, SortOrder: 8},
		&card.Card{Value: 10, Suit: suit.Club, SortOrder: 10, IsCut: true},
	},
}
var flushOf4Crib = &Hand{
	IsCrib: true,
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2},
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Club, SortOrder: 6},
		&card.Card{Value: 8, Suit: suit.Club, SortOrder: 8},
		&card.Card{Value: 10, Suit: suit.Diamond, SortOrder: 10, IsCut: true},
	},
}
var flushOf5Crib = &Hand{
	IsCrib: true,
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2},
		&card.Card{Value: 4, Suit: suit.Club, SortOrder: 4},
		&card.Card{Value: 6, Suit: suit.Club, SortOrder: 6},
		&card.Card{Value: 8, Suit: suit.Club, SortOrder: 8},
		&card.Card{Value: 10, Suit: suit.Club, SortOrder: 10, IsCut: true},
	},
}
var cleanEight = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 12, Label: "Queen"},
		&card.Card{Value: 10, Suit: suit.Diamond, SortOrder: 11, Label: "Jack"},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 10, IsCut: true},
	},
}
var dozen = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 2, Suit: suit.Club, SortOrder: 2},
		&card.Card{Value: 4, Suit: suit.Heart, SortOrder: 4},
		&card.Card{Value: 3, Suit: suit.Spade, SortOrder: 3},
		&card.Card{Value: 10, Suit: suit.Diamond, SortOrder: 11, Label: "Jack"},
		&card.Card{Value: 3, Suit: suit.Heart, SortOrder: 3, IsCut: true},
	},
}
var sixteen = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 10, Suit: suit.Club, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Diamond, SortOrder: 11, Label: "Jack"},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 12, Label: "Queen"},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 11, Label: "Jack", IsCut: true},
	},
}
var tripleRun = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 10, Suit: suit.Club, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Spade, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Diamond, SortOrder: 12, Label: "Queen of Diamonds"},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 10},
		&card.Card{Value: 10, Suit: suit.Heart, SortOrder: 11, Label: "Jack", IsCut: true},
	},
}
var twentyFour = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 6, Suit: suit.Club, SortOrder: 6},
		&card.Card{Value: 7, Suit: suit.Spade, SortOrder: 7},
		&card.Card{Value: 7, Suit: suit.Diamond, SortOrder: 7},
		&card.Card{Value: 8, Suit: suit.Heart, SortOrder: 8},
		&card.Card{Value: 8, Suit: suit.Heart, SortOrder: 8, IsCut: true},
	},
}
var twentyEight = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 5, Suit: suit.Club, SortOrder: 5},
		&card.Card{Value: 5, Suit: suit.Spade, SortOrder: 5},
		&card.Card{Value: 5, Suit: suit.Diamond, SortOrder: 5},
		&card.Card{Value: 5, Suit: suit.Heart, SortOrder: 5},
		&card.Card{Label: "Jack", Value: 10, Suit: suit.Heart, SortOrder: 11, IsCut: true},
	},
}
var twentyNine = &Hand{
	Cards: []*card.Card{
		&card.Card{Value: 5, Suit: suit.Club, SortOrder: 5},
		&card.Card{Value: 5, Suit: suit.Spade, SortOrder: 5},
		&card.Card{Value: 5, Suit: suit.Diamond, SortOrder: 5},
		&card.Card{Value: 5, Suit: suit.Heart, SortOrder: 5, IsCut: true},
		&card.Card{Label: "Jack", Value: 10, Suit: suit.Heart, SortOrder: 11},
	},
}

func TestCount(t *testing.T) {
	tests := []countTest{
		{
			desc: "Zero points",
			hand: gooseEgg,
			x:    0,
		},
		{
			desc: "Good Jack",
			hand: goodJackOnly,
			x:    1,
		},
		{
			desc: "1 15",
			hand: one15,
			x:    2,
		},
		{
			desc: "2 15",
			hand: two15s,
			x:    4,
		},
		{
			desc: "1 pair",
			hand: onePair,
			x:    2,
		},
		{
			desc: "2 pair",
			hand: twoPair,
			x:    4,
		},
		{
			desc: "3 pair",
			hand: threePair,
			x:    6,
		},
		{
			desc: "4 pair",
			hand: fourPair,
			x:    12,
		},
		{
			desc: "Run of 3",
			hand: runOf3,
			x:    3,
		},
		{
			desc: "Run of 4",
			hand: runOf4,
			x:    4,
		},
		{
			desc: "Run of 5",
			hand: runOf5,
			x:    5,
		},
		{
			desc: "Flush of 4 in regular hand",
			hand: flushOf4NoCrib,
			x:    4,
		},
		{
			desc: "Flush of 5 in regular hand",
			hand: flushOf5NoCrib,
			x:    5,
		},
		{
			desc: "Flush of 4 in crib",
			hand: flushOf4Crib,
			x:    0,
		},
		{
			desc: "Flush of 5 in crib",
			hand: flushOf5Crib,
			x:    5,
		},
		{
			desc: "Simple double run",
			hand: cleanEight,
			x:    8,
		},
		{
			desc: "Two 15s and double run",
			hand: dozen,
			x:    12,
		},
		{
			desc: "Triple run",
			hand: tripleRun,
			x:    15,
		},
		{
			desc: "Double double run",
			hand: sixteen,
			x:    16,
		},
		{
			desc: "24 hand",
			hand: twentyFour,
			x:    24,
		},
		{
			desc: "28 hand",
			hand: twentyEight,
			x:    28,
		},
		{
			desc: "29 hand!!!",
			hand: twentyNine,
			x:    29,
		},
	}

	for _, test := range tests {
		if actual := test.hand.Count(); actual != test.x {
			t.Errorf("%q failed; Got %d; Expected %d", test.desc, actual, test.x)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twentyFour.Count()
	}
}

func TestFifteensCount(t *testing.T) {
	tests := []countTest{
		{
			desc: "5 cards",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 1},
					&card.Card{Value: 2},
					&card.Card{Value: 3},
					&card.Card{Value: 4},
					&card.Card{Value: 5},
				},
			},
			x: 2,
		},
		{
			desc: "4 cards",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 2},
					&card.Card{Value: 2},
					&card.Card{Value: 1},
					&card.Card{Value: 5},
					&card.Card{Value: 6},
				},
			},
			x: 2,
		},
		{
			desc: "3 cards twice",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 3},
					&card.Card{Value: 3},
					&card.Card{Value: 4},
					&card.Card{Value: 5},
					&card.Card{Value: 6},
				},
			},
			x: 4,
		},
		{
			desc: "Single 15 - 2 cards",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 8},
					&card.Card{Value: 4},
					&card.Card{Value: 4},
					&card.Card{Value: 5},
					&card.Card{Value: 10},
				},
			},
			x: 2,
		},
		{
			desc: "7s and 8s",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 7},
					&card.Card{Value: 7},
					&card.Card{Value: 7},
					&card.Card{Value: 8},
					&card.Card{Value: 8},
				},
			},
			x: 12,
		},
		{
			desc: "10 and 5s",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 10},
					&card.Card{Value: 1},
					&card.Card{Value: 5},
					&card.Card{Value: 5},
					&card.Card{Value: 5},
				},
			},
			x: 8,
		},
		{
			desc: "3 5s",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Value: 8},
					&card.Card{Value: 1},
					&card.Card{Value: 5},
					&card.Card{Value: 5},
					&card.Card{Value: 5},
				},
			},
			x: 2,
		},
	}

	runTests(tests, fifteensCount, t)
}

func TestPairsCount(t *testing.T) {
	tests := []countTest{
		{
			desc: "One pair",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 1},
					&card.Card{SortOrder: 2},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 5},
				},
			},
			x: 2,
		},
		{
			desc: "Two pair",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 1},
					&card.Card{SortOrder: 2},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 2},
				},
			},
			x: 4,
		},
		{
			desc: "Three of a kind",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 1},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 2},
				},
			},
			x: 6,
		},
		{
			desc: "Four of a kind",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 1},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 3},
				},
			},
			x: 12,
		},
	}

	runTests(tests, pairsCount, t)
}

func BenchmarkPairsCount(b *testing.B) {
	doubleDoubleRun := &Hand{
		Cards: []*card.Card{
			&card.Card{SortOrder: 4},
			&card.Card{SortOrder: 3},
			&card.Card{SortOrder: 4},
			&card.Card{SortOrder: 5},
			&card.Card{SortOrder: 3},
		},
	}

	for n := 0; n < b.N; n++ {
		pairsCount(doubleDoubleRun)
	}
}
func TestRunsCount(t *testing.T) {
	tests := []countTest{
		{
			desc: "Run of 2",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 11},
					&card.Card{SortOrder: 13},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 8},
					&card.Card{SortOrder: 3},
				},
			},
			x: 0,
		},
		{
			desc: "Run of 3",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 11},
					&card.Card{SortOrder: 13},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 5},
					&card.Card{SortOrder: 3},
				},
			},
			x: 3,
		},
		{
			desc: "Double run of 3",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 13},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 5},
					&card.Card{SortOrder: 3},
				},
			},
			x: 6,
		},
		{
			desc: "Triple run of 3",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 5},
					&card.Card{SortOrder: 3},
				},
			},
			x: 9,
		},
		{
			desc: "Double double run",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 3},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 5},
					&card.Card{SortOrder: 3},
				},
			},
			x: 12,
		},
		{
			desc: "Run of 4",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 6},
					&card.Card{SortOrder: 13},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 5},
					&card.Card{SortOrder: 3},
				},
			},
			x: 4,
		},
		{
			desc: "Double run of 4",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 2},
					&card.Card{SortOrder: 4},
					&card.Card{SortOrder: 5},
					&card.Card{SortOrder: 3},
				},
			},
			x: 8,
		},
	}

	runTests(tests, runsCount, t)
}

func BenchmarkRunsCount(b *testing.B) {
	doubleDoubleRun := &Hand{
		Cards: []*card.Card{
			&card.Card{SortOrder: 4},
			&card.Card{SortOrder: 3},
			&card.Card{SortOrder: 4},
			&card.Card{SortOrder: 5},
			&card.Card{SortOrder: 3},
		},
	}

	// run the runsCount function b.N times
	for n := 0; n < b.N; n++ {
		runsCount(doubleDoubleRun)
	}
}
func TestFlushCount(t *testing.T) {
	tests := []countTest{
		{
			desc: "3 similar suits",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Heart},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Diamond},
					&card.Card{Suit: suit.Club},
				},
			},
			x: 0,
		},
		{
			desc: "4 similar suits not crib",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Heart},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
				},
			},
			x: 4,
		},
		{
			desc: "4 similar suits in crib",
			hand: &Hand{
				IsCrib: true,
				Cards: []*card.Card{
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Heart},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
				},
			},
			x: 0,
		},
		{
			desc: "5 similar suits not crib",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
				},
			},
			x: 5,
		},
		{
			desc: "5 similar suits in crib",
			hand: &Hand{
				IsCrib: true,
				Cards: []*card.Card{
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
					&card.Card{Suit: suit.Club},
				},
			},
			x: 5,
		},
	}

	runTests(tests, flushCount, t)
}
func TestNobsCount(t *testing.T) {
	tests := []countTest{
		{
			desc: "No Jack",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3, Suit: suit.Club},
					&card.Card{SortOrder: 5, Suit: suit.Heart},
					&card.Card{SortOrder: 7, Suit: suit.Club},
					&card.Card{SortOrder: 2, Suit: suit.Diamond, IsCut: true},
					&card.Card{SortOrder: 1, Suit: suit.Club},
				},
			},
			x: 0,
		},
		{
			desc: "Has Incorrect Jack",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3, Suit: suit.Club, IsCut: true},
					&card.Card{SortOrder: 11, Suit: suit.Heart},
					&card.Card{SortOrder: 7, Suit: suit.Club},
					&card.Card{SortOrder: 2, Suit: suit.Diamond},
					&card.Card{SortOrder: 1, Suit: suit.Club},
				},
			},
			x: 0,
		},
		{
			desc: "Jack is cut",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3, Suit: suit.Club},
					&card.Card{SortOrder: 11, Suit: suit.Heart, IsCut: true},
					&card.Card{SortOrder: 7, Suit: suit.Club},
					&card.Card{SortOrder: 2, Suit: suit.Diamond},
					&card.Card{SortOrder: 1, Suit: suit.Club},
				},
			},
			x: 0,
		},
		{
			desc: "Has the Good Jack",
			hand: &Hand{
				Cards: []*card.Card{
					&card.Card{SortOrder: 3, Suit: suit.Heart, IsCut: true},
					&card.Card{SortOrder: 11, Suit: suit.Heart},
					&card.Card{SortOrder: 7, Suit: suit.Club},
					&card.Card{SortOrder: 2, Suit: suit.Diamond},
					&card.Card{SortOrder: 1, Suit: suit.Club},
				},
			},
			x: 1,
		},
	}

	runTests(tests, nobsCount, t)
}

func runTests(tests []countTest, testingFn func(*Hand) int, t *testing.T) {
	for _, test := range tests {
		if actual := testingFn(test.hand); actual != test.x {
			t.Errorf("%q failed; Got %d; Expected %d", test.desc, actual, test.x)
		}
	}
}
