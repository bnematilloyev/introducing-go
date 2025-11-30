package main

// Communicating between multiple computers: TCP, HTTP and RPC servers

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

// TCP -> Transmission Control Protocol

//	Qachon ishlatiladi:
//		Past darajadagi, maxsus protokol yaratmoqchi bo‘lsangiz.
//		Masalan, o‘z chat serveringiz, real-time o‘yin serveri, IoT qurilmalari.

/*
type Listener interface {
	Accept() (c Conn, err error)
	// Accept waits for and returns the next connection to the listener

	Close() error
	// Close() closes the listener
	// Any blocked Accept operations will be unblocked and return errors

	Addr()
	// Addr returns the listener's network address
}
*/

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	// connect the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	// TCP
	go server()
	time.Sleep(time.Second)
	go client()

	var input string
	fmt.Scanln(&input)

	//Sending Hello, World
	//Received Hello, World
}
