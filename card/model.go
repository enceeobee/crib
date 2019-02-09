package card

import "github.com/enceeobee/crib/suit"

// TODO - suit should probably be a pointer

// Card represents a playing card
type Card struct {
	Value     int       `json:"value"`
	Suit      suit.Suit `json:"suit"`
	Label     string    `json:"label"`
	SortOrder int       `json:"sortOrder"`
	IsCut     bool      `json:"isCut"`
}
