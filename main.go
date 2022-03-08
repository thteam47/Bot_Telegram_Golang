package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gorilla/mux"
	"github.com/thteam47/Bot_Telegram_Golang/drive"
)

var bot *drive.BotDb
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/approve"),
		tgbotapi.NewKeyboardButton("/decline"),
	),
)

func sendMess(rw http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pass := r.FormValue("pass")
	x := make(chan int)

	if user == "admin" && pass == "admin" {
		msg := tgbotapi.NewMessage(5172255611, "Admin is Logging ... Choose options:")
		msg.ReplyMarkup = numericKeyboard
		if _, err := bot.BotTel.Send(msg); err != nil {
			log.Panic(err)
		}

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 0
		ch := make(chan tgbotapi.Update, bot.BotTel.Buffer)

		go func() {
			for {
				select {
				case st := <-x:
					if st == 5 {
						st = 10
						close(ch)
						return
					}
				default:
				}
				updates, err := bot.BotTel.GetUpdates(u)
				if err != nil {
					continue
				}

				for _, update := range updates {
					if update.UpdateID >= u.Offset {
						u.Offset = update.UpdateID + 1
						ch <- update
					}
				}
			}
		}()
		var check bool
		resp := ""
		for update := range ch {
			if update.Message == nil {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "close")
			if update.Message.Command() == "approve" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Login Success")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				check = true
				if _, err := bot.BotTel.Send(msg); err != nil {
					log.Panic(err)
				}
				x <- 5
				close(x)
				break
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Decline")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				if _, err := bot.BotTel.Send(msg); err != nil {
					log.Panic(err)
				}
				check = false
				x <- 5
				close(x)
				break
			}

			// default:
			// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Admin is Logging ... Choose options:")
			// 	msg.ReplyMarkup = numericKeyboard

			// }

		}

		if check {
			resp = "Login Succes"

		} else {
			resp = "Decline"
		}
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(resp)
	}

}
func getlogin(rw http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("template.html"))
	tmp.Execute(rw, nil)
}
func sendMessAu() {
	for {
		if bot != nil {
			msg := tgbotapi.NewMessage(5172255611, "Report:")
			if _, err := bot.BotTel.Send(msg); err != nil {
				fmt.Println(err)
			}
			time.Sleep(20 * time.Second)
		}
	}
}
func main() {
	bot = drive.ConnectBot("zg")
	fmt.Println(bot)
	go sendMessAu()
	r := mux.NewRouter()
	r.HandleFunc("/login", getlogin).Methods("GET")
	r.HandleFunc("/login", sendMess).Methods("POST")

	http.ListenAndServe(":8080", r)

}
