package main

import (
	"fmt"
	"net"
	"parsing"
)

func readloop(conn *net.TCPConn) {
	address := conn.RemoteAddr()
	player := parsing.NewPlayer(address.String())
	for {
		b := make([]byte, 4096)
		n, err := conn.Read(b[:])
		data := b[:n]
		if err != nil {
			fmt.Println(err)
		}

		if len(string(data)) == 0 {
			fmt.Println("PARTING:", address.String())
			conn.Close()
			return
		}
		conn.Write(parsing.Parse(player,string(data)))
		//fmt.Printf("%s", string(data))
	}
}

func main() {
	i, _adderr := net.ResolveTCPAddr("tcp", "0.0.0.0:6580")
	l, _liserr := net.ListenTCP("tcp", i)
	for {
		conn, _accerr := l.AcceptTCP()

		address := conn.RemoteAddr()
		fmt.Println("INCOMING:", address.String())

		go readloop(conn)

		if _accerr != nil {
			fmt.Println(_accerr)
		}
	}
	l.Close()
	fmt.Println(_adderr)
	fmt.Println(_liserr)
}
