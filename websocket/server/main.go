
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
	homeTemplate.Execute(w, "ws://"+r.Host+"/ws")
	fmt.Println("ws://"+r.Host+"/ws")
}

func WebServer() {
	flag.Parse()
	log.SetFlags(0)

	client.InitAllGroup()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
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

</head>
<body>
<table>
<tr><td valign="top" width="50%">
这里是群聊~
<form>
<!-- 1 -->
<input id="inputGroup" type="text" value="Hello world!">
<input id="groupId" type="text" value="输入组名!">

<button id="sendGroup">SendGroup</button>
</form>
<br>
<br>

这里是私人聊~
<form>
<!-- 2 -->
<input id="inputSingle" type="text" value="Hello world!">
<input id="singleId" type="text" value="输入要发送的朋友Id!">
<button id="sendSingle">SendSingle</button>
</form>
<br>
<br>

创建新群~,好友id之间用,来分隔
<form>
<input id="inputCreateGroup" type="text" value="新建群名">
<input id="CreateGroupUserIds" type="text" value="输入拉入群的好友id">
<button id="sendCreateGroup">CreateGroup</button>
</form>
<br>
<br>
点击open时，第一次发送信息前要输入userId（模拟登录）

<form>
<input id="inputUserId" type="text" value="输入userId">
<button id="sendUserId">sendUserID</button>
</form>
<button id="open">Open</button>
<button id="close">Close</button>
<br>
<br>

</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
<script>  
console.log( document.getElementById("sendUserId"));
 
  onclick =()=>{
      console.log(111112111);
  }
    window.addEventListener("load", function(evt) {
        var output = document.getElementById("output");
        var input = document.getElementById("input");
    
    
        var inputUserId=document.getElementById("inputUserId");
        var sendUserID=document.getElementById("sendUserId");
    
    
        var singleId=document.getElementById("singleId");
        var inputSingle=document.getElementById("inputSingle");
        var sendSingle=document.getElementById("sendSingle");
    
        
        var groupId=document.getElementById("groupId");
        var inputGroup=document.getElementById("inputGroup");
        var sendGroup=document.getElementById("sendGroup");
    
    
     	 var inputCreateGroup=document.getElementById("inputCreateGroup");
        var CreateGroupUserIds=document.getElementById("CreateGroupUserIds");
        var sendCreateGroup=document.getElementById("sendCreateGroup");
     
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
        //发送群聊
        document.getElementById("sendGroup").onclick = function(evt) {
            if (!ws) {
                return false;
            }
        var messageObj = {group_id:groupId.value,content:inputGroup.value,type:"group"};
        var messageJson = JSON.stringify(messageObj);
        ws.send(messageJson);
            return false;
   
        //var groupId=document.getElementById("groupId");
        //var inputGroup=document.getElementById("inputGroup");
        //var sendGroup=document.getElementById("sendGroup");
        };
    
        //发送私人聊天
        document.getElementById("sendSingle").onclick = function(evt) {
            if (!ws) {
                return false;
            }
        var messageObj = {accepter_id:singleId.value,content:inputSingle.value,type:"single"};
        var messageJson = JSON.stringify(messageObj);
        ws.send(messageJson);
            return false;
        };
    //创建新群
           document.getElementById("sendCreateGroup").onclick = function(evt) {
            if (!ws) {
                return false;
            }
        var user_idsData =CreateGroupUserIds.value.split(",")
        var messageObj = {group_name:inputCreateGroup.value,user_ids:user_idsData,type:"createGroup"};
        var messageJson = JSON.stringify(messageObj);
        ws.send(messageJson);
            return false;
        };
    
    
        // 模拟登录 发送userid
        document.getElementById("sendUserId").onclick = function(evt) {
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