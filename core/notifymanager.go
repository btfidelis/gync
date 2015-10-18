package core 

import (
	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

const APP_ICON = "./../icon.png"

func InitNotifyManager() {
	notify = notificator.New(notificator.Options{
		DefaultIcon: APP_ICON,
		AppName:     "Gync",
  	})
}

func Notify(message string, title string) {
	notify.Push(title, message, "")
}
