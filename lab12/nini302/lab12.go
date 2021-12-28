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
func readlines(filename string) map[string]string {
	ret := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//  fmt.Println(scanner.Text())
		s := scanner.Text()
		censored := []rune(s)
		censored[1] = '*'
		ret[scanner.Text()] = string(censored)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ret
}
func InitObservable() {

	dirtytalk := readlines("dirtytalk.txt")
	sensitive_name := readlines("sensitive_name.txt")

	ObservableMsg = ObservableMsg.Filter(func(i interface{}) bool {
		i_str, _ := i.(string)
		for index, _ := range dirtytalk {
			if strings.Contains(i_str, index) {
				return false
			}
		}
		return true

	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		i_str, _ := i.(string)
		for index, value := range sensitive_name {
			if strings.Contains(i_str, index) {
				i_str = strings.Replace(i_str, index, value, -1)
			}
		}
		return i_str, nil
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
