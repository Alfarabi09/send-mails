package handlers

import (
	"fara/pkg/emails"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/robfig/cron/v3"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func PostSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Sending email.
	err1 := emails.SendAll()
	if err1 != nil {
		log.Fatal(err1)
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	fmt.Println("Email Sent!")
}

func PostDelay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	text = cron_format(text)
	// Sending email.
	MoscowTime, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	scheduler := cron.New(cron.WithLocation(MoscowTime))

	// stop scheduler
	defer scheduler.Stop()

	// set task, message will be sent in 2 minutes
	EntryID, err := scheduler.AddFunc(text, NotifyNewOrder)
	if err != nil || EntryID == 0 {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	// start scheduler
	go scheduler.Start()

	// trap SIGINT
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	fmt.Println("Email Sent!")
}

func NotifyNewOrder() {
	err1 := emails.SendAll()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + "NotifyNewOrder")
}

func cron_format(s string) string {
	if len(s) > 5 {
		return "err"
	}
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return "err"
	}
	text := split[1] + " " + split[0] + " * * *"
	return text
}
