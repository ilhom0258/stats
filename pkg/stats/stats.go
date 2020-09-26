package stats

import (
	"strings"

	"github.com/ilhom0258/bank/pkg/types"
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

//Avg - рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var avgPayment types.Money = types.Money(0)
	for _, payment := range payments {
		avgPayment += payment.Amount
	}
	avgPayment /= types.Money(len(payments))
	return avgPayment
}
