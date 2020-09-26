package stats

import (
	"fmt"

	"github.com/ilhom0258/bank/v2/pkg/types"
)

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10222200,
			Category: "Auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   5000,
			Category: "Auto",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   4444,
			Category: "Book",
			Status:   types.StatusOk,
		},
		{
			ID:       4,
			Amount:   3000,
			Category: "Book",
			Status:   types.StatusOk,
		},
	}
	result := TotalInCategory(payments, "Book")
	fmt.Println(result)
	//Output:
	//7444
}
func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   12,
			Category: "Auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   12,
			Category: "Auto",
			Status:   types.StatusOk,
		},
	}
	result := Avg(payments)
	fmt.Println(result)
	//Output:
	//12
}
