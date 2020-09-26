package stats

import (
	"reflect"
	"testing"

	"github.com/ilhom0258/bank/v2/pkg/types"
)

func TestCategoriesAvg_success(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 1_000_00},
		{ID: 3, Category: "auto", Amount: 1_000_00},
		{ID: 4, Category: "school", Amount: 1_000_00},
		{ID: 5, Category: "fun", Amount: 1_000_00},
		{ID: 6, Category: "food", Amount: 1_000_00},
		{ID: 7, Category: "school", Amount: 1_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto":   100000,
		"food":   100000,
		"school": 100000,
		"fun":    100000,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected %v, actual %v", expected, result)
	}
}
