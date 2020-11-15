package ziface

// Package model { [ len | id ] | [body] }
type IDataPack interface {
    // 包头长度
    GetHeadLen() uint32
    // 封包
    Pack(IMessage) ([]byte, error)
    // 拆包
    UnPack([]byte) (IMessage, error)
}
