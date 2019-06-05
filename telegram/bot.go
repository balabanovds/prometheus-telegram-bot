package telegram

import (
	"log"

	"github.com/balabanovds/prometheus-telegram-bot/tor"
	"github.com/balabanovds/prometheus-telegram-bot/util"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot struct
type Bot struct {
	API *tgbot.BotAPI
}

// NewBot creates telegram bot, set webhook
func NewBot() (*Bot, error) {
	token := util.Cfg.Telegram.Token
	bot, err := tgbot.NewBotAPIWithClient(token, tor.GetClient())
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// link := util.Cfg.Telegram.Webhook + bot.Token

	// bot.RemoveWebhook()

	// if _, err := bot.SetWebhook(tgbot.NewWebhookWithCert(link, util.Cfg.Telegram.Cert)); err != nil {
	// 	return nil, err
	// }

	// info, err := bot.GetWebhookInfo()
	// if err != nil {
	// 	return nil, err
	// }
	// if info.LastErrorDate != 0 {
	// 	return nil, fmt.Errorf("Telegram callback failed: %s", info.LastErrorMessage)
	// }

	return &Bot{API: bot}, nil
}

func (b *Bot) Run() {

	ucfg := tgbot.NewUpdate(0)
	ucfg.Timeout = 60

	updates, err := b.API.GetUpdatesChan(ucfg)
	if err != nil {
		log.Fatalf("Updates err: %v", err)
	}

	// updates := b.API.ListenForWebhook("/" + b.API.Token)
	// go http.ListenAndServeTLS(util.Cfg.Service.TLSAddr, util.Cfg.Telegram.Cert, util.Cfg.Telegram.Key, nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("Update: %+v", update)
		// username := u.Message.From.UserName
		// chatID := u.Message.Chat.ID
		// text :=

		// if u.Message == nil || !ContainsInt64(tgbot.chats, u.Message.Chat.ID) || !strings.HasPrefix(u.Message.Text, BotCommandPrefix) {
		// 	continue
		// }

		// loggerInfo.Printf("id: %v, text: %v\n", u.Message.Chat.ID, u.Message.Text)

		// command := tgbot.commandRegister.GetCommand(u.Message.Text)
		// msg, err := command.MakeResponseMessage()
		// if err != nil {
		// 	loggerError.Println(err)
		// }

		// msgConfig := tgbotapi.NewMessage(u.Message.Chat.ID, msg)
		// msgConfig.ParseMode = "HTML"

		// if _, err := tgbot.botAPI.Send(msgConfig); err != nil {
		// 	loggerError.Println(err)
		// }
	}
}

// func (tgbot *TelegramBot) HandlePrometheusAlert(w http.ResponseWriter, r *http.Request) {
// 	var alertResponse AlertmanagerWebhookResponse
// 	if err := json.NewDecoder(r.Body).Decode(&alertResponse); err != nil {
// 		loggerError.Println(err)
// 	}

// 	var buf bytes.Buffer

// 	for _, alert := range alertResponse.Alerts {
// 		if alert.Status == "resolved" && isResolveDisabled(alert.Labels.Alertname) {
// 			continue
// 		}

// 		buf.WriteString("🔥 Firing: 🔥\n")
// 		st, _ := time.Parse(time.RFC3339, alert.StartsAt)
// 		startTime := st.Format(time.RFC1123)

// 		var endTime string
// 		if alert.Status == "resolved" {
// 			et, _ := time.Parse(time.RFC3339, alert.EndsAt)
// 			endTime = et.Format(time.RFC1123)
// 		} else {
// 			endTime = "none"
// 		}

// 		msg := fmt.Sprintf("%s\n<code>%s</code>\nJob: %s\nStartsAt: %s\nEndsAt: %s\nStatus: %s\n",
// 			alert.Labels.Alertname, alert.Labels.Instance, alert.Labels.Job, startTime, endTime, alert.Status)
// 		buf.WriteString(msg)

// 		if alert.Annotations.Value != "" {
// 			buf.WriteString("Value: " + alert.Annotations.Value + "\n")
// 		}
// 	}

// 	for _, chatID := range tgbot.chats {
// 		msgConfig := tgbotapi.NewMessage(chatID, buf.String())
// 		msgConfig.ParseMode = "HTML"

// 		if _, err := tgbot.botAPI.Send(msgConfig); err != nil {
// 			loggerError.Println(err)
// 		}
// 	}
// }
