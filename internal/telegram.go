package internal

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

type TelegramIntegration struct {
	bot     *tgbotapi.BotAPI
	replier *Replier
}

func NewTelegramIntegration(replier *Replier) *TelegramIntegration {
	return &TelegramIntegration{
		bot:     nil,
		replier: replier,
	}
}

func (t *TelegramIntegration) Init() {
	if t.bot != nil {
		return
	}
	var err error
	if token, exists := os.LookupEnv("TELEGRAM_TOKEN"); exists {
		t.bot, err = tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Authorized on account %s", t.bot.Self.UserName)
		go t.process()
	} else {
		log.Fatal("Please setup the TELEGRAM_TOKEN environment variable")
	}
}

func (t *TelegramIntegration) process() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := t.bot.GetUpdatesChan(u)
	if err != nil {
		log.Printf("Can't take telegram updates channel.")
		return
	}

	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}

		log.Printf(
			"Message received on Telegram. From: [%s]; Message: [%s]",
			update.Message.Chat.UserName,
			update.Message.Text,
		)

		if t.replyCommand(update) {
			continue
		}

		sm := SuspiciousMessage{
			Text: &update.Message.Text,
			Link: nil,
		}

		vm, err := t.replier.CheckMessage(sm)
		if err != nil || vm == nil {
			errText := "Ih foi mal, mas rolou um erro absurdo aqui. " +
				"Manda mensagem pro @eduardompinto pra ele arrumar"
			t.replyMessage(update, errText)
		} else {
			t.replyMessage(update, vm.Verdict())
		}
	}

}

// replyCommand if the message is a command reply it and returns true, otherwise returns false
func (t *TelegramIntegration) replyCommand(update tgbotapi.Update) bool {
	switch update.Message.Command() {
	case "ajuda",
		"help",
		"inicio",
		"start":
		helpText := "Me envie uma mensagem e/ou notícia, e eu verifico se é uma fake news."
		t.replyMessage(update, helpText)
		return true
	default:
		return false
	}
}

func (t *TelegramIntegration) replyMessage(update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	_, _ = t.bot.Send(msg)
}
