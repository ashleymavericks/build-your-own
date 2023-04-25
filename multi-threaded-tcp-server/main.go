package main

import (
		"fmt"
		"net"
)

func do(conn net.Conn) {
		buff := make([]byte, 1024) // maximum length of request bytes is 1024
		_, err := conn.Read(buf) // system call -> blocking call
		if err != nil {
				log.Fatal(err)
		}

		log.Println("processing the request")
		time.Sleep(8 * time.Second) // mimicking processing time
		
		// format the string as an HTTP response as per the protocol
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
		conn.Close()

func main() {
		listener, err := net.Listen("tcp", ":1729")
		if err != nil {
			log.Fatal(err)
		}
		
		// infinite for loop
		for {
			log.Println("waiting for a client to connect")
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("client connected"))
	
			// fmt.Println(conn) 
			// do(conn)
			go do(conn) // go routine to fork a new thread to handle requests concurrently, basically spinning off a thread to handle that request independently
		}
}