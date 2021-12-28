package main

import (
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

func get_content_of_file(filepath string) []string {
	dat, err := os.ReadFile(filepath)
	check(err)

	var return_strings []string
	var tmp_str []byte
	for i := 0; i < len(dat); i++ {
		if string(dat[i]) == "\n" {
			// fmt.Printf("%s", string(tmp_str))
			return_strings = append(return_strings, string(tmp_str))
			tmp_str = []byte{}
		} else {
			tmp_str = append(tmp_str, dat[i])
		}
	}

	return return_strings
}

func InitObservable() {
	/* 參考
	ObservableMsg = ObservableMsg.Filter(...) ... {
	}).Map(...) {
		...
	})
	*/
	ObservableMsg = ObservableMsg.Filter(func(inte interface{}) bool {
		// ObservableMsg = ObservableMsg.Filter(func(inte interface{}) bool {}
		inp_str := fmt.Sprintf("%v", inte)

		var dirtytalks = get_content_of_file("dirtytalk.txt")

		find_begin := strings.LastIndex(inp_str, ": ")
		find_end := strings.LastIndex(inp_str, "\n")
		// fmt.Println(find_begin, find_end)
		cut_str := inp_str[find_begin+2 : find_end]

		for i := 0; i < len(dirtytalks); i++ {
			if strings.Contains(cut_str, dirtytalks[i]) {
				return false
			}
			// if cut_str == dirtytalks[i] {
			// 	return false
			// }
		}
		return true
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		inp_str := fmt.Sprintf("%v", i)

		// var dirtytalks = get_content_of_file("dirtytalk.txt")
		var sensitive_names = get_content_of_file("sensitive_name.txt")

		// for idx := 0; idx < len(dirtytalks); idx++ {
		// 	inp_str = strings.Replace(inp_str, dirtytalks[idx], "", -1)
		// }

		for idx := 0; idx < len(sensitive_names); idx++ {
			change_name := sensitive_names[idx][0:3] + "*" + sensitive_names[idx][6:]
			// fmt.Printf("%s, %s, %d\n", sensitive_names[idx][0:1], sensitive_names[idx][2:], len(sensitive_names[idx]))
			inp_str = strings.Replace(inp_str, sensitive_names[idx], change_name, -1)
		}
		// fmt.Println(inp_str)
		return inp_str, nil
	})

	// for i := 0; i < len(dirtytalks); i++ {
	// 	fmt.Printf("my word: \"%s\" \n", dirtytalks[i])
	// }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	InitObservable()
	go broadcaster()
	http.HandleFunc("/wschatroom", wshandle)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("server start at :8899")
	log.Fatal(http.ListenAndServe(":8899", nil))
}
