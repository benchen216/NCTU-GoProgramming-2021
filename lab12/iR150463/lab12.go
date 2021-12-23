package main

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/reactivex/rxgo/v2"
)

type client chan<- string // an outgoing message channel

var (
	entering      = make(chan client)
	leaving       = make(chan client)
	messages      = make(chan rxgo.Item) // all incoming client messages
	ObservableMsg = rxgo.FromChannel(messages)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	MessageBoardcast := ObservableMsg.Observe()
	for {
		select {
		case msg := <-MessageBoardcast:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg.V.(string)
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//clientWriter and main
func clientWriter(conn *websocket.Conn, ch <-chan string) {
	for msg := range ch {
		conn.WriteMessage(1, []byte(msg))
	}
}

func wshandle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "你是 " + who + "\n"
	messages <- rxgo.Of(who + " 來到了現場" + "\n")
	entering <- ch

	defer func() {
		log.Println("disconnect !!")
		leaving <- ch
		messages <- rxgo.Of(who + " 離開了" + "\n")
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		messages <- rxgo.Of(who + " 表示: " + string(msg))
	}
}

// dt: dirty talk
// sn: senstive name
func InitObservable() {
	dtFile, err1 := os.Open("dirtytalk.txt")
	snFile, err2 := os.Open("sensitive_name.txt")
	defer dtFile.Close()
	defer snFile.Close()

	if err1 != nil {
		log.Println("Error1: ", err1)
		os.Exit(1)
	}

	if err1 != nil {
		log.Println("Error2: ", err2)
		os.Exit(1)
	}

	dt := []string{}
	sn := make(map[string]string)

	dtScanner := bufio.NewScanner(dtFile)
	for dtScanner.Scan() {
		dt = append(dt, dtScanner.Text())
	}

	snScanner := bufio.NewScanner(snFile)
	for snScanner.Scan() {
		s_rune := []rune(snScanner.Text())
		new_rune := append(s_rune[:1], rune('*'))
		new_rune = append(new_rune, s_rune[2:]...)

		replace_word := ""
		for _, s := range new_rune {
			replace_word += string(s)
		}

		sn[snScanner.Text()] = replace_word
	}

	ObservableMsg = ObservableMsg.Filter(func(msg interface{}) bool {
		for word := range dt {
			if strings.Contains(msg.(string), dt[word]) {
				return false
			}
		}
		return true
	}).Map(func(_ context.Context, msg interface{}) (interface{}, error) {
		newMsg := msg.(string)

		for word := range sn {
			log.Println("debug: ", word, " ", sn[word])
			if strings.Contains(newMsg, word) {
				newMsg = strings.Replace(newMsg, word, sn[word], -1)
			}
		}

		return newMsg, nil
	})
}

func main() {
	InitObservable()
	go broadcaster()
	http.HandleFunc("/wschatroom", wshandle)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("server start at :8899")
	log.Fatal(http.ListenAndServe(":8899", nil))
}
