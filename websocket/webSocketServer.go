package websocket

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var addr = flag.String("addr", ":8083", "http service address")
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
} // use default options
func checkOrigin(r *http.Request) bool {
	return true
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/ws")
	fmt.Println("ws://" + r.Host + "/ws")
}

func WebServer(c *gin.Context) {
	flag.Parse()
	log.SetFlags(0)

	//client.InitAllGroup()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//ServerWs(w, r)
	})
	logrus.Println("收到ws")
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

//func ServerWs(w http.ResponseWriter, r *http.Request) {

var homeTemplate = template.Must(template.New("").Parse(`

`))

/*var upgrader =websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}
func ServerHTTP(w http.ResponseWriter,r *http.Request)  {
	if websocket.IsWebSocketUpgrade(r){
		conn,_:=upgrader.Upgrade(w,r,w.Header())
		conn.WriteMessage(websocket.TextMessage,[]byte("haha"))
		go func() {
			for  {
				ReadMessage(conn)
			}
		}()
	}else {

	}
}
func ReadMessage(conn *websocket.Conn )  {
	messageType,content,_:=conn.ReadMessage()
	fmt.Println(messageType,string(content))
	if messageType==-1 {
		return
	}
}*/
