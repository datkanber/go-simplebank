package util

// Constants for supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	TRY = "TRY"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TRY, CAD:
		return true
	default:
		return false
	}
}
