package main

import (
	"bufio"
	"context"
	"fmt"
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

func InitObservable() {
	file_swear, _ := os.Open("dirtytalk.txt")
	file_name, _ := os.Open("sensitive_name.txt")
	defer file_swear.Close()
	defer file_name.Close()

	dirtytalk := []string{}
	names := []string{}

	buf := bufio.NewScanner(file_swear)
	for buf.Scan() {
		str := buf.Text()
		dirtytalk = append(dirtytalk, str)
	}

	buf = bufio.NewScanner(file_name)
	for buf.Scan() {
		str := buf.Text()
		names = append(names, str)
	}

	ObservableMsg = ObservableMsg.Filter(func(i interface{}) bool {
		for _, v := range dirtytalk {
			if strings.Contains(i.(string), v) {
				fmt.Println(v)
				return false
			}
		}
		return true
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		str := i.(string)
		for _, v := range names {
			if strings.Contains(str, v) {
				tmp := v[0:3] + "*"
				if len(v) > 2 {
					tmp += v[6:]
				}
				str = strings.Replace(str, v, tmp, -1)
			}
		}
		return str, nil
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
