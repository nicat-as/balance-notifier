package channel

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type TelegramClient interface {
	SendMessage(text string) error
}

type DefaultTelegramClient struct {
}

const (
	telegramPrefix  = "NOTIFICATION.TELEGRAM."
	telegramBaseUrl = telegramPrefix + "BASE_URL"
	telegramApiKey  = telegramPrefix + "API_KEY"
	telegramChatId  = telegramPrefix + "CHAT_ID"

	sendMessage = "/sendMessage"
)

var FnSendMessage = DefaultTelegramClient{}.SendMessage

func (d DefaultTelegramClient) SendMessage(text string) error {
	parsedUrl, err := url.Parse(getEnv(telegramBaseUrl, telegramApiKey) + sendMessage)
	if err != nil {
		log.Fatalln("SendMessage.error - ", err)
		return err
	}
	chatId, _ := strconv.ParseInt(os.Getenv(telegramChatId), 10, 64)
	query := parsedUrl.Query()
	query.Add("chat_id", strconv.FormatInt(chatId, 10))
	query.Add("text", text)
	query.Add("parse_mode", "HTML")
	parsedUrl.RawQuery = query.Encode()
	response, err := http.Get(parsedUrl.String())
	if err != nil {
		log.Fatalln("SendMessage.error - ", err)
		return err
	}
	if response.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(response.Body)
		log.Fatal(string(b))
	}
	return err
}

func getEnv(envs ...string) string {
	sb := strings.Builder{}
	for _, env := range envs {
		sb.WriteString(os.Getenv(env))
	}
	return sb.String()
}

var _ TelegramClient = (*DefaultTelegramClient)(nil)
