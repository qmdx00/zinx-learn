package utils

import (
    "encoding/json"
    "io/ioutil"
    "qmdx00.cn/zinx/ziface"
)

type GlobalObj struct {
    TcpServer ziface.IServer
    Host      string
    TcpPort   uint
    Name      string

    Version     string
    MaxConn     int
    MaxPackSize uint32
}

var GlobalObject *GlobalObj

// 从 zinx.json 加载自定义配置
func (g *GlobalObj) load() {
    cfg, err := ioutil.ReadFile("conf/zinx.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(cfg, &GlobalObject)
    if err != nil {
        panic(err)
    }
}

func init() {
    // 默认配置
    GlobalObject = &GlobalObj{
        Name:        "ZinxServerApp",
        Version:     "v0.4",
        TcpPort:     2333,
        Host:        "0.0.0.0",
        MaxConn:     100,
        MaxPackSize: 4096,
    }
    // 自定义配置
    GlobalObject.load()

}
