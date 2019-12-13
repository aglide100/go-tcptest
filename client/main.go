
package main

import (
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:5000"

	conn, err := net.Dial("tcp", addr)	
	if err != nil{
		log.Fatal(err)
	}


	defer conn.Close()

	msg := []byte("hello you to aaangfkrgzfdjbbk")
	_, err = conn.Write(msg)
	if err != nil{
		log.Fatal(err)
	}


}