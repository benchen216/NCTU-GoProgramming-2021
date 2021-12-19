package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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
	file, err := os.Open("dirtytalk.txt")
	if err != nil {
		fmt.Printf("file load error %s", err)
		os.Exit(1)
	}
	DirtyTalks := make([]string, 0)
	reader := bufio.NewReader(file)

	for {

		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		DirtyTalks = append(DirtyTalks, strings.TrimSuffix(str, "\n"))
	}
	file.Close()

	file, err = os.Open("sensitive_name.txt")
	if err != nil {
		fmt.Printf("file load error %s", err)
		os.Exit(2)
	}
	SensativeName := make([]string, 0)
	reader = bufio.NewReader(file)

	for {

		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		SensativeName = append(SensativeName, strings.TrimSuffix(str, "\n"))
	}
	file.Close()

	ObservableMsg = ObservableMsg.Filter(func(i interface{}) bool {
		for _, s := range DirtyTalks {
			if strings.Contains(i.(string), s) == true {
				return false
			}
		}
		return true
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		for _, s := range SensativeName {
			r := []rune(s)
			r[1] = '*'
			i = strings.Replace(i.(string), s, string(r), -1)
		}
		return i.(string), nil
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
