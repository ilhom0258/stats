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

func TestPeriodsDynamic_1(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("invelid output: expected %v, actual %v", expected, result)
	}
}

func TestPeriodsDynamic_2(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("invelid output: expected %v, actual %v", expected, result)
	}
}

func TestPeriodsDynamic_3(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("invelid output: expected %v, actual %v", expected, result)
	}
}

func TestPeriodsDynamic_4(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}
	expected := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("invelid output: expected %v, actual %v", expected, result)
	}
}
