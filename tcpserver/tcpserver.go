package tcpserver

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

func Listen(port int, callback func(bs []byte)) (net.Listener, error) {
	p := strconv.Itoa(port)
	adress := "127.0.0.1:" + p
	ln, err := net.Listen("tcp", adress)
	if err != nil {
		return nil, err
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("发生了未处理的异常", err)
			}
		}()
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println("accept err: ", err.Error())
				continue
			}
			go handle(conn, callback)
		}
	}()
	return ln, nil
}

func handle(conn net.Conn, callback func(bs []byte)) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("发生了未处理的异常", err)
		}
	}()
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Bytes()
		callback(line)
	}
	if err := scanner.Err(); err != nil {
		log.Println("错误：", err)
	}
}
