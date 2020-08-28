package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"protobuf_test/proto_def"
	"time"
)

func main() {
	strIpP:="localhost:6600"
	var conn net.Conn
	var err error

	for conn,err= net.Dial("tcp",strIpP);err!=nil;conn,err=net.Dial("tcp",strIpP){
		fmt.Printf("connect %s failed",err)
		time.Sleep(time.Second)
		fmt.Println("reconnect")
	}

	fmt.Printf("connected %s success",strIpP)
	defer conn.Close()
	cnt:= 0
	sender:=bufio.NewScanner(os.Stdin)
	for sender.Scan(){
		cnt++
		stSend:=&proto_def.UserInfo{
			Message:       sender.Text(),
			Length:        *proto.Int(len(sender.Text())),
			Cnt:           *proto.Int(cnt),
		}
		pData,err:=proto.Marshal(stSend)

		if err != nil {
			panic(err)
		}
		conn.Write(pData)
		if sender.Text()=="stop"{
			return
		}
	}

}
