package account

import (
	"errors"
	account2 "github.com/nicat-as/balance-notifier/pkg/account"
	"github.com/nicat-as/balance-notifier/pkg/http"
	"github.com/nicat-as/balance-notifier/pkg/provider"
	"github.com/nicat-as/balance-notifier/pkg/provider/authorization"
	"log"
	"os"
	"strconv"
)

var _ account2.AccountClient[string, any] = (*KapitalAccountClient)(nil)

func (k KapitalAccountClient) FetchAccountList(_ string) (*account2.AccountList[any], error) {
	log.Printf("action.account.FetchAccountList.start")
	accountList, err := http.DoRequest[string, provider.KapitalResponseData[AccountList]](
		provider.GetEndpoint(provider.KapitalAccountList),
		http.Get,
		map[string]string{
			"Authorization": authorization.KapitalAuthorization{}.GetToken(authorization.UsernamePassword{
				Username: os.Getenv(authorization.Username),
				Password: os.Getenv(authorization.Password),
				JwtToken: nil,
			}),
		},
		"",
	)
	if err != nil {
		return nil, err
	}
	var accounts []account2.Account[any]
	for _, a := range accountList.ResponseData.AccountsList {
		amount, err := strconv.ParseFloat(a.CurrentAmount, 64)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account2.Account[any]{
			AccountNo:        a.AccountNo,
			CurrentBalance:   amount,
			Currency:         a.Currency,
			AdditionalFields: nil,
		})
	}

	if len(accounts) <= 0 {
		return nil, errors.New("account:not-found")
	}
	log.Printf("action.account.FetchAccountList.end")
	return &account2.AccountList[any]{
		Data: accounts,
	}, err
}
