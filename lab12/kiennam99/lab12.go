package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
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

	// ObservableMsg = ObservableMsg.Filter(...) ... {
	// }).Map(...) {
	// 	...
	// })
	content, err := ioutil.ReadFile("dirtytalk.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))
	thecont := strings.Split(string(content), "\n")
	content, err = ioutil.ReadFile("sensitive_name.txt")
	if err != nil {
		log.Fatal(err)
	}
	themap := strings.Split(string(content), "\n")
	// fmt.Println(thecont[0])
	ObservableMsg = ObservableMsg.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		inputstr := i.(string)
		for _, k := range themap {
			if strings.Contains(inputstr, k) {
				inputstr = strings.Replace(inputstr, k, string(k[0:3])+"*"+string(k[6:]), -1)
				// fmt.Printf("%x ", k)
			}
		}
		return inputstr, nil
	})
	ObservableMsg = ObservableMsg.Filter(func(i interface{}) bool {

		for _, k := range thecont {
			if strings.Contains(i.(string), k) {
				return false
			}
		}
		return true
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
