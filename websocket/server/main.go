
package webSocketServer

import (
	client "apiproject/websocket/client/service"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8081", "http service address")
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: checkOrigin,
} // use default options
func checkOrigin(r *http.Request) bool {
	return true
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
	fmt.Println("ws://"+r.Host+"/ws")
}

func WebServer() {
	flag.Parse()
	log.SetFlags(0)

	client.InitAllGroup()

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		ServerWs( w, r)
	})
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}


func ServerWs(w http.ResponseWriter, r *http.Request) {
	//升级成websocket协议
	//每次都创建一个client

	conn, err := upgrader.Upgrade(w, r, nil)
	client:=&client.ClientService{
		Conn: conn,
		AcceptedMessages: make(chan[]byte,256),
	}





	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	if err != nil {
		log.Println("read:", err)
		return
	}

	go client.ReturnAccectped()
	go client.SendOther()


}


var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
	var inputG = document.getElementById("inputGroup");
	var inputType=document.getElementById("inputType");
	var inputUserId=document.getElementById("inputUserId");

    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
	var messageObj = {group_id:inputG.value,content:input.value,type:inputType.value};
	var messageJson = JSON.stringify(messageObj);
	ws.send(messageJson);
        return false;
    };
    document.getElementById("sendUserID").onclick = function(evt) {
       if (!ws) {
           return false;
       }
	var messageObj = {user_id:inputUserId.value};
	var messageJson = JSON.stringify(messageObj);
	ws.send(messageJson);
       return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">

<form>
<button id="open">Open</button>
<button id="close">Close</button>
<input id="input" type="text" value="Hello world!">
<input id="inputGroup" type="text" value="输入组名!">
<input id="inputType" type="text" value="输入类型">


<button id="send">Send</button>

</form>
<form>
<input id="inputUserId" type="text" value="输入userId">
<button id="sendUserID">sendUserID</button>

</form>


</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
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