package main

import (
	"log"
	"net/http"
	"context"
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

func InitObservable() {
	data1, _ := os.ReadFile("dirtytalk.txt")
	dirty := strings.Split(string(data1), "\n")
	data2, _ := os.ReadFile("sensitive_name.txt")
	name := strings.Split(string(data2), "\n")
	
	ObservableMsg = ObservableMsg.Filter(func(s0 interface{}) bool {
		for i:=0;i<len(dirty);i++ {
			if strings.Contains(s0.(string), dirty[i]) {
				return false
			}
		}
		return true
	}).Map(func(_ context.Context, s0 interface{}) (interface{}, error) {
		s := s0.(string)
		for i:=0;i<len(name);i++ {
			tmp := []rune(name[i])
			s = strings.Replace(s, name[i], string(tmp[0])+"*"+string(tmp[2:]), -1)
		}
		return s, nil
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
