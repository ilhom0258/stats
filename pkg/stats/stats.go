package stats

import (
	"strings"

	"github.com/ilhom0258/bank/v2/pkg/types"
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
		if payment.Status == types.StatusFail {
			continue
		}

		totalValue += payment.Amount
	}

	return totalValue
}

//Avg - рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var avgPayment types.Money = types.Money(0)
	var amountPayment types.Money = types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		amountPayment++
		avgPayment += payment.Amount
	}

	avgPayment /= amountPayment
	return avgPayment
}

// CategoriesAvg calculates avg spending for each category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	categoriesQt := map[types.Category]int64{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		categoriesQt[payment.Category] ++
	}
	for key := range categories {
		categories[key] /= types.Money(categoriesQt[key])
	}
	return categories
}
