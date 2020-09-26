package stats

import (
	"strings"

	"github.com/ilhom0258/bank/pkg/bank/types"
)

// TotalInCategory - function for calculating sum in category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	totalValue := types.Money(0)
	searchCategory := string(category)

	for _, payment := range payments {

		paymentCategoryValue := string(payment.Category)

		if strings.ToLower(paymentCategoryValue) != strings.ToLower(searchCategory) {
			continue
		}
		totalValue += payment.Amount
	}

	return totalValue
}
