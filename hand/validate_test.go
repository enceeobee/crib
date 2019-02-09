package hand

import (
	"testing"

	"github.com/enceeobee/crib/card"
)

func TestValidate(t *testing.T) {
	// Too few cards
	hand := &Hand{
		Cards: []*card.Card{
			&card.Card{},
			&card.Card{},
			&card.Card{},
		},
	}
	err := hand.Validate()

	if err == nil || err.Error() != "Incorrect number of cards in hand" {
		t.Errorf("Too few cards failed, err = %q", err.Error())
	}

	// Too many cards
	hand = &Hand{
		Cards: []*card.Card{
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
		},
	}
	err = hand.Validate()

	if err == nil || err.Error() != "Incorrect number of cards in hand" {
		t.Errorf("Too many cards failed, err = %q", err.Error())
	}

	// No cut card
	hand = &Hand{
		Cards: []*card.Card{
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
		},
	}
	err = hand.Validate()

	if err == nil || err.Error() != "Hand does not include cut card" {
		t.Errorf("No cut card failed, err = %q", err.Error())
	}

	// Happy path
	hand = &Hand{
		Cards: []*card.Card{
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{},
			&card.Card{IsCut: true},
		},
	}
	err = hand.Validate()

	if err != nil {
		t.Errorf("Valid hand failed, err = %q", err.Error())
	}
}
