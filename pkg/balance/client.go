package balance

type BalanceClient interface {
	FetchCurrentBalance(accountNo string) (*BalanceInfo, error)
}
