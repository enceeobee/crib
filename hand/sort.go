package hand

import (
	"sort"

	"github.com/enceeobee/crib/card"
)

// BySortOrder implements sort.Interface for []*card.Card based on
// the SortOrder field.
type BySortOrder []*card.Card

func (bso BySortOrder) Len() int           { return len(bso) }
func (bso BySortOrder) Swap(i, j int)      { bso[i], bso[j] = bso[j], bso[i] }
func (bso BySortOrder) Less(i, j int) bool { return bso[i].SortOrder < bso[j].SortOrder }

// Sort sorts a hand
func (h *Hand) Sort() {
	sort.Sort(BySortOrder(h.Cards))
}
