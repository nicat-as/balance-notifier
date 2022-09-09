package main

import (
	"github.com/go-co-op/gocron"
	"github.com/nicat-as/balance-notifier/pkg/notification"
	"github.com/nicat-as/balance-notifier/pkg/provider"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	log.Println("action.balance-notifier.app.start")
	initEnv()
	log.Println("action.schedule.start")
	s := gocron.NewScheduler(time.UTC)
	accountNo := "410500D8402765243103"
	var amount float64 = 30

	s.Cron(os.Getenv("CRON.NOTIFY")).Do(notification.Notify, &amount, accountNo, provider.Kapital)

	s.Cron(os.Getenv("CRON.RESET")).Do(notification.ResetThreshold, accountNo, amount)

	log.Println("action.tasks.scheduled")
	s.StartBlocking()
}

func initEnv() {
	log.Println("action.env.initialize.start")
	b, err := ioutil.ReadFile("app/app.properties")
	if err != nil {
		log.Fatal(err)
		return
	}
	split := strings.Split(string(b), "\n")
	for _, line := range split {
		if len(line) > 0 && line[:1] == "#" {
			continue
		}
		lineSplit := strings.Split(line, "=")
		if len(lineSplit) == 2 {
			err := os.Setenv(lineSplit[0], lineSplit[1])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println("action.env.initialize.end")
}
