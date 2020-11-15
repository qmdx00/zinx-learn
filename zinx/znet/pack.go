package znet

import (
    "bytes"
    "encoding/binary"
    "errors"
    "qmdx00.cn/zinx/utils"
    "qmdx00.cn/zinx/ziface"
)

type DataPack struct {

}

func NewDataPack() *DataPack {
    return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
    // uint32 len + uint32 id
    // uint32 = 8 byte = 32 bit
    return 8
}

// Message 封包成 byte 切片
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

    buf := bytes.NewBuffer([]byte{})
    // write msg length
    if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
        return nil, err
    }
    // write msg id
    if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
        return nil, err
    }
    // write msg data
    if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgData()); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

// byte 切片拆包成 Message，读取成 Head，在根据 Head 读取 body
func (d *DataPack) UnPack(bin []byte) (ziface.IMessage, error) {

    buf := bytes.NewReader(bin)
    msg := &Message{}
    // read head msg len
    if err := binary.Read(buf, binary.LittleEndian, &msg.Len); err != nil {
        return nil, err
    }
    // read head msg id
    if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
        return nil, err
    }
    if utils.GlobalObject.MaxPackSize > 0 && msg.GetMsgLen() > utils.GlobalObject.MaxPackSize {
        return nil, errors.New("[UnPack Message Error] Too Large Message Data Received ")
    }

    return msg, nil
}
