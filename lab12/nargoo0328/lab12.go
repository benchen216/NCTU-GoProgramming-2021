package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"bufio"
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

func InitObservable(dirty_talk , sensitive_name []string) {
	ObservableMsg = ObservableMsg.Filter(func(i interface{}) bool{
		for _,talk := range dirty_talk{
			if strings.Contains(i.(string),talk){
				return false
			}
		}
		return true
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		str := i.(string)
		chat_index := strings.Index(i.(string),"表示:")
		chat := []rune(i.(string))
		chat_raw :=chat[chat_index+4:]
		chat_content := string(chat_raw)
		for _,name := range sensitive_name{
			chat = []rune(str)
			if strings.Contains(chat_content,name){
				// log.Println(name)
				raw_name := []rune(name)
				pos_2 := check_substrings(chat_raw,raw_name)
				// log.Println(chat_index,pos_2)
				str = string(chat[:chat_index+4+pos_2+1])+"*"+string(chat[chat_index+4+pos_2+2:])
			}
		}
		return str,nil
	})

}

func check_substrings(str1 , str2 []rune) int{
	log.Println(len(str1),len(str2))
	for i:=0;i<=len(str1)-1-len(str2);i++{
		for j,char := range str2{
			if str1[i+j] != char{
				break
			}
			if j==len(str2)-1{
				return i
			}
		}
	}
	return -1
}

func readfile(file_name string) [] string{
	content := [] string{}
	f, _ := os.Open(file_name)
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
      line := scanner.Text()
      content = append(content,line)
    }
	return content
}

func main() {
	dirty_talk := readfile("dirtytalk.txt")
	sensitive_name := readfile("sensitive_name.txt")
	InitObservable(dirty_talk,sensitive_name)
	go broadcaster()
	http.HandleFunc("/wschatroom", wshandle)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("server start at :8899")
	log.Fatal(http.ListenAndServe(":8899", nil))
}
