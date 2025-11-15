package znet

type Message struct {
	Id      uint32 // 消息的 ID
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// 创建 Message 消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 还需要为 Message 实现 IMessage 的六个方法
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetDateLen(length uint32) {
	msg.DataLen = length
}

func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
