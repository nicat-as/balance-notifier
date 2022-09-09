package notification

import (
	"fmt"
	"github.com/nicat-as/balance-notifier/pkg/factory"
	"github.com/nicat-as/balance-notifier/pkg/notification/channel"
	"github.com/nicat-as/balance-notifier/pkg/provider"
	"log"
)

var thresholdMap = map[string]float64{}

func Notify(threshold *float64, accountNo string, provider provider.Provider) {
	lastThreshold := *threshold
	if v, ok := thresholdMap[accountNo]; ok {
		lastThreshold = v

	}
	log.Printf("action.notification.Notify.start - threshold: %.2f , accountNo: %s, provider: %s \n",
		lastThreshold, accountNo, provider)
	balanceClient, err := factory.BalanceClientFactory(provider)
	if err != nil {
		log.Fatalf("provider %s not found", provider)
		return
	}
	balanceInfo, err := balanceClient.FetchCurrentBalance(accountNo)
	if err != nil {
		log.Fatalf("balance fetching failed for acc: %s | %v", accountNo, err)
		return
	}

	if balanceInfo.Amount > lastThreshold {
		log.Printf("balance increased: %.2f and set value as threshold\n", balanceInfo.Amount)
		thresholdMap[accountNo] = balanceInfo.Amount
		message := fmt.Sprintf("your balance increased: %.2f %s", balanceInfo.Amount, balanceInfo.Currency)
		err := channel.FnSendMessage(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("action.notification.Notify.end - threshold: %.2f , accountNo: %s, provider: %s \n",
		lastThreshold, accountNo, provider)
}

func ResetThreshold(accountNo string, resetValue float64) {
	log.Printf("action.notification.ResetThreshold.start - accountNo: %s, resetValue: %.2f", accountNo, resetValue)
	thresholdMap[accountNo] = resetValue
	log.Printf("action.notification.ResetThreshold.end - accountNo: %s, resetValue: %.2f", accountNo, resetValue)
}
