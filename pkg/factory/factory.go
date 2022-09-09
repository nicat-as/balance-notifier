package factory

import (
	"errors"
	"github.com/nicat-as/balance-notifier/pkg/balance"
	"github.com/nicat-as/balance-notifier/pkg/provider"
	balanceProvider "github.com/nicat-as/balance-notifier/pkg/provider/balance"
)

func BalanceClientFactory(providerType provider.Provider) (balance.BalanceClient, error) {
	switch providerType {
	case provider.Kapital:
		return balanceProvider.KapitalBalanceClient{}, nil
	default:
		return nil, errors.New("provider:not-found")

	}
}
