package client

import (
	"apiproject/common"
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "127.0.0.1:8081", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	message:=common.GroupMessage{
		Type: common.GroupMessageType,
		Content:"hi",
		GroupId: "1",
	}
	b,_:=json.Marshal(&message)
	for  {
		err=conn.WriteMessage(websocket.TextMessage,b)
		if err != nil {
			log.Println("write:", err)
			return
		}
	////	_, response, _ := conn.ReadMessage()
	//	message1:=common.Message{}
	//	json.Unmarshal(response,&message1)//没有回复
	//	fmt.Println("收货")
	//	fmt.Println(message1)
		time.Sleep(5*time.Second)
	}

/*	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()*/

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	//for {
	//	select {
	//	case <-done:
	//		return
	//	case t := <-ticker.C:
	//		err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
	//		if err != nil {
	//			log.Println("write:", err)
	//			return
	//		}
	//	case <-interrupt:
	//		log.Println("interrupt")
	//
	//		// Cleanly close the connection by sending a close message and then
	//		// waiting (with timeout) for the server to close the connection.
	//		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//		if err != nil {
	//			log.Println("write close:", err)
	//			return
	//		}
	//		select {
	//		case <-done:
	//		case <-time.After(time.Second):
	//		}
	//		return
	//	}
	//}
}

