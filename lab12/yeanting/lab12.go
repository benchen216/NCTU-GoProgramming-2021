package main

import (
	"log"
	"net/http"
	"os"
	"bufio"
	"strings"
	"context"
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
	ObservableMsg = ObservableMsg.Filter(func(i interface{}) bool {
		f_dirty, err := os.Open("dirtytalk.txt")
		defer f_dirty.Close()
		if err != nil{
			log.Fatal("Open file failed!")
			return true
		}
		fileScanner := bufio.NewScanner(f_dirty)
		for fileScanner.Scan() {
			if strings.Contains( i.(string), strings.Replace(fileScanner.Text(), "\n", "", -1) ) {
				return false
			}
		}
		return true
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		f_sen, err := os.Open("sensitive_name.txt")
		defer f_sen.Close()
		if err != nil{
			log.Fatal("Open file failed!")
			return i.(string), nil
		}
		fileScanner := bufio.NewScanner(f_sen)
		var stp_string string
		for fileScanner.Scan() {
			stp_string = strings.Replace(fileScanner.Text(), "\n", "", -1)
			if strings.Contains(i.(string), stp_string) {
				idx := strings.Index(i.(string), stp_string)
				temp := []byte(i.(string))[0:idx]
				rune_idx := len([]rune(string(temp)))
				r_value := []rune(i.(string))
				r_value[rune_idx+1] = '*'
				return string(r_value), nil
			}
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
