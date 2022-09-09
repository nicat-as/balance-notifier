package balance

type BalanceInfo struct {
	Amount   float64
	Currency Currency
}

type Currency string
