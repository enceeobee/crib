package hand

import "github.com/enceeobee/crib/card"

// Hand represents a cribbage hand
type Hand struct {
	Cards  []*card.Card `json:"cards"`
	IsCrib bool         `json:"isCrib"`
}
