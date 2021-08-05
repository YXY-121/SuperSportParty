package client

var AllHub =make(map[string]*Hub)

type Hub struct {
	HubId string
	Register chan *ClientService//群的总人员
	Group 	map[string] []*ClientService//jilu
	Clients  map[*ClientService]bool//当前在线人员
	GroupBroadCast chan []byte

	//map[string]chan[]byte
}

func NewHub(hubId string) *Hub {
	return &Hub{
		HubId: hubId,
		Register: make(chan *ClientService),
		Group:make(map[string][]*ClientService),
		Clients:make(map[*ClientService]bool),
		GroupBroadCast: make(chan []byte),
	}
}
func (h *Hub)Run(){
	for  {

		select {
		case client:=<-h.Register:
			h.Clients[client]=true

		case message:=<-h.GroupBroadCast:
			for client := range h.Clients {
				select {
				case client.AcceptedMessages <- message:
				default:
					close(client.AcceptedMessages)
					delete(h.Clients, client)
				}
			}

		}


	}
}