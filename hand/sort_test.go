package hand

import (
	"reflect"
	"testing"

	"github.com/enceeobee/crib/card"
)

func TestSort(t *testing.T) {
	actual := make([]int, 5)
	expected := []int{3, 5, 8, 9, 10}
	hand := &Hand{
		Cards: []*card.Card{
			&card.Card{SortOrder: 8},
			&card.Card{SortOrder: 10},
			&card.Card{SortOrder: 3},
			&card.Card{SortOrder: 5},
			&card.Card{SortOrder: 9},
		},
	}

	hand.Sort()

	for i := 0; i < 5; i++ {
		actual[i] = hand.Cards[i].SortOrder
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Sorting failed; Got %v; Expected %v", actual, expected)
	}
}
