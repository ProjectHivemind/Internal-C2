package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const BUFFSIZE = 5000

var HIVEMIND_ADDR string = "localhost:1234"

// StartListener start a tcp listening channel on that port
func StartListener(port, hivemindAddr string) {
	serverAddr := "0.0.0.0:" + port
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Starting TCP listener on %s\n", serverAddr)
	defer listener.Close()

	HIVEMIND_ADDR = hivemindAddr
	// hivemindConn, err := net.Dial("tcp", hivemindAddr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := make([]byte, BUFFSIZE) // Needs to be able to accept large registrations, may need to be bigger or done differently
	reader := bufio.NewReader(conn)
	n, err := io.ReadFull(reader, msg)
	if err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			fmt.Println("Read error:", err)
		}
	}
	fmt.Println(string(msg[:n]))

	hivemindConn, _ := net.Dial("tcp", HIVEMIND_ADDR)
	hivemindConnTCP, _ := hivemindConn.(*net.TCPConn)

	defer hivemindConnTCP.Close()

	_, err = hivemindConnTCP.Write(msg[:n])
	if err != nil {
		return
	}

	hivemindConnTCP.CloseWrite()

	msg2 := make([]byte, BUFFSIZE)
	reader2 := bufio.NewReader(hivemindConnTCP)
	n2, err := io.ReadFull(reader2, msg2)
	if err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			fmt.Println("Read error:", err)
		}
	}

	conn.Write(msg2[:n2])
}
