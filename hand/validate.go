package hand

import "errors"

// Validate determines if the provided Hand is valid
func (h *Hand) Validate() error {
	if len(h.Cards) != 5 {
		return errors.New("Incorrect number of cards in hand")
	}

	for _, card := range h.Cards {
		if card.IsCut {
			return nil
		}
	}

	return errors.New("Hand does not include cut card")
}
