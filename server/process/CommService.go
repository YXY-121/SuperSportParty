package process

import (
	"apiproject/common"
)


//检查是否注册了。
type ServerService struct {
	RegisterUser []string
	SingleSubscribers [] string
	GroupIdMap map[string][]string//每个id对应着一个chan，chan里面装着同一个群聊的人~
}

func NewServerService() *ServerService {
	return &ServerService{
		make([]string,0),
		make([] string,0),
		make(map[string][]string),
	}
}

//是否已经注册
func (c *ServerService)IsRegister(userId string) bool {
	for _,v:=range c.RegisterUser{
		if v==userId {
			return true
		}
	}
	return false
}

//接收到消息，然后群转发
func (c *ServerService)DealGroupMessage(message common.GroupMessage)  {
	if message.Type=="group" {
		groupId:=message.GroupId
		data:=message.Data
		c.BroadCast(data,groupId)
	}//todo 否则报错
}
//群转发
func (c *ServerService)BroadCast(data string,groupId string)  {
	groupSubscribers:=c.FindSubscribersByGroupId(groupId)
	if groupSubscribers==nil {
		//todo 报错
	}
	for _,accepter:=range groupSubscribers{
		c.SendMessage(data,accepter)
	}

}
//根据组 查找相对应的string[]，获取该组里的全部成员
func(c *ServerService) FindSubscribersByGroupId(groupId string)  []string{
	for key,v:=range c.GroupIdMap{
		if key==groupId {
			return v
		}
	}
	return nil
}
//处理点对点的发送
func (c *ServerService)DealSingleMessage(message common.SingleMessage)  {
	if message.Type!="single" {
		//todo 抛出异常
	}
		//	data:=message.Data
			accepter:=message.AccepterId
			//检查是否在线
			if c.IsAccepterOnline(accepter){
				//如果在线的话：


			}else{
	}

			//如果不在线的话：


}
//发送信息
func (c *ServerService)SendMessage(data string,accepter string)  {
		//如何发送给客户端进行一个交互呢。//根据注册的客户端的端口发送吗 如何作导呢
}
//点对点时，判断接收者是否在线
func (c *ServerService)IsAccepterOnline(accepter string) bool {
	for _,v:=range c.SingleSubscribers{
		if v==accepter {
			return true
		}
	}
	return false
}