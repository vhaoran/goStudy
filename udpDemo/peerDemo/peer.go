package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var tag string

const HAND_SHAKE_MSG = "我是打洞消息"
const SERVER_IP = "47.104.24.37"
const SERVER_PORT = 8888

func main() {
	// 当前进程标记字符串,便于显示
	tag = os.Args[1]
	portOfPeer := 9082
	if len(os.Args) > 2 {
		portOfPeer, _ = strconv.Atoi(os.Args[2])
	}

	// 注意端口必须固定
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: portOfPeer}

	serverAddr := &net.UDPAddr{
		IP:   net.ParseIP(SERVER_IP),
		Port: SERVER_PORT}

	conn, err := net.DialUDP("udp", srcAddr, serverAddr)
	if err != nil {
		fmt.Println(err)
	}
	if _, err = conn.Write([]byte("hello, I'm new peer:" + tag)); err != nil {
		log.Panic(err)
	}

	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}
	conn.Close()
	anotherPeer := parseAddr(string(data[:n]))
	fmt.Printf("local:%s server:%s another:%s\n",
		srcAddr,
		remoteAddr,
		anotherPeer.String())

	// 开始打洞
	digHole(srcAddr, &anotherPeer)
}

func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}

func digHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp",
		srcAddr,
		anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("after net.DialUDP")

	// 向另一个peer发送一条udp消息
	// (对方peer的nat设备会丢弃该消息,非法来源),
	// 用意是在自身的nat设备打开一条可进入的通道,
	// 这样对方peer就可以发过来udp消息
	if _, err = conn.Write([]byte(HAND_SHAKE_MSG)); err != nil {
		log.Println("send handshake:", err)
	}
	log.Println("after send handshake:")

	go func() {
		for {
			time.Sleep(3 * time.Second)
			if _, err = conn.Write([]byte(fmt.Sprint("from ["+tag+"]----> ", conn.LocalAddr().String(), "--->", conn.RemoteAddr().String())));
				err != nil {
				log.Println("############send msg fail", err)
			} else {
				log.Print(tag, " has send")
			}
		}
	}()
	for {
		data := make([]byte, 1024)
		n, cntInfo, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("#########error during read: %s\n", err)
		} else {
			log.Printf("收到数据:%s\n", data[:n], cntInfo.String())
		}
	}
}
