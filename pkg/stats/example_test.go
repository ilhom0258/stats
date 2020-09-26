package stats

import (
	"fmt"
)

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   1000,
			Category: "Auto",
		},
		{
			ID:       2,
			Amount:   3000,
			Category: "Auto",
		},
		{
			ID:       3,
			Amount:   1000,
			Category: "Book",
		},
		{
			ID:       4,
			Amount:   3000,
			Category: "Book",
		},
	}
	result := TotalInCategory(payments, "Book")
	fmt.Println(result)
	//Output:
	//4000
}
