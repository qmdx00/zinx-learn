package znet

type Message struct {
    Id   uint32
    Len  uint32
    Data []byte
}

func NewMessage(id uint32, data []byte) *Message {
    return &Message{
        Id:   id,
        Len:  uint32(len(data)),
        Data: data,
    }
}

func (m *Message) GetMsgId() uint32 {
    return m.Id
}

func (m *Message) GetMsgLen() uint32 {
    return m.Len
}

func (m *Message) GetMsgData() []byte {
    return m.Data
}

func (m *Message) SetMsgId(id uint32) {
    m.Id = id
}

func (m *Message) SetMsgLen(len uint32) {
    m.Len = len
}

func (m *Message) SetMsgData(data []byte) {
    m.Data = data
}
