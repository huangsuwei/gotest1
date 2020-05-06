package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

//创建全局map，存储在线用户
var onlineMap map[string]Client

//创建全局channel传递用户信息
var message = make(chan string)

func WriteMsgToClient(clnt Client, conn net.Conn) {
	//监听用户自带channel上是否有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(clnt Client, msg string) (buf string) {
	return "[" + clnt.Addr + "]" + clnt.Name + ":" + msg
}

func handlerConnect(conn net.Conn) {
	defer conn.Close()
	//获取用户ip port
	netAddr := conn.RemoteAddr().String()

	//创建新连接用户的结构体 默认是ip+port
	clnt := Client{make(chan string), netAddr, netAddr}

	//将新连接用户，添加到在线用户map中，key:ip+port，value:Client
	onlineMap[netAddr] = clnt

	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(clnt, conn)

	//发送用户上线消息到全局channel中
	//message <- "[" + netAddr + "]" + clnt.Name + "login"
	message <- MakeMsg(clnt, "login")

	//创建一个channel，用来判断用户是否退出
	isQuit := make(chan bool)
	//用来判断用户是否活跃，
	hasData := make(chan bool)

	//创建一个匿名go程，专门处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端：%s退出\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("coon.read err", err)
				return
			}
			//将读到的用户消息，广播给所有用户，写入到message中
			msg := string(buf[:n-1]) //这里减一是为了去掉\n
			//提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				//遍历当前map，获取在线用户
				for _, user := range onlineMap {
					conn.Write([]byte(user.Addr + ":" + user.Name + "\n"))
				}
				//判断用户发送了改名命令
			} else if len(msg) >= 8 && msg[:6] == "rename" { //rename|
				clnt.Name = strings.Split(msg, "|")[1]
				//或者 clnt.Name = msg[8:]
				//更新到onlineMap中
				onlineMap[clnt.Addr] = clnt
				conn.Write([]byte("rename success!" + "\n"))
			} else {
				message <- MakeMsg(clnt, msg)
			}
			hasData <- true
		}
	}()

	//保证不退出
	for {
		//监听channel上的数据流动
		select {
		case <-isQuit:
			delete(onlineMap, clnt.Addr)       //将用户从online移除
			message <- MakeMsg(clnt, "logout") //写入用户退出消息到全局channel
			return                             //结束当前go程
		case <-hasData:
			//条件满足不做任何操作，目的是为了重置定时器
		//超时强制踢出
		case <-time.After(120 * time.Second):
			delete(onlineMap, clnt.Addr)       //将用户从online移除
			message <- MakeMsg(clnt, "logout") //写入用户退出消息到全局channel
			return                             //结束当前go程
		}
	}
}

func Manager() {
	//初始化onlineMap
	onlineMap = make(map[string]Client)

	//循环从message中读取内容
	for {
		//监听全局chan中是否有数据，有数据存储到msg，无数据阻塞
		msg := <-message

		//循环发送消息给所有在线用户
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer listener.Close()

	//创建管理者go程，管理map和全局message
	go Manager()

	//循环监听客户端连接请求
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("listen.accept err", err1)
			return
		}
		//启动go程处理客户端数据请求
		go handlerConnect(conn)
	}
}
