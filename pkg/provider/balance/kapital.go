package balance

import (
	balance2 "github.com/nicat-as/balance-notifier/pkg/balance"
	accountProvider "github.com/nicat-as/balance-notifier/pkg/provider/account"
	"log"
)

var _ balance2.BalanceClient = (*KapitalBalanceClient)(nil)

func (k KapitalBalanceClient) FetchCurrentBalance(accountNo string) (*balance2.BalanceInfo, error) {
	log.Printf("action.balance.FetchCurrentBalance.start - accountNo: %s\n", accountNo)
	accounts, err := accountProvider.KapitalAccountClient{}.FetchAccountList(accountNo)
	if err != nil {
		log.Fatal("action.balance.FetchCurrentBalance.error - ", err)
		return nil, err
	}
	for _, account := range accounts.Data {
		if account.AccountNo == accountNo {
			log.Println("account found")
			return &balance2.BalanceInfo{
				Amount:   account.CurrentBalance,
				Currency: balance2.Currency(account.Currency),
			}, err
		}
	}
	log.Printf("action.balance.FetchCurrentBalance.end - accountNo: %s\n", accountNo)
	return nil, err
}
